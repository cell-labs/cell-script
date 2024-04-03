package compiler

import (
	"github.com/llir/llvm/ir"
	llvmTypes "github.com/llir/llvm/ir/types"

	"github.com/cell-labs/cell-script/compiler/compiler/types"
	"github.com/cell-labs/cell-script/compiler/compiler/value"
)

// OSFuncs and the "debug" package contains a mapping to glibc functions.
// These are used to make bootstrapping of the language easier. The end goal is to not depend on glibc.
type OSFuncs struct {
	Printf  value.Value
	Malloc  value.Value
	Realloc value.Value
	Memcpy  value.Value
	Strcat  value.Value
	Strcpy  value.Value
	Strncpy value.Value
	Strndup value.Value
	Exit    value.Value
}

func (c *Compiler) createExternalPackage() {
	external := NewPkg("debug")

	setExternal := func(internalName string, fn *ir.Func, variadic bool) value.Value {
		fn.Sig.Variadic = variadic
		val := value.Value{
			Type: &types.Function{
				LlvmReturnType: types.Void,
				FuncType:       fn.Type(),
				IsExternal:     true,
			},
			Value: fn,
		}
		external.DefinePkgVar(internalName, val)
		return val
	}

	c.osFuncs.Printf = setExternal("Printf", c.module.NewFunc("printf",
		i32.LLVM(),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
	), true)

	c.osFuncs.Malloc = setExternal("malloc", c.module.NewFunc("malloc",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", i64.LLVM()),
	), false)

	c.osFuncs.Realloc = setExternal("realloc", c.module.NewFunc("realloc",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", i64.LLVM()),
	), false)

	c.osFuncs.Memcpy = setExternal("memcpy", c.module.NewFunc("memcpy",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("dest", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("src", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("n", i64.LLVM()),
	), false)

	c.osFuncs.Strcat = setExternal("strcat", c.module.NewFunc("strcat",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
	), false)

	c.osFuncs.Strcpy = setExternal("strcpy", c.module.NewFunc("strcpy",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
	), false)

	c.osFuncs.Strncpy = setExternal("strncpy", c.module.NewFunc("strncpy",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", i64.LLVM()),
	), false)

	c.osFuncs.Strndup = setExternal("strndup", c.module.NewFunc("strndup",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", i64.LLVM()),
	), false)

	// c.osFuncs.Exit = setExternal("exit", c.module.NewFunc("exit",
	// 	llvmTypes.Void,
	// 	ir.NewParam("", i32.LLVM()),
	// ), false)

	c.packages["debug"] = external
}
