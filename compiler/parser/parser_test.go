package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cell-labs/cell-script/compiler/lexer"
	"github.com/cell-labs/cell-script/compiler/option"
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
				Arguments: []Node{&ConstantNode{Type: NUMBER, Value: 1}},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, &option.Options{Debug: false}))
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
					Value: 1,
				},
				Right: &ConstantNode{
					Type:  NUMBER,
					Value: 2,
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, &option.Options{Debug: false}))
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
					Value: 1,
				},
				Right: &OperatorNode{
					Operator: OP_MUL,
					Left: &ConstantNode{
						Type:  NUMBER,
						Value: 2,
					},
					Right: &ConstantNode{
						Type:  NUMBER,
						Value: 3,
					},
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, &option.Options{Debug: false}))
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
						Value: 1,
					},
					Right: &ConstantNode{
						Type:  NUMBER,
						Value: 2,
					},
				},
				Right: &ConstantNode{
					Type:  NUMBER,
					Value: 3,
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, &option.Options{Debug: false}))
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
						Value: 1,
					},
					Right: &ConstantNode{
						Type:  NUMBER,
						Value: 2,
					},
				},
				Right: &OperatorNode{
					Operator: OP_MUL,
					Left: &ConstantNode{
						Type:  NUMBER,
						Value: 3,
					},
					Right: &ConstantNode{
						Type:  NUMBER,
						Value: 4,
					},
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, &option.Options{Debug: false}))
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
						Left:     &ConstantNode{Type: NUMBER, Value: 100},
						Right:    &ConstantNode{Type: NUMBER, Value: 3},
					},
					Right: &ConstantNode{Type: NUMBER, Value: 4},
				},

				Right: &ConstantNode{Type: NUMBER, Value: 7},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, &option.Options{Debug: false}))
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
						Left:     &ConstantNode{Type: NUMBER, Value: 100},
						Right: &StructLoadElementNode{
							Struct:      &NameNode{Name: "f"},
							ElementName: "a",
						},
					},
					Right: &ConstantNode{Type: NUMBER, Value: 4},
				},

				Right: &ConstantNode{Type: NUMBER, Value: 7},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, &option.Options{Debug: false}))
}

func TestPramga(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.KEYWORD, Val: "pragma", Line: 1},
		{Type: lexer.IDENTIFIER, Val: "cellscript", Line: 1},
		{Type: lexer.NUMBER, Val: "0", Line: 1},
		{Type: lexer.NUMBER, Val: "0", Line: 1},
		{Type: lexer.NUMBER, Val: "1", Line: 1},
		{Type: lexer.EOL},
		{Type: lexer.EOF},
	}

	/*
		pramga cellscript 0.0.1
	*/

	expected := &FileNode{
		Instructions: []Node{
			&PragmaNode{
				Version: VersionScheme{
					Major: 0,
					Minor: 0,
					Patch: 1,
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, &option.Options{Debug: false}))
}

func TestFunction(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.KEYWORD, Val: "extern", Line: 1},
		{Type: lexer.OPERATOR, Val: "(", Line: 1},
		{Type: lexer.KEYWORD, Val: "func", Line: 1},
		{Type: lexer.IDENTIFIER, Val: "foo", Line: 1},
		{Type: lexer.OPERATOR, Val: "(", Line: 1},
		{Type: lexer.OPERATOR, Val: ")", Line: 1},
		{Type: lexer.IDENTIFIER, Val: "int32", Line: 1},
		{Type: lexer.OPERATOR, Val: ")", Line: 1},
		{Type: lexer.EOL},
		{Type: lexer.EOF},
	}

	/*
		extern (
			func foo() int32
			func bar() int32
		)
		extern func foo() int32
	*/

	expected := &FileNode{
		Instructions: []Node{
			&ExternNode{
				FuncNodes: []*DefineFuncNode{
					&DefineFuncNode{
						Name:     "foo",
						IsNamed:  true,
						IsMethod: false,
						IsExtern: true,

						MethodOnType:      nil,
						IsPointerReceiver: false,
						InstanceName:      "",

						Arguments: nil,
						ReturnValues: []*NameNode{
							&NameNode{
								Type: &SingleTypeNode{
									TypeName: "int32",
								},
							},
						},
						Body: nil,
					},
				},
			},
		},
	}

	assert.Equal(t, expected, Parse(input, &option.Options{Debug: false}))
}

func TestCondition(t *testing.T) {
	input := []lexer.Item{
		{Type: lexer.IDENTIFIER, Val: "a", Line: 1},
		{Type: lexer.OPERATOR, Val: ">", Line: 1},
		{Type: lexer.IDENTIFIER, Val: "b", Line: 1},
		{Type: lexer.OPERATOR, Val: "&&", Line: 1},
		{Type: lexer.IDENTIFIER, Val: "b", Line: 1},
		{Type: lexer.OPERATOR, Val: ">", Line: 1},
		{Type: lexer.IDENTIFIER, Val: "c", Line: 1},
		{Type: lexer.EOL},
		{Type: lexer.EOF},
	}

	/*
		a > b && b > c
	*/

	expected := &FileNode{
		Instructions: []Node{
			&OperatorNode{
				Operator: "&&",
				Left: &OperatorNode{
					Operator: ">",
					Left: &NameNode{Name: "a"},
					Right: &NameNode{Name: "b"},
				},
				Right: &OperatorNode{
					Operator: ">",
					Left: &NameNode{Name: "b"},
					Right: &NameNode{Name: "c"},
				},
			},
		},
	}

	/*
		a + d > 2 && b == false || c < 3
	*/

	input = []lexer.Item{
		{Type: lexer.IDENTIFIER, Val: "a", Line: 1},
		{Type: lexer.OPERATOR, Val: "+", Line: 1},
		{Type: lexer.IDENTIFIER, Val: "d", Line: 1},
		{Type: lexer.OPERATOR, Val: ">", Line: 1},
		{Type: lexer.NUMBER, Val: "2", Line: 1},
		{Type: lexer.OPERATOR, Val: "&&", Line: 1},
		{Type: lexer.IDENTIFIER, Val: "b", Line: 1},
		{Type: lexer.OPERATOR, Val: "==", Line: 1},
		{Type: lexer.KEYWORD, Val: "false", Line: 1},
		{Type: lexer.OPERATOR, Val: "||", Line: 1},
		{Type: lexer.IDENTIFIER, Val: "c", Line: 1},
		{Type: lexer.OPERATOR, Val: "<", Line: 1},
		{Type: lexer.NUMBER, Val: "3", Line: 1},
		{Type: lexer.EOL},
		{Type: lexer.EOF},
	}
              
	expected = &FileNode{
		Instructions: []Node{
			&OperatorNode{
				Operator: "&&",
				Left: &OperatorNode{
					Operator: ">",
					Left: &OperatorNode{
						Operator: "+",
						Left: &NameNode{Name: "a"},
						Right: &NameNode{Name: "d"},
					},
					Right: &ConstantNode{Type: NUMBER, Value: 2},
				},
				Right: &OperatorNode{
					Operator: "||",
					Left: &OperatorNode{
						Operator: "==",
						Left: &NameNode{Name: "b"},
						Right: &ConstantNode{Type: BOOL, Value: 0},
					},
					Right: &OperatorNode{
						Operator: "<",
						Left: &NameNode{Name: "c"},
						Right: &ConstantNode{Type: NUMBER, Value: 3},
					},
				},
			},
		},
	}
	
	actual := Parse(input, &option.Options{Debug: false})
	assert.Equal(t, expected, actual)
}
