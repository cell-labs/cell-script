package parser

import (
	"fmt"
	"strings"

	"github.com/cell-labs/cell-script/compiler/lexer"
)

// DefineFuncNode creates a new named function
type DefineFuncNode struct {
	baseNode

	Name    string
	IsNamed bool

	IsMethod bool
	IsExtern bool

	MethodOnType      *SingleTypeNode
	IsPointerReceiver bool
	InstanceName      string

	Arguments    []*NameNode
	ReturnValues []*NameNode
	Body         []Node
}

func (dfn DefineFuncNode) String() string {
	var body []string

	for _, b := range dfn.Body {
		body = append(body, fmt.Sprintf("%+v", b))
	}

	if dfn.IsMethod {
		return fmt.Sprintf("func+m (%+v) %s(%+v) %+v {\n\t%s\n}", dfn.InstanceName, dfn.Name, dfn.Arguments, dfn.ReturnValues, strings.Join(body, "\n\t"))
	} else if dfn.IsNamed {
		return fmt.Sprintf("func+n %s(%+v) %+v {\n\t%s\n}", dfn.Name, dfn.Arguments, dfn.ReturnValues, strings.Join(body, "\n\t"))
	} else {
		return fmt.Sprintf("func+v (%+v) %+v {\n\t%s\n}", dfn.Arguments, dfn.ReturnValues, strings.Join(body, "\n\t"))
	}
}

func (dfn DefineFuncNode) Mangling() string {
	if dfn.IsNamed {
		return fmt.Sprintf("%s%+v%+v", dfn.Name, dfn.Arguments, dfn.ReturnValues)
	} else {
		return fmt.Sprintf("%+v%+v", dfn.Arguments, dfn.ReturnValues)
	}
}

type ExternNode struct {
	baseNode
	FuncNodes []*DefineFuncNode
}

func (en ExternNode) String() string {
	var fns []string

	for _, f := range en.FuncNodes {
		fns = append(fns, fmt.Sprintf("%+v", f.String()))
	}

	return fmt.Sprintf("extern {\n\t%s\n}", strings.Join(fns, "\n\t"))
}

func (p *parser) parseExtern() *ExternNode {
	// Single extern statement
	expectFuncString := p.lookAhead(0)
	if expectFuncString.Type == lexer.KEYWORD && expectFuncString.Val == "func" {
		p.i++
		dfn := p.parseFuncDefinition()
		dfn.IsExtern = true
		return &ExternNode{
			FuncNodes: []*DefineFuncNode{dfn},
		}
	}

	fns := []*DefineFuncNode{}
	// Multiple extern
	p.expect(lexer.Item{Type: lexer.OPERATOR, Val: "("}, p.lookAhead(0))
	p.i++
	for {
		checkIfEndParen := p.lookAhead(0)
		if checkIfEndParen.Type == lexer.OPERATOR && checkIfEndParen.Val == ")" {
			break
		}
		if checkIfEndParen.Type == lexer.EOL {
			p.i++
			continue
		}

		if checkIfEndParen.Type == lexer.KEYWORD && checkIfEndParen.Val == "func" {
			p.i++
			fn := p.parseFuncDefinition()
			fn.IsExtern = true
			fns = append(fns, fn)
			continue
		}

		panic(fmt.Sprintf("Failed to parse extern: %+v", checkIfEndParen))
	}
	return &ExternNode{
		FuncNodes: fns,
	}
}

// The tokens after keyword "func"
func (p *parser) parseFuncDefinition() *DefineFuncNode {
	defineFunc := &DefineFuncNode{}
	var argsOrMethodType []*NameNode
	var canBeMethod bool

	checkIfOpeningParen := p.lookAhead(0)
	if checkIfOpeningParen.Type == lexer.OPERATOR && checkIfOpeningParen.Val == "(" {
		p.i++
		argsOrMethodType = p.parseFunctionArguments()
		canBeMethod = true
	}

	checkIfIdentifier := p.lookAhead(0)
	checkIfOpeningParen = p.lookAhead(1)

	if canBeMethod && checkIfIdentifier.Type == lexer.IDENTIFIER &&
		checkIfOpeningParen.Type == lexer.OPERATOR && checkIfOpeningParen.Val == "(" {

		defineFunc.IsMethod = true
		defineFunc.IsNamed = true
		defineFunc.Name = checkIfIdentifier.Val

		if len(argsOrMethodType) != 1 {
			panic("Unexpected count of types in method")
		}

		defineFunc.InstanceName = argsOrMethodType[0].Name

		methodOnType := argsOrMethodType[0].Type

		if pointerSingleTypeNode, ok := methodOnType.(*PointerTypeNode); ok {
			defineFunc.IsPointerReceiver = true
			methodOnType = pointerSingleTypeNode.ValueType
		}

		if singleTypeNode, ok := methodOnType.(*SingleTypeNode); ok {
			defineFunc.MethodOnType = singleTypeNode
		} else {
			panic(fmt.Sprintf("could not find type in method defitition: %T", methodOnType))
		}
	}

	name := p.lookAhead(0)
	openParen := p.lookAhead(1)
	if name.Type == lexer.IDENTIFIER && openParen.Type == lexer.OPERATOR && openParen.Val == "(" {
		defineFunc.Name = name.Val
		defineFunc.IsNamed = true

		p.i++
		p.i++

		// Parse argument list
		defineFunc.Arguments = p.parseFunctionArguments()
	} else {
		defineFunc.Arguments = argsOrMethodType
	}

	// Parse return types
	var retTypesNodeNames []*NameNode

	checkIfOpeningCurly := p.lookAhead(0)
	if checkIfOpeningCurly.Type != lexer.OPERATOR || checkIfOpeningCurly.Val != "{" {

		// Check if next is an opening parenthesis
		// Is optional if there's only one return type, is required
		// if there is multiple ones
		var allowMultiRetVals bool

		checkIfOpenParen := p.lookAhead(0)
		if checkIfOpenParen.Type == lexer.OPERATOR && checkIfOpenParen.Val == "(" {
			allowMultiRetVals = true
			p.i++
		}

		for {
			if checkIfOpenParen.Type == lexer.EOL || checkIfOpenParen.Type == lexer.EOF {
				break
			}

			nameNode := &NameNode{}

			// Support both named return values and when we only get the type
			retTypeOrNamed, err := p.parseOneType()
			if err != nil {
				panic(err)
			}
			p.i++

			// Next can be type, that means that the previous was the name of the var
			isType := p.lookAhead(0)
			if isType.Type == lexer.IDENTIFIER ||
				isType.Type == lexer.OPERATOR && isType.Val == "[" {
				retType, err := p.parseOneType()
				if err != nil {
					panic(err)
				}
				p.i++

				nameNode.Name = retTypeOrNamed.Type()
				nameNode.Type = retType
			} else {
				nameNode.Type = retTypeOrNamed
			}

			retTypesNodeNames = append(retTypesNodeNames, nameNode)

			if !allowMultiRetVals {
				break
			}

			// Check if comma or end parenthesis
			commaOrEnd := p.lookAhead(0)
			if commaOrEnd.Type == lexer.OPERATOR && commaOrEnd.Val == "," {
				p.i++
				continue
			}

			if commaOrEnd.Type == lexer.OPERATOR && commaOrEnd.Val == ")" {
				p.i++
				break
			}
			panic("Could not parse function return types")
		}
	}

	defineFunc.ReturnValues = retTypesNodeNames
	return defineFunc
}
