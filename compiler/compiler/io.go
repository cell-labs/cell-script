package compiler

import (
	"github.com/cell-labs/cell-script/compiler/compiler/syscall"
	"github.com/cell-labs/cell-script/compiler/compiler/types"
	"github.com/cell-labs/cell-script/compiler/compiler/value"
	"github.com/cell-labs/cell-script/compiler/parser"
)

func (c *Compiler) printFuncCall(v *parser.CallNode) value.Value {
	arg := c.compileValue(v.Arguments[0])
	syscall.Print(c.contextBlock, arg, c.GOOS)
	return value.Value{Type: types.Void}
}
