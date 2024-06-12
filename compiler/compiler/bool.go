package compiler

import (
	"github.com/cell-labs/cell-script/compiler/compiler/internal/pointer"
	"github.com/cell-labs/cell-script/compiler/compiler/types"
	"github.com/cell-labs/cell-script/compiler/compiler/value"
	"github.com/cell-labs/cell-script/compiler/parser"
	"github.com/llir/llvm/ir/constant"
	llvmTypes "github.com/llir/llvm/ir/types"
)

func (c *Compiler) compileNegateBoolNode(v *parser.NegateNode) value.Value {
	val := c.compileValue(v.Item)
	loadedVal := c.contextBlock.NewLoad(pointer.ElemType(val.Value), val.Value)

	return value.Value{
		Type:       types.Bool,
		Value:      c.contextBlock.NewXor(loadedVal, constant.NewInt(llvmTypes.I8, 1)),
		IsVariable: false,
	}
}
