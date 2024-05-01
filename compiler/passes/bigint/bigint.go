package bigint

import (
	"fmt"
	"math/big"

	"github.com/cell-labs/cell-script/compiler/parser"
)

func BigInt(root *parser.FileNode) *parser.FileNode {
	parser.Walk(&bigIntVisitor{}, root)
	return root
}

type bigIntVisitor struct{}

func (b *bigIntVisitor) Visit(node parser.Node) (n parser.Node, v parser.Visitor) {
	v = b
	n = node

	// transform var a uint128 = 123456
	// to		 var a bigint = "123456"
	// transform var b uint128
	// to		 var b bigint
	// transform var c := a + b
	// to		 var c bigint = a + b
	if a, ok := n.(*parser.AllocNode); ok {
		fmt.Println(a)
		if a.Type == nil {

		} else if typeName := a.Type.Type(); typeName == "uint128" || typeName == "uint256" {
			if s, ok := a.Type.(*parser.SingleTypeNode); ok {
				s.TypeName = "bigint"
			}
		}
		if len(a.Val) > 0 {
			for _, v := range a.Val {
				if c, ok := v.(*parser.ConstantNode); ok {
					c.Type = parser.STRING
					c.ValueStr = c.Value.Text(10)
					c.Value = nil
				}
			}
		}
	}

	// transform a = 123456
	// to		 var a bigint = bigIntFromString("123456")
	// transform b = a
	// to		 b bigint = bigIntFromAssign(a)
	// transform c := a
	// to		 c bigint = bigIntFromAssign(a)
	if a, ok := node.(*parser.AssignNode); ok {
		if a.Target == nil {
			return
		}

		// TODO: check the target type is bigint
		if len(a.Val) == 0 {
			return
		} else if c, ok := a.Val[0].(*parser.ConstantNode); ok {
			a.Val = []parser.Node{
				&parser.CallNode{
					Function: &parser.NameNode{Package: "global", Name: "bigIntFromString"},
					Arguments: []parser.Node{
						&parser.ConstantNode{
							Type:     parser.STRING,
							Value:    c.Value,
							ValueStr: c.ValueStr,
						},
					},
				},
			}
		} else {

			a.Val = []parser.Node{
				&parser.CallNode{
					Function:  &parser.NameNode{Package: "global", Name: "bigIntAssign"},
					Arguments: a.Val,
				},
			}
		}
		return a, nil
	}

	// transform a == 12345
	// to		 bigIntEqual(a, 12345) == true
	if c, ok := node.(*parser.ConditionNode); ok {
		transformTo := func(funcName string) {
			l := c.Cond.Left
			r := c.Cond.Right
			c.Cond.Left = &parser.CallNode{
				Function:  &parser.NameNode{Package: "global", Name: funcName},
				Arguments: []parser.Node{l, r},
			}
			c.Cond.Right = &parser.ConstantNode{
				Type:  parser.BOOL,
				Value: big.NewInt(1), // set true
			}
		}
		if c.Cond.Operator == "==" {
			transformTo("bigIntEqual")
		} else if c.Cond.Operator == ">" {
			transformTo("bigIntGT")
		} else if c.Cond.Operator == ">=" {
			transformTo("bigIntGTE")
		} else if c.Cond.Operator == "<" {
			transformTo("bigIntLT")
		} else if c.Cond.Operator == "<=" {
			transformTo("bigIntLTE")
		}
	}

	// transform sum := a + b
	// to		 sum = bigIntAdd(a, b)
	if c, ok := node.(*parser.OperatorNode); ok {
		transformTo := func(funcName string) {
			l := c.Left
			r := c.Right
			n = &parser.CallNode{
				Function:  &parser.NameNode{Package: "global", Name: funcName},
				Arguments: []parser.Node{l, r},
			}
		}
		switch c.Operator {
		case "+":
			transformTo("bigIntAdd")
		case "-":
			transformTo("bigIntSub")
		case "*":
			transformTo("bigIntMul")
		case "/":
			transformTo("bigIntDiv")
		case "%":
			transformTo("bigIntMod")
		}
	}

	// transform print(a)
	// to		 sum = bigIntPrint(a, b)
	return
}
