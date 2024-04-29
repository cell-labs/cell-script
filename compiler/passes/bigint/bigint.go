package bigint

import (
	"math/big"

	"github.com/cell-labs/cell-script/compiler/parser"
	"github.com/cell-labs/cell-script/compiler/utils"
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
	// to		 var a bigint = bigIntFromString("123456")
	// transform var b uint128
	// to		 var b bigint = bigIntNew()
	// transform var c := a + b
	// to		 var c bigint = bigIntAdd()
	if a, ok := node.(*parser.AllocNode); ok {
		// fmt.Println(a)
		// fmt.Println(a.Val)
		if a.Type == nil {

		} else if a.Type.Type() == "bigint" {
			a.Type.SetName("bigint")
			if len(a.Val) > 0 {
				c, ok := a.Val[0].(*parser.ConstantNode)
				if !ok {
					utils.Ice("type error")
				}
				a.Val = []parser.Node{
					&parser.CallNode{
						Function: &parser.NameNode{Package: "global", Name:"bigIntFromString"},
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
						Function: &parser.NameNode{Package: "global", Name:"bigIntNew"},
						Arguments: []parser.Node{},
					},
				}
			}
			return a, nil
		}

	}

	// parser.ConstantNode
	if a, ok := node.(*parser.ConstantNode); ok {
		if a.Type == parser.BIGNUMBER {
			// fmt.Println(a)
			return &parser.CallNode{
				Function: &parser.NameNode{Package: "global", Name:"bigIntFromString"},
				Arguments: []parser.Node{
					&parser.ConstantNode{
						Type:     parser.STRING,
						Value:    a.Value,
						ValueStr: a.ValueStr,
					},
				},
			}, nil
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
					Function:  &parser.NameNode{Package: "global", Name:"bigIntFromString"},
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
					Function:  &parser.NameNode{Package: "global", Name:"bigIntAssign"},
					Arguments: a.Val,
				},
			}
		}
		return a, nil
	}

	// transform a == 12345
	// to		 big_int_equal(a, 12345) == true
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
	return
}
