package compiler

import (
	"github.com/llir/llvm/ir/constant"
	llvmTypes "github.com/llir/llvm/ir/types"
	llvmValue "github.com/llir/llvm/ir/value"

	"github.com/cell-labs/cell-script/compiler/compiler/internal"
	"github.com/cell-labs/cell-script/compiler/compiler/internal/pointer"
	"github.com/cell-labs/cell-script/compiler/compiler/name"
	"github.com/cell-labs/cell-script/compiler/compiler/value"
	"github.com/cell-labs/cell-script/compiler/parser"
)

func (c *Compiler) compileInitializeStringWithSliceNode(v *parser.InitializeStringWithSliceNode) value.Value {
	sliceValue := c.compileValue(v.Items[0])
	srcVal := internal.LoadIfVariable(c.contextBlock, sliceValue)
	srcLen := c.contextBlock.NewExtractValue(srcVal, 0)
	srcOff := c.contextBlock.NewExtractValue(srcVal, 2)
	srcArr := c.contextBlock.NewExtractValue(srcVal, 3)
	srcArrStartPtr := c.contextBlock.NewGetElementPtr(llvmTypes.I8, srcArr, srcOff)
	// create new string
	var len64 llvmValue.Value
	len64 = srcLen
	if srcLen.Type() != llvmTypes.I64 {
		len64 = c.contextBlock.NewZExt(srcLen, i64.LLVM())
	}
	// todo: c.contextBlock.NewCall(c.osFuncs.Strndup.Value.(llvmValue.Named), srcArrStartPtr, len64)
	// construct a new string {i64, i8*}
	strVal := srcArrStartPtr
	sType, ok := c.packages["global"].GetPkgType("string", true)
	if !ok {
		panic("string type not found")
	}
	alloc := c.contextBlock.NewAlloca(sType.LLVM())

	// Save length of the string
	lenItem := c.contextBlock.NewGetElementPtr(pointer.ElemType(alloc), alloc, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, 0))
	lenItem.SetName(name.Var("len"))
	c.contextBlock.NewStore(len64, lenItem)

	// Save i8* version of string
	strItem := c.contextBlock.NewGetElementPtr(pointer.ElemType(alloc), alloc, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, 1))
	strItem.SetName(name.Var("str"))
	c.contextBlock.NewStore(strVal, strItem)

	return value.Value{
		Value:      c.contextBlock.NewLoad(pointer.ElemType(alloc), alloc),
		Type:       sType,
		IsVariable: false,
	}
}
