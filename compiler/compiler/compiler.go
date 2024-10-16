package compiler

import (
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/debug"

	"github.com/cell-labs/cell-script/compiler/compiler/internal"
	"github.com/cell-labs/cell-script/compiler/compiler/internal/pointer"
	"github.com/cell-labs/cell-script/compiler/compiler/name"
	"github.com/cell-labs/cell-script/compiler/compiler/strings"
	"github.com/cell-labs/cell-script/compiler/compiler/types"
	"github.com/cell-labs/cell-script/compiler/compiler/value"
	"github.com/cell-labs/cell-script/compiler/parser"

	"errors"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	llvmTypes "github.com/llir/llvm/ir/types"
	llvmValue "github.com/llir/llvm/ir/value"
)

type Options struct {
	Target string
}

type Compiler struct {
	module *ir.Module

	// functions provided by the OS, such as printf and malloc
	osFuncs OSFuncs

	packages       map[string]*pkg
	currentPackage *pkg

	// TODO: Replace with currentPackage.Name()
	currentPackageName string

	contextFunc *types.Function

	initGlobalsFunc *ir.Func
	mainFunc        *ir.Func

	// Stack of return values pointers, is used both used if a function returns more
	// than one value (arg pointers), and single stack based returns
	contextFuncRetVals [][]value.Value

	contextBlock *ir.Block

	// Stack of variables that are in scope
	contextBlockVariables []map[string]value.Value

	// What a break or continue should resolve to
	contextLoopBreak    []*ir.Block
	contextLoopContinue []*ir.Block

	// Where a condition should jump when done
	contextCondAfter []*ir.Block

	// What type the current assign operation is assigning to.
	// Is used when evaluating what type an integer constant should be.
	contextAssignDest []value.Value

	// Stack of Alloc instructions
	// Is used to decide if values should be stack or heap allocated
	contextAlloc []*parser.AllocNode

	stringConstants map[string]*ir.Global

	// runtime.GOOS and runtime.GOARCH
	GOOS, GOARCH string
}

var (
	boolean = types.Bool
	i8      = types.I8
	i32     = types.I32
	i64     = types.I64
)

func NewCompiler(options *Options) *Compiler {
	c := &Compiler{
		module: ir.NewModule(),

		packages: make(map[string]*pkg),

		contextFuncRetVals: make([][]value.Value, 0),

		contextBlockVariables: make([]map[string]value.Value, 0),

		contextLoopBreak:    make([]*ir.Block, 0),
		contextLoopContinue: make([]*ir.Block, 0),
		contextCondAfter:    make([]*ir.Block, 0),

		contextAssignDest: make([]value.Value, 0),

		stringConstants: make(map[string]*ir.Global),
	}

	c.createExternalPackage()
	c.addGlobal()
	c.pushVariablesStack()

	// Triple examples:
	// x86_64-apple-macosx10.13.0
	// x86_64-pc-linux-gnu
	var targetTriple [2]string

	switch runtime.GOARCH {
	case "amd64":
		targetTriple[0] = "x86_64"
	case "arm64":
		targetTriple[0] = "aarch64"
	default:
		panic("unsupported GOARCH: " + runtime.GOARCH)
	}

	switch runtime.GOOS {
	case "darwin":
		targetTriple[1] = "apple-macosx10.13.0"
	case "linux":
		targetTriple[1] = "pc-linux-gnu"
	case "windows":
		targetTriple[1] = "pc-windows"
	default:
		panic("unsupported GOOS: " + runtime.GOOS)
	}

	if options.Target != "native" {
		targetTriple[0] = options.Target
		targetTriple[1] = "unknown"
	}

	c.module.TargetTriple = fmt.Sprintf("%s-%s", targetTriple[0], targetTriple[1])

	// TODO: Allow cross compilation
	c.GOOS = runtime.GOOS
	c.GOARCH = runtime.GOARCH

	return c
}

func (c *Compiler) Compile(root parser.PackageNode) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// Compile time panics, that are not errors in the compiler
			if _, ok := r.(Panic); ok {
				err = errors.New(fmt.Sprint(r))
				return
			}

			// Bugs in the compiler
			err = fmt.Errorf("%s\n\nInternal compiler stacktrace:\n%s",
				fmt.Sprint(r),
				string(debug.Stack()),
			)
		}
	}()

	c.currentPackageName = filepath.Base(root.Name)
	if !c.IsPackageImported(c.currentPackageName) {
		c.currentPackage = NewPkg(c.currentPackageName)
		c.packages[c.currentPackageName] = c.currentPackage
	} else {
		c.currentPackage = c.packages[c.currentPackageName]
	}

	for _, fileNode := range root.Files {
		c.compile(fileNode.Instructions)
	}

	return
}

func (c *Compiler) GetIR() string {
	return c.module.String()
}

func (c *Compiler) IsPackageImported(name string) bool {
	_, ok := c.packages[name]
	return ok
}

func (c *Compiler) addGlobal() {
	types.ModuleStringType = c.module.NewTypeDef("string", internal.String())

	// Create empty string constant
	types.EmptyStringConstant = c.module.NewGlobalDef(strings.NextStringName(), strings.Constant(""))
	types.EmptyStringConstant.Immutable = true

	// TODO: Use a different name? Runtime?
	global := NewPkg("global")

	strLen := internal.StringLen(types.ModuleStringType)
	global.DefinePkgVar("len_string", value.Value{
		Type: &types.Function{
			FuncType:       strLen.Type(),
			LlvmReturnType: types.I64,
		},
		Value:      strLen,
		IsVariable: false,
	})

	global.DefinePkgType("bool", types.Bool)
	global.DefinePkgType("int", types.I64)  // TODO: Size based on arch
	global.DefinePkgType("uint", types.U64) // TODO: Size based on arch
	global.DefinePkgType("int8", types.I8)
	global.DefinePkgType("uint8", types.U8)
	global.DefinePkgType("int16", types.I16)
	global.DefinePkgType("uint16", types.U16)
	global.DefinePkgType("int32", types.I32)
	global.DefinePkgType("uint32", types.U32)
	global.DefinePkgType("int64", types.I64)
	global.DefinePkgType("uint64", types.U64)
	global.DefinePkgType("uint128", types.U128)
	global.DefinePkgType("uint256", types.U256)
	global.DefinePkgType("uintptr", types.Uintptr)
	global.DefinePkgType("string", types.String)
	global.DefinePkgType("byte", types.U8)

	c.packages["global"] = global

	c.module.Funcs = append(c.module.Funcs, strLen)

	// Initialization function
	c.initGlobalsFunc = c.module.NewFunc(name.Var("global-init"), types.Void.LLVM())
	b := c.initGlobalsFunc.NewBlock(name.Block())
	b.NewRet(nil)

	// main.main function, body will be added later
	c.mainFunc = c.module.NewFunc("main", types.I64.LLVM()) // TODO: Size based on arch
	mainBlock := c.mainFunc.NewBlock(name.Block())
	mainBlock.NewCall(c.initGlobalsFunc)
}

func (c *Compiler) compile(instructions []parser.Node) {
	for _, i := range instructions {
		switch v := i.(type) {
		case *parser.ConditionNode:
			c.compileConditionNode(v)
		case *parser.DefineFuncNode:
			c.compileDefineFuncNode(v)
		case *parser.ReturnNode:
			c.compileReturnNode(v)
		case *parser.AllocNode:
			c.compileAllocNode(v)
		case *parser.AllocGroup:
			for _, a := range v.Allocs {
				c.compileAllocNode(a)
			}
		case *parser.AssignNode:
			c.compileAssignNode(v)
		case *parser.ForNode:
			c.compileForNode(v)
		case *parser.BreakNode:
			c.compileBreakNode(v)
		case *parser.ContinueNode:
			c.compileContinueNode(v)
		case *parser.ExternNode:
			for _, fn := range v.FuncNodes {
				c.compileDefineFuncNode(fn)
			}

		case *parser.DeclarePackageNode:
			// TODO: Make use of it
			break
		case *parser.ImportNode:
			// NOOP
			break
		case *parser.PragmaNode:
			// NOOP
			break

		case *parser.DefineTypeNode:
			t := c.parserTypeToType(v.Type)

			// Add type to module and override the structtype to use the named
			// type in the module
			if structType, ok := t.(*types.Struct); ok {
				structType.Type = c.module.NewTypeDef(v.Name, t.LLVM())
			}

			// Add to tre mapping
			c.currentPackage.DefinePkgType(v.Name, t)
		case *parser.SwitchNode:
			c.compileSwitchNode(v)

		default:
			c.compileValue(v)
			break
		}
	}
}

func (c *Compiler) compileNameNode(v *parser.NameNode) value.Value {
	pkg := c.currentPackage
	inSamePackage := true

	if len(v.Package) > 0 {
		// Imported package?
		if p, ok := c.packages[v.Package]; ok {
			pkg = p
			inSamePackage = false
		} else {
			panic(fmt.Sprintf("package %s does not exist", v.Package))
		}
	}

	// Search scope in reverse (most specific first)
	for i := len(c.contextBlockVariables) - 1; i >= 0; i-- {
		if val, ok := c.contextBlockVariables[i][v.Name]; ok {
			return val
		}
	}

	if pkgVar, ok := pkg.GetPkgVar(v.Mangling, inSamePackage); ok {
		return pkgVar
	}

	if pkgVar, ok := pkg.GetPkgVar(v.Name, inSamePackage); ok {
		return pkgVar
	}


	panic(fmt.Sprintf("package %s has no memeber %s/%s", v.Package, v.Name, v.Mangling))
}

func (c *Compiler) setVar(name string, val value.Value) {
	c.contextBlockVariables[len(c.contextBlockVariables)-1][name] = val
}

func (c *Compiler) pushVariablesStack() {
	c.contextBlockVariables = append(c.contextBlockVariables, make(map[string]value.Value))
}

func (c *Compiler) popVariablesStack() {
	c.contextBlockVariables = c.contextBlockVariables[0 : len(c.contextBlockVariables)-1]
}

func (c *Compiler) compileValue(node parser.Node) value.Value {
	switch v := node.(type) {

	case *parser.ConstantNode:
		return c.compileConstantNode(v)
	case *parser.OperatorNode:
		return c.compileOperatorNode(v)
	case *parser.SubNode:
		return c.compileSubNode(v)
	case *parser.NameNode:
		return c.compileNameNode(v)
	case *parser.CallNode:
		return c.compileCallNode(v)
	case *parser.TypeCastNode:
		return c.compileTypeCastNode(v)
	case *parser.StructLoadElementNode:
		return c.compileStructLoadElementNode(v)
	case *parser.LoadArrayElement:
		return c.compileLoadArrayElement(v)
	case *parser.GetReferenceNode:
		return c.compileGetReferenceNode(v)
	case *parser.DereferenceNode:
		return c.compileDereferenceNode(v)
	case *parser.NegateNode:
		return c.compileNegateBoolNode(v)
	case *parser.InitializeSliceNode:
		return c.compileInitializeSliceNode(v)
	case *parser.InitializeStringWithSliceNode:
		return c.compileInitializeStringWithSliceNode(v)
	case *parser.SliceArrayNode:
		src := c.compileValue(v.Val)

		if _, ok := src.Type.(*types.StringType); ok {
			return c.compileSubstring(src, v)
		}

		if ty, ok := src.Type.(*types.Slice); ok {
			src.Value = internal.LoadIfVariable(c.contextBlock, src)
			src.Type = ty.Type
			return c.compileSliceArray(src, v, true)
		}
		// array type as pointer receriver
		if ptr, ok := src.Type.(*types.Pointer); ok {
			src.Type = ptr.Type
			src.Value = c.contextBlock.NewLoad(pointer.ElemType(src.Value), src.Value)
		}
		return c.compileSliceArray(src, v, false)
	case *parser.InitializeStructNode:
		return c.compileInitStructWithValues(v)
	case *parser.TypeCastInterfaceNode:
		return c.compileTypeCastInterfaceNode(v)
	case *parser.DefineFuncNode:
		return c.compileDefineFuncNode(v)
	case *parser.InitializeArrayNode:
		return c.compileInitializeArrayNode(v)
	case *parser.DecrementNode:
		return c.compileDecrementNode(v)
	case *parser.IncrementNode:
		return c.compileIncrementNode(v)
	case *parser.GroupNode:
		return c.compileGroupNode(v)
	}

	panic("compileValue fail: " + fmt.Sprintf("%T: %+v", node, node))
}

func (c *Compiler) panic(block *ir.Block, message string) {
	globMsg := c.module.NewGlobalDef(strings.NextStringName(), strings.Constant("runtime panic: "+message+"\n"))
	globMsg.Immutable = true
	block.NewCall(c.osFuncs.Printf.Value.(llvmValue.Named), strings.Toi8Ptr(block, globMsg))
	// todo: specify panic exit value
	block.NewCall(c.osFuncs.Exit.Value.(llvmValue.Named), constant.NewInt(llvmTypes.I8, -1))
}

type Panic string

func compilePanic(message string) {
	panic(Panic(fmt.Sprintf("compile panic: %s\n", message)))
}
