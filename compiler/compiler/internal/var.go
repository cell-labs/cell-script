package internal

import (
	"github.com/cell-labs/cell-script/compiler/compiler/internal/pointer"
	"github.com/cell-labs/cell-script/compiler/compiler/value"
	"github.com/llir/llvm/ir"
	llvmValue "github.com/llir/llvm/ir/value"
)

func LoadIfVariable(block *ir.Block, val value.Value) llvmValue.Value {
	if val.IsVariable {
		return block.NewLoad(pointer.ElemType(val.Value), val.Value)
	}
	return val.Value
}
