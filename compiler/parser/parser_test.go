package parser

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cell-labs/cell-script/compiler/lexer"
)

func TestCall(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.IDENTIFIER, Val: "printf"},
		{Type: lexer.OPERATOR, Val: "("},
		{Type: lexer.NUMBER, Val: "1"},
		{Type: lexer.OPERATOR, Val: ")"},
		{Type: lexer.EOF, Val: ""},
	}

	expected := &FileNode{
		Instructions: []Node{
			&CallNode{
				Function:  &NameNode{Name: "printf"},
				Arguments: []Node{&ConstantNode{Type: NUMBER, Value: big.NewInt(1)}},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, false))
}

func TestAdd(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.NUMBER, Val: "1"},
		{Type: lexer.OPERATOR, Val: "+"},
		{Type: lexer.NUMBER, Val: "2"},
		{Type: lexer.EOF, Val: ""},
	}

	expected := &FileNode{
		Instructions: []Node{
			&OperatorNode{
				Operator: OP_ADD,
				Left: &ConstantNode{
					Type:  NUMBER,
					Value: big.NewInt(1),
				},
				Right: &ConstantNode{
					Type:  NUMBER,
					Value: big.NewInt(2),
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, false))
}

func TestInfixPriority(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.NUMBER, Val: "1"},
		{Type: lexer.OPERATOR, Val: "+"},
		{Type: lexer.NUMBER, Val: "2"},
		{Type: lexer.OPERATOR, Val: "*"},
		{Type: lexer.NUMBER, Val: "3"},
		{Type: lexer.EOF, Val: ""},
	}

	expected := &FileNode{
		Instructions: []Node{
			&OperatorNode{
				Operator: OP_ADD,
				Left: &ConstantNode{
					Type:  NUMBER,
					Value: big.NewInt(1),
				},
				Right: &OperatorNode{
					Operator: OP_MUL,
					Left: &ConstantNode{
						Type:  NUMBER,
						Value: big.NewInt(2),
					},
					Right: &ConstantNode{
						Type:  NUMBER,
						Value: big.NewInt(3),
					},
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, false))
}

func TestInfixPriority2(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.NUMBER, Val: "1"},
		{Type: lexer.OPERATOR, Val: "*"},
		{Type: lexer.NUMBER, Val: "2"},
		{Type: lexer.OPERATOR, Val: "+"},
		{Type: lexer.NUMBER, Val: "3"},
		{Type: lexer.EOF, Val: ""},
	}

	expected := &FileNode{
		Instructions: []Node{
			&OperatorNode{
				Operator: OP_ADD,
				Left: &OperatorNode{
					Operator: OP_MUL,
					Left: &ConstantNode{
						Type:  NUMBER,
						Value: big.NewInt(1),
					},
					Right: &ConstantNode{
						Type:  NUMBER,
						Value: big.NewInt(2),
					},
				},
				Right: &ConstantNode{
					Type:  NUMBER,
					Value: big.NewInt(3),
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, false))
}

func TestInfixPriority3(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.NUMBER, Val: "1"},
		{Type: lexer.OPERATOR, Val: "*"},
		{Type: lexer.NUMBER, Val: "2"},
		{Type: lexer.OPERATOR, Val: "+"},
		{Type: lexer.NUMBER, Val: "3"},
		{Type: lexer.OPERATOR, Val: "*"},
		{Type: lexer.NUMBER, Val: "4"},
		{Type: lexer.EOF, Val: ""},
	}

	expected := &FileNode{
		Instructions: []Node{
			&OperatorNode{
				Operator: OP_ADD,
				Left: &OperatorNode{
					Operator: OP_MUL,
					Left: &ConstantNode{
						Type:  NUMBER,
						Value: big.NewInt(1),
					},
					Right: &ConstantNode{
						Type:  NUMBER,
						Value: big.NewInt(2),
					},
				},
				Right: &OperatorNode{
					Operator: OP_MUL,
					Left: &ConstantNode{
						Type:  NUMBER,
						Value: big.NewInt(3),
					},
					Right: &ConstantNode{
						Type:  NUMBER,
						Value: big.NewInt(4),
					},
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, false))
}

func TestInfixPriority4(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.NUMBER, Val: "100"},
		{Type: lexer.OPERATOR, Val: "/"},
		{Type: lexer.NUMBER, Val: "3"},
		{Type: lexer.OPERATOR, Val: "/"},
		{Type: lexer.NUMBER, Val: "4"},
		{Type: lexer.OPERATOR, Val: "*"},
		{Type: lexer.NUMBER, Val: "7"},
		{Type: lexer.EOF, Val: ""},
	}

	/*
		OP(OP(OP(100/300)/4) * 7)

	*/

	expected := &FileNode{
		Instructions: []Node{
			&OperatorNode{
				Operator: OP_MUL,

				Left: &OperatorNode{
					Operator: OP_DIV,
					Left: &OperatorNode{
						Operator: OP_DIV,
						Left:     &ConstantNode{Type: NUMBER, Value: big.NewInt(100)},
						Right:    &ConstantNode{Type: NUMBER, Value: big.NewInt(3)},
					},
					Right: &ConstantNode{Type: NUMBER, Value: big.NewInt(4)},
				},

				Right: &ConstantNode{Type: NUMBER, Value: big.NewInt(7)},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, false))
}

func TestInfixPriority4Load(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.NUMBER, Val: "100"},
		{Type: lexer.OPERATOR, Val: "/"},
		{Type: lexer.IDENTIFIER, Val: "f"},
		{Type: lexer.OPERATOR, Val: "."},
		{Type: lexer.IDENTIFIER, Val: "a"},
		{Type: lexer.OPERATOR, Val: "/"},
		{Type: lexer.NUMBER, Val: "4"},
		{Type: lexer.OPERATOR, Val: "*"},
		{Type: lexer.NUMBER, Val: "7"},
		{Type: lexer.EOF, Val: ""},
	}

	/*
		OP(OP(OP(100/f.a)/4) * 7)
	*/

	expected := &FileNode{
		Instructions: []Node{
			&OperatorNode{
				Operator: OP_MUL,

				Left: &OperatorNode{
					Operator: OP_DIV,
					Left: &OperatorNode{
						Operator: OP_DIV,
						Left:     &ConstantNode{Type: NUMBER, Value: big.NewInt(100)},
						Right: &StructLoadElementNode{
							Struct:      &NameNode{Name: "f"},
							ElementName: "a",
						},
					},
					Right: &ConstantNode{Type: NUMBER, Value: big.NewInt(4)},
				},

				Right: &ConstantNode{Type: NUMBER, Value: big.NewInt(7)},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, false))
}

func TestAllocConstantWithSuffix(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.IDENTIFIER, Val: "a"},
		{Type: lexer.OPERATOR, Val: ":="},
		{Type: lexer.BIGNUMBER, Val: "3"},
		{Type: lexer.EOF},
	}

	expected := &FileNode{
		Instructions: []Node{
			&AllocNode{
				Escapes: false,
				Name: []string{"a"},
				Val: []Node{
					&ConstantNode{
						Type: BIGNUMBER,
						Value: big.NewInt(3),
						ValueStr: "3",
					},
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, false))
}

func TestAssignConstantWithSuffix(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.IDENTIFIER, Val: "a"},
		{Type: lexer.OPERATOR, Val: "="},
		{Type: lexer.BIGNUMBER, Val: "3", Suffix: "u128"},
		{Type: lexer.EOF},
	}

	expected := &FileNode{
		Instructions: []Node{
			&AssignNode{
				Target: []Node{
					&NameNode{
						Name: "a",
					},
				},
				Val: []Node{
					&ConstantNode{
						Type: BIGNUMBER,
						Value: big.NewInt(3),
						ValueStr: "3",
					},
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, false))
}
