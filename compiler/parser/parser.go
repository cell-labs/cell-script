package parser

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"errors"

	"github.com/cell-labs/cell-script/compiler/lexer"
	"github.com/cell-labs/cell-script/compiler/option"
)

type parser struct {
	i     int
	input []lexer.Item

	inAllocRightHand bool

	debug bool

	types map[string]struct{}

	// names of packages imported in this file
	// used to detect the difference in "a.b" where a is a package or a struct
	packages map[string]struct{}
}

func Parse(input []lexer.Item, options *option.Options) *FileNode {
	p := &parser{
		i:        0,
		input:    input,
		debug:    options.Debug,
		packages: map[string]struct{}{},
		// todo: loading builtin types all in once
		types: map[string]struct{}{
			"int":     {},
			"uint":    {},
			"int8":    {},
			"uint8":   {},
			"int16":   {},
			"uint16":  {},
			"int32":   {},
			"uint32":  {},
			"int64":   {},
			"uint64":  {},
			"uint128": {},
			"uint256": {},
			"uintptr": {},
			"byte":    {},
			"string":  {},
			"error":   {},
		},
	}

	return &FileNode{
		Instructions: p.parseUntil(lexer.Item{Type: lexer.EOF}),
	}
}

func (p *parser) parseOne(withAheadParse bool) (res Node) {
	return p.parseOneWithOptions(withAheadParse, withAheadParse, withAheadParse)
}

func (p *parser) parseNumberSuffix() (ty TypeNode) {
	next := p.lookAhead(1)
	if next.Type == lexer.IDENTIFIER {
		switch next.Val {
		case "u8":
			ty = &SingleTypeNode{SourceName: "uint8", TypeName: "uint8"}
			p.i++
		case "u16":
			ty = &SingleTypeNode{SourceName: "uint16", TypeName: "uint16"}
			p.i++
		case "u32":
			ty = &SingleTypeNode{SourceName: "uint32", TypeName: "uint32"}
			p.i++
		case "u64":
			ty = &SingleTypeNode{SourceName: "uint64", TypeName: "uint64"}
			p.i++
		case "u128":
			ty = &SingleTypeNode{SourceName: "uint128", TypeName: "uint128"}
			p.i++
		case "u256":
			ty = &SingleTypeNode{SourceName: "uint256", TypeName: "uint256"}
			p.i++
		}
	}
	return
}

func (p *parser) parseOneWithOptions(withAheadParse, withArithAhead, withIdentifierAhead bool) (res Node) {
	current := p.input[p.i]

	switch current.Type {

	case lexer.EOF:
		panic("unexpected EOF")

	case lexer.EOL:
		// Ignore the EOL, continue further
		// p.i++
		// return p.parseOne()
		return nil

	// IDENTIFIERS are converted to either:
	// - a CallNode if followed by an opening parenthesis (a function call), or
	// - a NodeName (variables)
	case lexer.IDENTIFIER:
		res = &NameNode{Name: current.Val}
		next := p.lookAhead(1)
		// array or slice may load element
		isIndex := next.Type == lexer.OPERATOR && next.Val == "["
		isPackageMember := next.Type == lexer.OPERATOR && next.Val == "."
		if withIdentifierAhead || isIndex || isPackageMember {
			res = p.aheadParseWithOptions(res, withArithAhead, withIdentifierAhead)
		}
		return


	// HEX always returns a ConstantNode
	// Convert string hex representation to int64
	case lexer.HEX:
		val, err := strconv.ParseInt(current.Val, 16, 64)
		if err != nil {
			panic(err)
		}
		res = &ConstantNode{
			Type:       NUMBER,
			TargetType: p.parseNumberSuffix(),
			Value:      val,
		}
		if withAheadParse {
			res = p.aheadParse(res)
		}
		return

	// NUMBER always returns a ConstantNode
	// Convert string representation to int64
	case lexer.NUMBER:
		val, err := strconv.ParseInt(current.Val, 10, 64)
		if err != nil {
			panic(err)
		}

		res = &ConstantNode{
			Type:       NUMBER,
			TargetType: p.parseNumberSuffix(),
			Value:      val,
		}
		if withAheadParse {
			res = p.aheadParse(res)
		}
		return

	// STRING is always a ConstantNode, the value is not modified
	case lexer.STRING:
		res = &ConstantNode{
			Type:     STRING,
			ValueStr: current.Val,
		}
		if withAheadParse {
			res = p.aheadParse(res)
		}
		return
	case lexer.BYTE:
		res = &ConstantNode{
			Type:  BYTE,
			Value: int64(current.Val[0]),
		}
		if withAheadParse {
			res = p.aheadParse(res)
		}
		return

	case lexer.OPERATOR:
		if current.Val == "&" {
			p.i++
			res = &GetReferenceNode{Item: p.parseOne(true)}
			if withAheadParse {
				res = p.aheadParse(res)
			}
			return
		}
		if current.Val == "*" {
			p.i++
			res = &DereferenceNode{Item: p.parseOne(false)}
			if withAheadParse {
				res = p.aheadParse(res)
			}
			return
		}

		if current.Val == "!" {
			p.i++
			res = &NegateNode{Item: p.parseOne(false)}
			return
		}

		if current.Val == "-" {
			p.i++
			res = p.aheadParse(&SubNode{Item: p.parseOne(true)})
			return
		}

		// Slice or array initalization
		if current.Val == "[" {
			next := p.lookAhead(1)

			// Slice init
			if next.Type == lexer.OPERATOR && next.Val == "]" {
				p.i += 2

				sliceItemType, err := p.parseOneType()
				if err != nil {
					panic(err)
				}

				p.i++

				next = p.lookAhead(0)
				if next.Type != lexer.OPERATOR || next.Val != "{" {
					log.Printf("%+v", next)
					panic("expected { after type in slice init")
				}

				p.i++

				prevInAlloc := p.inAllocRightHand
				p.inAllocRightHand = false
				items := p.parseUntil(lexer.Item{Type: lexer.OPERATOR, Val: "}"})
				p.inAllocRightHand = prevInAlloc

				len := &ConstantNode{
					Type:  NUMBER,
					Value: int64(len(items)),
				}
				for _, i := range items {
					if ii, ok := i.(*ConstantNode); ok {
						ii.TargetType = sliceItemType
					}
				}
				res = &InitializeSliceNode{
					Type:  sliceItemType,
					Len:   len,
					Items: items,
				}
				if withAheadParse {
					res = p.aheadParse(res)
				}
				return
			}

			p.i++

			// TODO: Support for compile-time artimethic ("[1+2]int{1,2,3}")
			arraySize := p.lookAhead(0)
			if arraySize.Type != lexer.NUMBER {
				panic("expected number in array size")
			}
			size, err := strconv.Atoi(arraySize.Val)
			if err != nil {
				panic("expected number in array size")
			}

			p.i++
			p.expect(p.lookAhead(0), lexer.Item{
				Type: lexer.OPERATOR,
				Val:  "]",
			})

			p.i++
			arrayItemType, err := p.parseOneType()
			if err != nil {
				panic(err)
			}

			p.i++
			p.expect(p.lookAhead(0), lexer.Item{
				Type: lexer.OPERATOR,
				Val:  "{",
			})

			p.i++

			prevInAlloc := p.inAllocRightHand
			p.inAllocRightHand = false
			items := p.parseUntil(lexer.Item{Type: lexer.OPERATOR, Val: "}"})
			p.inAllocRightHand = prevInAlloc

			for _, i := range items {
				if ii, ok := i.(*ConstantNode); ok {
					ii.TargetType = arrayItemType
				}
			}
			// Array init
			res = &InitializeArrayNode{
				Type:  arrayItemType,
				Size:  size,
				Items: items,
			}
			if withAheadParse {
				res = p.aheadParse(res)
			}
			return
		}

		if current.Val == "(" {
			p.i++
			i := p.parseUntil(lexer.Item{Type: lexer.OPERATOR, Val: ")"})
			if len(i) != 1 {
				panic("Expected exactly one item in GroupNode '('")
			}
			return p.aheadParseWithOptions(&GroupNode{Item: i[0]}, true, false)
		}

	case lexer.KEYWORD:

		// "if" gets converted to a ConditionNode
		// the keyword "if" is followed by
		// - a condition
		// - an opening curly bracket ({)
		// - a body
		// - a closing bracket (})
		if current.Val == "if" {

			getCondition := func() *OperatorNode {
				p.i++

				condNodes := p.parseUntil(lexer.Item{Type: lexer.OPERATOR, Val: "{"})
				if len(condNodes) != 1 {
					panic("could not parse if-condition")
				}

				p.i++

				if cond, ok := condNodes[0].(*OperatorNode); ok {
					return cond
				}

				// Add implicit == true
				return &OperatorNode{
					Left: condNodes[0],
					Right: &ConstantNode{
						Type:  BOOL,
						Value: 1,
					},
					Operator: OP_EQ,
				}

			}

			outerConditionNode := &ConditionNode{
				Cond: getCondition(),
				True: p.parseUntil(lexer.Item{Type: lexer.OPERATOR, Val: "}"}),
			}

			lastConditionNode := outerConditionNode

			p.i++

			// Check if the next keyword is "if" + "else" or "else"
			for {
				checkIfElse := p.lookAhead(0)
				if checkIfElse.Type != lexer.KEYWORD || checkIfElse.Val != "else" {
					break
				}

				p.i++

				checkIfElseIf := p.lookAhead(0)
				if checkIfElseIf.Type == lexer.KEYWORD && checkIfElseIf.Val == "if" {

					newCondNode := &ConditionNode{
						Cond: getCondition(),
						True: p.parseUntil(lexer.Item{Type: lexer.OPERATOR, Val: "}"}),
					}

					lastConditionNode.False = []Node{newCondNode}
					lastConditionNode = newCondNode
					p.i++
					continue
				}

				expectOpenBrack := p.lookAhead(0)
				if expectOpenBrack.Type != lexer.OPERATOR || expectOpenBrack.Val != "{" {
					panic("Expected { after else")
				}

				p.i++
				lastConditionNode.False = p.parseUntil(lexer.Item{Type: lexer.OPERATOR, Val: "}"})
			}

			return outerConditionNode
		}

		if current.Val == "panic" {
			p.i++
			lParent := p.lookAhead(0)
			p.expect(lParent, lexer.Item{Type: lexer.OPERATOR, Val: "("})
			p.i++
			args := p.parseUntil(lexer.Item{Type: lexer.OPERATOR, Val: ")"})
			if len(args) != 1 {
				panic("wrong number of arguments for panic(message)")
			}
			if _, ok := args[0].(*ConstantNode); !ok {
				panic("wrong type of argument for panic(message)")
			}
			return &CallNode{
				Function:  &NameNode{Name: "panic"},
				Arguments: args,
			}
		}

		// "make" is a construtor command for composed types
		if current.Val == "make" {
			p.i++

			lParent := p.lookAhead(0)
			p.expect(lParent, lexer.Item{Type: lexer.OPERATOR, Val: "("})
			p.i++

			ty, err := p.parseOneType()
			if err != nil {
				panic(err)
			}
			p.i++

			items := p.parseUntil(lexer.Item{Type: lexer.OPERATOR, Val: ")"})
			switch t := ty.(type) {
			case *SliceTypeNode:
				if len(items) == 2 {
					return &InitializeSliceNode{
						Type: t.ItemType,
						Len:  items[0],
						Cap:  items[1],
					}
				} else if len(items) == 1 {
					return &InitializeSliceNode{
						Type: t.ItemType,
						Len:  items[0],
						Cap:  items[0],
					}
				} else {
					panic("wrong argument for slice constructor")
				}
			case *SingleTypeNode:
				if t.TypeName == "string" {
					return &InitializeStringWithSliceNode{
						Items: items,
					}
				}
			default:
				panic("not supported")
			}
		}
		// "extern" is external function without function body

		// single extern: 	extern func foo() int32
		// multiple extern:	extern (
		// 						func foo() int32
		// 						func bar() int32
		// 					)
		if current.Val == "extern" {
			p.i++
			return p.parseExtern()
		}
		// "func" gets converted into a DefineFuncNode
		// the keyword "func" is followed by
		// - a IDENTIFIER (function name)
		// - opening parenthesis
		// - optional: arguments (name type, name2 type2, ...)
		// - closing parenthesis
		// - optional: return type
		// - opening curly bracket ({)
		// - function body
		// - closing curly bracket (})

		// named func:  func abc() {
		// method:      func (a abc) abc() {
		// value func:  func (a abc) {

		if current.Val == "func" {
			p.i++
			defineFunc := p.parseFuncDefinition()

			openBracket := p.lookAhead(0)
			if openBracket.Type != lexer.OPERATOR || openBracket.Val != "{" {
				panic("func arguments must be followed by {. Got " + openBracket.Val)
			}

			p.i++

			defineFunc.Body = p.parseUntil(lexer.Item{Type: lexer.OPERATOR, Val: "}"})

			return p.aheadParse(defineFunc)
		}

		// "return" creates a ReturnNode
		if current.Val == "return" {
			p.i++

			var retVals []Node

			for {
				checkIfEOL := p.lookAhead(0)
				if checkIfEOL.Type == lexer.EOL {
					break
				}

				retVal := p.parseOne(true)
				p.i++

				retVals = append(retVals, retVal)

				checkIfComma := p.lookAhead(0)
				if checkIfComma.Type == lexer.OPERATOR && checkIfComma.Val == "," {
					p.i++
					continue
				}

				break
			}

			res = &ReturnNode{Vals: retVals}
			if withAheadParse {
				res = p.aheadParse(res)
			}
			return
		}

		// Declare a new type
		if current.Val == "type" {
			name := p.lookAhead(1)
			if name.Type != lexer.IDENTIFIER {
				panic("type must be followed by IDENTIFIER")
			}

			p.i += 2

			typeType, err := p.parseOneType()
			if err != nil {
				panic(err)
			}

			// Save the name of the type
			typeType.SetName(name.Val)

			res = &DefineTypeNode{
				Name: name.Val,
				Type: typeType,
			}
			if withAheadParse {
				res = p.aheadParse(res)
			}

			// Register that this type exists
			// TODO: Make it context sensitive (eg package level types, types in functions etc)
			p.types[name.Val] = struct{}{}

			return
		}

		// New instance of type
		if current.Val == "var" || current.Val == "const" {
			p.i++

			isConst := current.Val == "const"

			isGroup := p.lookAhead(0)
			if isGroup.Val == "(" {
				p.i++

				var allocs []*AllocNode
				for {
					nextToken := p.lookAhead(0)
					if nextToken.Val == ")" {
						break
					}
					if nextToken.Type == lexer.EOL {
						p.i++
						continue
					}
					allocs = append(allocs, p.parseVarDecl(isConst))
				}
				return &AllocGroup{Allocs: allocs}
			}

			return p.parseVarDecl(isConst)
		}

		if current.Val == "package" {
			packageName := p.lookAhead(1)

			if packageName.Type != lexer.IDENTIFIER {
				panic("package must be followed by a IDENTIFIER")
			}

			p.i += 1

			return &DeclarePackageNode{
				PackageName: packageName.Val,
			}
		}

		if current.Val == "pragma" {
			p.i++

			key := p.lookAhead(0)
			p.expect(key, lexer.Item{Type: lexer.IDENTIFIER})
			p.i++
			if key.Val != "cellscript" {
				panic("")
			}
			tokenMajor := p.lookAhead(0)
			p.expect(tokenMajor, lexer.Item{Type: lexer.NUMBER})
			major, _ := strconv.Atoi(tokenMajor.Val)
			p.i++
			tokenMinor := p.lookAhead(0)
			p.expect(tokenMajor, lexer.Item{Type: lexer.NUMBER})
			minor, _ := strconv.Atoi(tokenMinor.Val)
			p.i++
			tokenPatch := p.lookAhead(0)
			p.expect(tokenPatch, lexer.Item{Type: lexer.NUMBER})
			patch, _ := strconv.Atoi(tokenPatch.Val)
			p.i++
			return &PragmaNode{
				Version: VersionScheme{
					Major: major,
					Minor: minor,
					Patch: patch,
				},
			}
		}

		if current.Val == "for" {
			return p.parseFor()
		}

		if current.Val == "break" {
			return &BreakNode{}
		}

		if current.Val == "continue" {
			return &ContinueNode{}
		}

		if current.Val == "import" {
			imp := p.parseImport()
			for _, path := range imp.PackagePaths {
				n := filepath.Base(path)
				p.packages[n] = struct{}{}
			}
			return imp
		}

		if current.Val == "true" || current.Val == "false" {
			var v int64 = 0
			if current.Val == "true" {
				v = 1
			}

			return &ConstantNode{
				Type:  BOOL,
				Value: v,
			}
		}

		if current.Val == "range" {
			return p.parseRange()
		}

		if current.Val == "switch" {
			return p.parseSwitch()
		}
	}

	p.printInput()
	log.Panicf("unable to handle default: %d - %+v", p.i, current)
	panic("")
}

func (p *parser) parseVarDecl(isConst bool) *AllocNode {
	allocNode := &AllocNode{Name: p.identifierList(), IsConst: isConst}

	isEq := p.lookAhead(0)
	if (isEq.Type != lexer.OPERATOR || isEq.Val != "=") && isEq.Type != lexer.EOL {
		if isConst {
			// panic("unexpected type in const declaration")
		}

		tp, err := p.parseOneType()
		if err != nil {
			panic(err)
		}
		allocNode.Type = tp
		p.i++
	}

	isEq = p.lookAhead(0)
	if isEq.Type == lexer.OPERATOR && isEq.Val == "=" {
		p.i++
		allocNode.Val = p.expressionList()
	}

	return allocNode
}

func (p *parser) identifierList() (res []string) {
	for {
		n := p.lookAhead(0)
		p.expect(n, lexer.Item{Type: lexer.IDENTIFIER})
		res = append(res, n.Val)

		p.i++

		isComma := p.lookAhead(0)
		if isComma.Type == lexer.OPERATOR && isComma.Val == "," {
			p.i++
			continue
		}

		return
	}
}

func (p *parser) expressionList() (res []Node) {
	for {
		res = append(res, p.parseOne(true))
		p.i++

		isComma := p.lookAhead(0)
		if isComma.Type == lexer.OPERATOR && isComma.Val == "," {
			p.i++
			continue
		}

		return
	}
}

func (p *parser) aheadParse(input Node) Node {
	return p.aheadParseWithOptions(input, true, true)
}

func (p *parser) parseOperation(input Node, withArithAhead bool) Node {
	next := p.lookAhead(1)
	// Handle "Operations" both arith and comparision
	operator := opsCharToOp[next.Val]
	_, isArithOp := arithOperators[operator]

	if !withArithAhead && isArithOp {
		return input
	}

	p.i += 2
	res := &OperatorNode{
		Operator: operator,
		Left:     input,
	}

	if isArithOp {
		res.Right = p.parseOneWithOptions(false, false, false)
		// Sort infix operations if necessary (eg: apply OP_MUL before OP_ADD)
		res = sortInfix(res)
	} else {
		res.Right = p.parseOneWithOptions(true, true, true)
	}

	return p.aheadParseWithOptions(res, true, true)
}

func (p *parser) parseInitializeStructNode(inputType TypeNode) *InitializeStructNode {
	items := make(map[string]Node)

	for {
		// Skip EOLs
		checkIfEOL := p.lookAhead(0)
		if checkIfEOL.Type == lexer.EOL {
			p.i++
		}

		// Find end of parsing
		checkIfEndBracket := p.lookAhead(0)
		if checkIfEndBracket.Type == lexer.OPERATOR && checkIfEndBracket.Val == "}" {
			break
		}

		key := p.lookAhead(0)
		if key.Type != lexer.IDENTIFIER {
			panic("Expected IDENTIFIER in struct initialization")
		}

		col := p.lookAhead(1)
		p.expect(col, lexer.Item{Type: lexer.OPERATOR, Val: ":"})

		p.i += 2

		prevInAlloc := p.inAllocRightHand
		p.inAllocRightHand = false
		items[key.Val] = p.parseOne(true)
		p.inAllocRightHand = prevInAlloc

		p.i++

		commaOrEnd := p.lookAhead(0)
		if commaOrEnd.Type == lexer.OPERATOR && commaOrEnd.Val == "," {
			p.i++
			continue
		}

		if commaOrEnd.Type == lexer.OPERATOR && commaOrEnd.Val == "}" {
			break
		}
	}
	return &InitializeStructNode{
		Type:  inputType,
		Items: items,
	}
}

func (p *parser) aheadParseWithOptions(input Node, withArithAhead, withIdentifierAhead bool) Node {
	next := p.lookAhead(1)

	if next.Type == lexer.OPERATOR {
		if next.Val == "." {
			p.i++

			next = p.lookAhead(1)
			if next.Type == lexer.IDENTIFIER {
				p.i++

				if prevNameNode, ok := input.(*NameNode); ok {
					if _, isPkg := p.packages[prevNameNode.Name]; isPkg {
						return p.aheadParseWithOptions(&NameNode{
							Package: prevNameNode.Name,
							Name:    next.Val,
						}, withArithAhead, withIdentifierAhead)
					}
				}

				return p.aheadParseWithOptions(&StructLoadElementNode{
					Struct:      input,
					ElementName: next.Val,
				}, withArithAhead, withIdentifierAhead)
			}

			if next.Type == lexer.OPERATOR && next.Val == "(" {
				p.i++
				p.i++

				castToType, err := p.parseOneType()
				if err != nil {
					panic(err)
				}

				p.i++

				expectEndParen := p.lookAhead(0)
				p.expect(expectEndParen, lexer.Item{Type: lexer.OPERATOR, Val: ")"})

				p.i++

				return p.aheadParse(&TypeCastInterfaceNode{
					Item: input,
					Type: castToType,
				})
			}

			panic(fmt.Sprintf("Expected IDENTFIER or ( after . Got: %+v", next))
		}

		if next.Val == ":=" || next.Val == "=" {
			p.i += 2

			// TODO: This needs to be a stack
			prevInRight := p.inAllocRightHand
			val := p.parseOne(true)
			p.inAllocRightHand = prevInRight

			if nameNode, ok := input.(*NameNode); ok {
				if next.Val == ":=" {
					return &AllocNode{
						Name: []string{nameNode.Name},
						Val:  []Node{val},
					}
				} else {
					return &AssignNode{
						Target: []Node{nameNode},
						Val:    []Node{val},
					}
				}
			}

			if next.Val == "=" {
				return &AssignNode{
					Target: []Node{input},
					Val:    []Node{val},
				}
			}

			panic(fmt.Sprintf("%s can only be used after a name. Got: %+v", next.Val, input))
		}

		if next.Val == "+=" || next.Val == "-=" || next.Val == "*=" || next.Val == "/=" {
			p.i++
			p.i++

			var op Operator
			switch next.Val {
			case "+=":
				op = OP_ADD
			case "-=":
				op = OP_SUB
			case "*=":
				op = OP_MUL
			case "/=":
				op = OP_DIV
			}

			return &AssignNode{
				Target: []Node{input},
				Val: []Node{&OperatorNode{
					Operator: op,
					Left:     input,
					Right:    p.parseOne(true),
				}},
			}
		}

		if next.Val == "..." {
			p.i++
			return &DeVariadicSliceNode{
				Item: input,
			}
		}

		// Array slicing
		if next.Val == "[" {
			p.i += 2

			index := p.parseOne(true)

			var res Node

			checkIfColon := p.lookAhead(1)
			if checkIfColon.Type == lexer.OPERATOR && checkIfColon.Val == ":" {
				checkIfEndSquare := p.lookAhead(2)
				if checkIfEndSquare.Type == lexer.OPERATOR && checkIfEndSquare.Val == "]" {
					p.i++
					res = &SliceArrayNode{
						Val:   input,
						Start: index,
					}
				} else {
					p.i += 2
					res = &SliceArrayNode{
						Val:    input,
						Start:  index,
						HasEnd: true,
						End:    p.parseOne(true),
					}
				}

			} else {
				res = &LoadArrayElement{
					Array: input,
					Pos:   index,
				}
			}

			p.i++

			expectEndBracket := p.lookAhead(0)
			if expectEndBracket.Type == lexer.OPERATOR && expectEndBracket.Val == "]" {
				return p.aheadParse(res)
			}

			panic(fmt.Sprintf("Unexpected %+v, expected ]", expectEndBracket))
		}

		if next.Val == "--" {
			p.i++
			return &DecrementNode{Item: input}
		}

		if next.Val == "++" {
			p.i++
			return &IncrementNode{Item: input}
		}

		if _, ok := opsCharToOp[next.Val]; ok {
			return p.parseOperation(input, withArithAhead)
		}
	}

	if next.Type == lexer.OPERATOR && next.Val == "(" {
		current := p.lookAhead(0)

		p.i += 2 // identifier and left paren

		vals := p.parseUntil(lexer.Item{Type: lexer.OPERATOR, Val: ")"})
		if _, ok := p.types[current.Val]; ok && len(vals) != 0 { // len(var) == 0 means it is a function call
			if len(vals) != 1 {
				panic("type conversion must take only one argument")
			}
			return p.aheadParse(&TypeCastNode{
				Type: &SingleTypeNode{
					TypeName: current.Val,
				},
				Val: vals[0],
			})
		}

		beforeAllocRightHand := p.inAllocRightHand
		p.inAllocRightHand = false
		callNode := p.aheadParse(&CallNode{
			Function:  input,
			Arguments: vals,
		})
		p.inAllocRightHand = beforeAllocRightHand
		return callNode
	}

	// Initialize structs with values:
	//   Foo{Bar: 123}
	//   Foo{Bar: 123, Bax: hello(123)}
	if next.Type == lexer.OPERATOR && next.Val == "{" {
		nameNode, isNamedNode := input.(*NameNode)
		if isNamedNode {
			_, isType := p.types[nameNode.Name]
			if isType {
				inputType := &SingleTypeNode{
					PackageName: nameNode.Package,
					TypeName:    nameNode.Name,
				}

				p.i += 2
				initStruct := p.parseInitializeStructNode(inputType)
				return p.aheadParse(initStruct)
			}
		}
	}

	if next.Type == lexer.OPERATOR && next.Val == "," {
		// MultiName node parsing ("a, b, c := ...")
		if inputNamedNode, ok := input.(*NameNode); ok {

			// This bit of parsing is specualtive
			//
			// The parser will restore it's position to this index in case it
			// turns out that we can not convert this into a MultiNameNode
			preIndex := p.i

			p.i++
			p.i++

			nextName := p.parseOne(true)

			if nextAlloc, ok := nextName.(*AllocNode); ok {
				nextNames := []string{inputNamedNode.Name}
				nextNames = append(nextNames, nextAlloc.Name...)
				nextAlloc.Name = nextNames

				prev := p.inAllocRightHand
				p.inAllocRightHand = true
				r := p.aheadParse(nextAlloc)
				p.inAllocRightHand = prev

				return r
			}

			if nextAssign, ok := nextName.(*AssignNode); ok {
				nextTargets := []Node{input}
				nextTargets = append(nextTargets, nextAssign.Target...)
				nextAssign.Target = nextTargets

				prev := p.inAllocRightHand
				p.inAllocRightHand = true
				r := p.aheadParse(nextAssign)
				p.inAllocRightHand = prev

				return r
			}

			// A MultiNameNode could not be created
			// Reset the parsing index
			p.i = preIndex
		}

		// MultiValueNode node parsing ("... := 1, 2, 3")
		if p.inAllocRightHand {
			p.i += 2

			// Add to existing multi value if possible. Done on third argument
			// and forward
			if inAllocNode, ok := input.(*AllocNode); ok {
				inAllocNode.Val = append(inAllocNode.Val, p.parseOne(false))
				return p.aheadParse(inAllocNode)
			}

			if inAssignValue, ok := input.(*AssignNode); ok {
				inAssignValue.Val = append(inAssignValue.Val, p.parseOne(false))
				return p.aheadParse(inAssignValue)
			}

			panic("unexpected in alloc right hand")
		}
	}

	return input
}

func (p *parser) lookAhead(steps int) lexer.Item {
	return p.input[p.i+steps]
}

func (p *parser) parseUntil(until lexer.Item) []Node {
	n, _ := p.parseUntilEither([]lexer.Item{until})
	return n
}

// parseUntilEither reads lexer items until it finds one that equals to a item in "untils"
// The list of parsed nodes is returned in res. The lexer item that stopped the iteration
// is returned in "reached"
func (p *parser) parseUntilEither(untils []lexer.Item) (res []Node, reached lexer.Item) {
	for {
		current := p.input[p.i]

		// Check if we have reached the end
		for _, until := range untils {
			if current.Type == until.Type && current.Val == until.Val {
				return res, until
			}
		}

		// Ignore comma
		if current.Type == lexer.OPERATOR && current.Val == "," {
			p.i++
			continue
		}

		// Ignore EOL EOF
		if current.Type == lexer.EOL || current.Type == lexer.EOF {
			p.i++
			continue
		}

		one := p.parseOne(true)
		if one != nil {
			next := p.lookAhead(1)
			if _, isOperationNode := opsCharToOp[next.Val]; next.Type == lexer.OPERATOR && isOperationNode {
				one = p.parseOperation(one, false)
			}
			res = append(res, one)
		}

		p.i++
	}
}

func (p *parser) parseFunctionArguments() []*NameNode {
	var res []*NameNode
	var i int

	for {
		current := p.input[p.i]
		if current.Type == lexer.OPERATOR && current.Val == ")" {
			p.i++
			return res
		}

		if i > 0 {
			if current.Type != lexer.OPERATOR && current.Val != "," {
				panic("arguments must be separated by commas. Got: " + fmt.Sprintf("%+v", current))
			}

			p.i++
			current = p.input[p.i]
		}

		name := p.lookAhead(0)
		if name.Type != lexer.IDENTIFIER {
			panic("function arguments: variable name must be identifier. Got: " + fmt.Sprintf("%+v", name))
		}
		p.i++

		argType, err := p.parseOneType()
		if err != nil {
			panic(err)
		}
		p.i++

		res = append(res, &NameNode{
			Name: name.Val,
			Type: argType,
		})

		i++
	}
}

func (p *parser) parseOneType() (TypeNode, error) {
	current := p.lookAhead(0)

	isVariadic := false
	if current.Type == lexer.OPERATOR && current.Val == "..." {
		isVariadic = true

		p.i++
		current = p.lookAhead(0)
	}

	// pointer types
	if current.Type == lexer.OPERATOR && current.Val == "*" {
		p.i++
		valType, err := p.parseOneType()
		if err != nil {
			panic(err)
		}
		return &PointerTypeNode{
			ValueType:  valType,
			IsVariadic: isVariadic,
		}, nil
	}

	// struct parsing
	if current.Type == lexer.KEYWORD && (current.Val == "table" || current.Val == "struct") {
		p.i++

		res := &StructTypeNode{
			Types:      make([]TypeNode, 0),
			Names:      make(map[string]int),
			IsVariadic: isVariadic,
		}

		current = p.lookAhead(0)
		if current.Type != lexer.OPERATOR || current.Val != "{" {
			panic("struct must be followed by {")
		}
		p.i++

		for {
			itemName := p.lookAhead(0)

			// Ignore EOL
			if itemName.Type == lexer.EOL {
				p.i++
				continue
			}

			// Stop at }
			if itemName.Type == lexer.OPERATOR && itemName.Val == "}" {
				break
			}

			if itemName.Type != lexer.IDENTIFIER {
				panic("expected IDENTIFIER in struct{}, got " + fmt.Sprintf("%+v", itemName))
			}
			p.i++

			itemType, err := p.parseOneType()
			if err != nil {
				panic("expected TYPE in struct{}, got: " + err.Error())
			}
			p.i++

			res.Types = append(res.Types, itemType)
			res.Names[itemName.Val] = len(res.Types) - 1

			current = p.lookAhead(0)
		}

		return res, nil
	}

	if current.Type == lexer.KEYWORD && current.Val == "interface" {
		p.i++
		p.expect(p.lookAhead(0), lexer.Item{Type: lexer.OPERATOR, Val: "{"})
		p.i++

		ifaceType := &InterfaceTypeNode{
			IsVariadic: isVariadic,
		}

		// Parse methods if set
		for {
			current := p.lookAhead(0)
			if current.Type == lexer.EOL {
				p.i++
				continue
			}
			if current.Type == lexer.OPERATOR && current.Val == "}" {
				break
			}

			// Expect method name
			p.expect(current, lexer.Item{Type: lexer.IDENTIFIER})

			methodName := current.Val
			methodDef := InterfaceMethod{}

			p.i++
			p.expect(p.lookAhead(0), lexer.Item{Type: lexer.OPERATOR, Val: "("})

			// Check if the method takes any arguments
			for {
				p.i++
				current = p.lookAhead(0)
				if current.Type == lexer.OPERATOR && current.Val == ")" {
					p.i++
					break
				}

				current = p.lookAhead(0)
				if current.Type == lexer.OPERATOR && current.Val == "," {
					continue
				}

				argumentType, err := p.parseOneType()
				if err != nil {
					panic(err)
				}

				methodDef.ArgumentTypes = append(methodDef.ArgumentTypes, argumentType)
			}

			// Function return types
			for {
				current = p.lookAhead(0)
				if current.Type == lexer.EOL {
					p.i++
					break
				}

				if current.Type == lexer.OPERATOR && current.Val == "(" {
					p.i++
					continue
				}

				if current.Type == lexer.OPERATOR && current.Val == "," {
					p.i++
					continue
				}

				if current.Type == lexer.OPERATOR && current.Val == ")" {
					p.i++
					break
				}

				returnType, err := p.parseOneType()
				if err != nil {
					panic(err)
				}

				methodDef.ReturnTypes = append(methodDef.ReturnTypes, returnType)
				p.i++
			}

			if ifaceType.Methods == nil {
				ifaceType.Methods = make(map[string]InterfaceMethod)
			}
			ifaceType.Methods[methodName] = methodDef
		}

		return ifaceType, nil
	}

	if current.Type == lexer.IDENTIFIER {
		// Types from other packages ("foo.Bar")
		isTypeFromPkg := p.lookAhead(1)
		if isTypeFromPkg.Type == lexer.OPERATOR && isTypeFromPkg.Val == "." {
			typeName := p.lookAhead(2)
			p.expect(typeName, lexer.Item{Type: lexer.IDENTIFIER})
			p.i += 2
			return &SingleTypeNode{
				PackageName: current.Val,
				TypeName:    typeName.Val,
				IsVariadic:  isVariadic,
			}, nil
		}

		// Types from the current package
		return &SingleTypeNode{
			TypeName:   current.Val,
			IsVariadic: isVariadic,
		}, nil
	}

	// Array parsing
	if current.Type == lexer.OPERATOR && current.Val == "[" {
		arrayLenght := p.lookAhead(1)

		// Slice parsing
		if arrayLenght.Type == lexer.OPERATOR && arrayLenght.Val == "]" {
			p.i += 2

			sliceItemType, err := p.parseOneType()
			if err != nil {
				return nil, errors.New("arrayParse failed: " + err.Error())
			}

			return &SliceTypeNode{
				ItemType:   sliceItemType,
				IsVariadic: isVariadic,
			}, nil
		}

		if arrayLenght.Type != lexer.NUMBER {
			return nil, errors.New("parseArray failed: Expected number or ] after [")
		}
		arrayLengthInt, err := strconv.Atoi(arrayLenght.Val)
		if err != nil {
			return nil, err
		}

		expectEndBracket := p.lookAhead(2)
		if expectEndBracket.Type != lexer.OPERATOR || expectEndBracket.Val != "]" {
			return nil, errors.New("parseArray failed: Expected ] in array type")
		}

		p.i += 3

		arrayItemType, err := p.parseOneType()
		if err != nil {
			return nil, errors.New("arrayParse failed: " + err.Error())
		}

		return &ArrayTypeNode{
			ItemType:   arrayItemType,
			Len:        int64(arrayLengthInt),
			IsVariadic: isVariadic,
		}, nil
	}

	// Func type parsing
	if current.Type == lexer.KEYWORD && current.Val == "func" {
		p.i++

		expectOpenParen := p.lookAhead(0)
		if expectOpenParen.Type != lexer.OPERATOR && expectOpenParen.Val != "(" {
			return nil, errors.New("parse func failed, expected ( after func")
		}
		p.i++

		fn := &FuncTypeNode{}

		multiTypeParse := func() ([]TypeNode, error) {
			var typeList []TypeNode

			for {
				checkIfEndParen := p.lookAhead(0)
				if checkIfEndParen.Type == lexer.OPERATOR && checkIfEndParen.Val == ")" {
					break
				}

				argType, err := p.parseOneType()
				if err != nil {
					return nil, err
				}

				typeList = append(typeList, argType)

				p.i++

				expectCommaOrEndParen := p.lookAhead(0)
				if expectCommaOrEndParen.Type == lexer.OPERATOR && expectCommaOrEndParen.Val == "," {
					p.i++
					continue
				}

				if expectCommaOrEndParen.Type == lexer.OPERATOR && expectCommaOrEndParen.Val == ")" {
					continue
				}

				return nil, errors.New("expected ) or , in func arg parsing")
			}

			return typeList, nil
		}

		// List of arguments
		var err error
		fn.ArgTypes, err = multiTypeParse()
		if err != nil {
			return nil, fmt.Errorf("unable to parse func arguments: %s", err)
		}

		// Return type parsing
		// Possible formats:
		// - Nothing
		// - T
		// - (T1, T2, ... )

		checkIfParenOrType := p.lookAhead(1)

		// Multiple types
		if checkIfParenOrType.Type == lexer.OPERATOR && checkIfParenOrType.Val == "(" {
			p.i++
			p.i++

			fn.RetTypes, err = multiTypeParse()
			if err != nil {
				return nil, fmt.Errorf("unable to parse func return values: %s", err)
			}

			return fn, nil
		}

		// Single types
		isPointerType := checkIfParenOrType.Type == lexer.OPERATOR && checkIfParenOrType.Val == "*"
		if checkIfParenOrType.Type == lexer.IDENTIFIER || isPointerType {
			p.i++

			t, err := p.parseOneType()
			if err != nil {
				return nil, err
			}

			fn.RetTypes = []TypeNode{t}
			return fn, nil
		}

		return fn, nil
	}

	return nil, errors.New("parseOneType failed: " + fmt.Sprintf("%+v", current))
}

// panics if check fails
func (p *parser) expect(input lexer.Item, expected lexer.Item) {
	if expected.Type != input.Type {
		panic(fmt.Sprintf("Expected %+v got %+v", expected, input))
	}

	if expected.Val != "" && expected.Val != input.Val {
		panic(fmt.Sprintf("Expected %+v got %+v", expected, input))
	}
}
