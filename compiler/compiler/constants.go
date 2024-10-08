package compiler

import (
	"github.com/cell-labs/cell-script/compiler/compiler/internal/pointer"
	"github.com/cell-labs/cell-script/compiler/compiler/strings"
	"github.com/cell-labs/cell-script/compiler/compiler/types"
	"github.com/cell-labs/cell-script/compiler/compiler/value"
	"github.com/cell-labs/cell-script/compiler/parser"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	llvmTypes "github.com/llir/llvm/ir/types"
)

func (c *Compiler) compileConstantNode(v *parser.ConstantNode) value.Value {
	switch v.Type {
	case parser.BYTE:
		return value.Value{
			Value:      constant.NewInt(i8.Type, v.Value),
			Type:       i8,
			IsVariable: false,
		}
	case parser.NUMBER:
		// To make primitive have method table, there is a hacking way:
		// by defining the type with the same name: `type uint64 uint64`,
		// then new methods can be added to primitive types.
		// Here we try to simulate the hacking approach.
		// The default number type is uint64.
		var intType *types.Int = c.parserTypeToType(&parser.SingleTypeNode{
			SourceName: "uint64", TypeName: "uint64",
		}).(*types.Int)

		if v.TargetType != nil {
			intType = c.parserTypeToType(v.TargetType).(*types.Int)
		}
		// Use context to detect which type that should be returned
		// Is used to detect if a number should be i32 or i64 etc...
		// todo: better strategy
		// var wantedType types.Type
		// if len(c.contextAssignDest) > 0 {
		// 	wantedType = c.contextAssignDest[len(c.contextAssignDest)-1].Type
		// }

		// Create the correct type of int based on context
		// if t, ok := wantedType.(*types.Int); ok {
		// 	intType = t
		// }

		return value.Value{
			Value:      constant.NewInt(intType.Type, v.Value),
			Type:       intType,
			IsVariable: false,
		}

	case parser.STRING:
		var constString *ir.Global

		// Reuse the *ir.Global if it has already been created
		if reusedConst, ok := c.stringConstants[v.ValueStr]; ok {
			constString = reusedConst
		} else {
			constString = c.module.NewGlobalDef(strings.NextStringName(), strings.Constant(v.ValueStr))
			constString.Immutable = true
			c.stringConstants[v.ValueStr] = constString
		}

		sType, ok := c.packages["global"].GetPkgType("string", true)
		if !ok {
			panic("string type not found")
		}
		alloc := c.contextBlock.NewAlloca(sType.LLVM())

		// Save length of the string
		lenItem := c.contextBlock.NewGetElementPtr(pointer.ElemType(alloc), alloc, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, 0))
		c.contextBlock.NewStore(constant.NewInt(llvmTypes.I64, int64(len(v.ValueStr))), lenItem)

		// Save i8* version of string
		strItem := c.contextBlock.NewGetElementPtr(pointer.ElemType(alloc), alloc, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, 1))
		c.contextBlock.NewStore(strings.Toi8Ptr(c.contextBlock, constString), strItem)

		return value.Value{
			Value:      c.contextBlock.NewLoad(pointer.ElemType(alloc), alloc),
			Type:       types.String,
			IsVariable: false,
		}

	case parser.BOOL:
		return value.Value{
			// todo: optimise bool memory
			Value:      constant.NewInt(llvmTypes.I1, v.Value),
			Type:       types.Bool,
			IsVariable: false,
		}

	default:
		panic("Unknown constant Type")
	}
}
