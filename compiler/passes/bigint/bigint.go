package bigint

import (
	"fmt"
	"math/big"

	"github.com/cell-labs/cell-script/compiler/parser"
	"github.com/cell-labs/cell-script/compiler/passes/intrinsic"
)

func BigInt(root *parser.FileNode) *parser.FileNode {
	parser.Walk(&bigIntVisitor{}, root)
	return root
}

type bigIntVisitor struct{}

func (b *bigIntVisitor) Visit(node parser.Node) (n parser.Node, v parser.Visitor) {
	v = b
	n = node
	var bigIntNode = intrisinc.GetTypeNodeByName("bigint")

	if _, ok := node.(*parser.StructTypeNode); ok {
		fmt.Printf("%#v\n", node)
	}

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

		} else if a.Type.Type() == "uint128" || a.Type.Type() == "uint256" {
			a.Type.SetName("bigint")
			a.Val = []parser.Node{
				&parser.CallNode{
					Function: &parser.DefineFuncNode{
						Name:          "big_int_new",
						IsNamed:       true,
						IsCompilerAdd: true,
					},
					Arguments: a.Val,
				},
			}
			fmt.Println(a)
			return a, nil
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
		} else if _, ok := a.Val[0].(*parser.ConstantNode); ok {
			a.Val = []parser.Node{
				&parser.CallNode{
					Function: &parser.DefineFuncNode{
						Name:          "big_int_from_string",
						IsNamed:       true,
						IsCompilerAdd: true,
						ReturnValues: []*parser.NameNode{
							{
								Name: "bigint",
								Type: bigIntNode,
							},
						},
					},
					Arguments: a.Val,
				},
			}
		} else {

			a.Val = []parser.Node{
				&parser.CallNode{
					Function: &parser.DefineFuncNode{
						Name:          "big_int_assign",
						IsNamed:       true,
						IsCompilerAdd: true,
					},
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
				Function: &parser.DefineFuncNode{
					Name:          funcName,
					IsNamed:       true,
					IsCompilerAdd: true,
					ReturnValues:  []*parser.NameNode{&parser.NameNode{Type: parser.SingleTypeNode{TypeName: "bool"}}},
				},
				Arguments: []parser.Node{l, r},
			}
			c.Cond.Right = &parser.ConstantNode{
				Type:  parser.BOOL,
				Value: big.NewInt(1), // set true
			}
		}
		if c.Cond.Operator == "==" {
			transformTo("big_int_equal")
		} else if c.Cond.Operator == ">" {
			transformTo("big_int_gt")
		} else if c.Cond.Operator == ">=" {
			transformTo("big_int_gte")
		} else if c.Cond.Operator == "<" {
			transformTo("big_int_lt")
		} else if c.Cond.Operator == "<=" {
			transformTo("big_int_lte")
		}
	}
	return
}
