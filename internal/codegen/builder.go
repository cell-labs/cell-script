package codegen

import (
	"github.com/llir/llvm/ir"
)

type IrBuilder struct {
	module *ir.Module
}

func NewIrBuilder() *IrBuilder {
	return &IrBuilder{}
}
