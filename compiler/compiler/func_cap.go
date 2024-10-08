package compiler

import (
	"strings"

	"github.com/cell-labs/cell-script/compiler/compiler/internal/pointer"
	"github.com/cell-labs/cell-script/compiler/compiler/value"
	"github.com/cell-labs/cell-script/compiler/parser"
)

func (c *Compiler) capFuncCall(v *parser.CallNode) value.Value {
	arg := c.compileValue(v.Arguments[0])

	if strings.HasPrefix(arg.Type.Name(), "slice") {
		val := arg.Value
		val = c.contextBlock.NewLoad(pointer.ElemType(val), val)

		return value.Value{
			Value:      c.contextBlock.NewExtractValue(val, 1),
			Type:       i64,
			IsVariable: false,
		}
	}

	c.panic(c.contextBlock, "Can not call cap on "+arg.Type.Name())
	return value.Value{}
}
