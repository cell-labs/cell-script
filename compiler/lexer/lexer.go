package lexer

import (
	"fmt"
	"strings"
)

type lexType uint8

const (
	IDENTIFIER lexType = iota
	KEYWORD
	NUMBER
	HEX
	STRING
	BYTE
	OPERATOR
	EOF
	EOL
)

type Item struct {
	Type lexType
	Val  string
	Line int
}

func (i Item) String() string {
	var t string
	switch i.Type {
	case IDENTIFIER:
		t = "IDENTIFIER"
	case KEYWORD:
		t = "KEYWORD"
	case NUMBER:
		t = "NUMBER"
	case STRING:
		t = "STRING"
	case BYTE:
		t = "BYTE"
	case OPERATOR:
		t = "OPERATOR"
	case EOF:
		t = "EOF"
	case EOL:
		t = "EOL"
	}

	return fmt.Sprintf("{Type:%s, Val:%s, Line:%d}", t, i.Val, i.Line)
}

// https://golang.org/ref/spec#Operators_and_punctuation
var operations = map[string]struct{}{
	"+":   {},
	"&":   {},
	"+=":  {},
	"&=":  {},
	"&&":  {},
	"==":  {},
	"!=":  {},
	"(":   {},
	")":   {},
	"-":   {},
	"|":   {},
	"-=":  {},
	"|=":  {},
	"||":  {},
	"<":   {},
	"<=":  {},
	"[":   {},
	"]":   {},
	"*":   {},
	"^":   {},
	"*=":  {},
	"^=":  {},
	"<-":  {},
	">":   {},
	">=":  {},
	"{":   {},
	"}":   {},
	"/":   {},
	"<<":  {},
	"/=":  {},
	"<<=": {},
	"++":  {},
	"=":   {},
	":=":  {},
	",":   {},
	";":   {},
	"%":   {},
	">>":  {},
	"%=":  {},
	">>=": {},
	"--":  {},
	"!":   {},
	"...": {},
	".":   {},
	":":   {},
	"&^":  {},
	"&^=": {},

	// TODO: Remove
	"..": {}, // is not a real operation. Is there so that ... can be found.
}

var escapedChar = map[byte]byte{'t': '\t', 'n': '\n', 'r': '\r', '\'': '\''}

func Lex(inputFullSource string) []Item {
	var res []Item

	for line, input := range strings.Split(inputFullSource, "\n") {
		if line > 0 {
			res = append(res, Item{Type: EOL, Line: line})
		}

		// Lines starts at 1
		line = line + 1

		i := 0

		for i < len(input) {

			// Comment, until end of line or end of file
			if input[i] == '/' && input[i+1] == '/' {
				break
			}

			if _, ok := operations[string(input[i])]; ok {

				operator := string(input[i])

				// Parse till end of the operator
				// Can be up to 3 characters long
				for {
					if len(input) == i+1 {
						break
					}

					checkIfOperator := operator + string(input[i+1])
					if _, ok := operations[checkIfOperator]; ok {
						operator = checkIfOperator
						i++
						continue
					} else {
						break
					}
				}

				res = append(res, Item{Type: OPERATOR, Val: operator, Line: line})
				i++
				continue
			}

			if input[i] == '"' {
				// String continues until next unescaped "
				var str string

				i++

				for i < len(input) {
					if input[i] == '"' {
						break
					}

					// parse escape sequences
					if input[i] == '\\' {
						if esc, ok := escapeSequences[string(input[i])+string(input[i+1])]; ok {
							str += esc
							i += 2
							continue
						}
					}

					str += string(input[i])
					i++
				}

				i++
				res = append(res, Item{Type: STRING, Val: str, Line: line})
				continue
			}

			if input[i] == '\'' {
				// String continues until next unescaped '
				var str string

				i++

				for i < len(input) {
					if input[i] == '\'' {
						break
					}
					// parse escape char
					if input[i] == '\\' {
						i++
						va, exist := escapedChar[input[i]]
						if !exist {
							panic("Unsupported escaped character")
						}
						str = string(va)
						i++
						continue
					}

					str = string(input[i])
					i++
				}

				i++
				if str == "" {
					panic("Ilegal byte")
				}
				res = append(res, Item{Type: BYTE, Val: string(str), Line: line})
				continue
			}

			// NAME
			// Consists of a-z, parse until the last allowed char
			if (input[i] >= 'a' && input[i] <= 'z') || (input[i] >= 'A' && input[i] <= 'Z') || input[i] == '_' {
				name := ""

				for i < len(input) && ((input[i] >= 'a' && input[i] <= 'z') ||
					(input[i] >= 'A' && input[i] <= 'Z') ||
					(input[i] >= '0' && input[i] <= '9') ||
					input[i] == '_') {
					name += string(input[i])
					i++
				}

				if _, ok := keywords[name]; ok {
					res = append(res, Item{Type: KEYWORD, Val: name, Line: line})
				} else {
					res = append(res, Item{Type: IDENTIFIER, Val: name, Line: line})
				}

				continue
			}

			// NUMBER
			// 0-9
			lexNumber := func() bool {
				isHex := false
				if i + 1 < len(input) {
					if input[i] == '0' && input[i+1] == 'x' {
						isHex = true
						i += 2
					}
				}
				if i >= len(input) {
					return false
				}
				if isHex {
					if (input[i] >= '0' && input[i] <= '9') || (input[i] >= 'a' && input[i] <= 'f') || (input[i] >= 'A' && input[i] <= 'F') {
						val := ""
						for (input[i] >= '0' && input[i] <= '9') || (input[i] >= 'a' && input[i] <= 'f') || (input[i] >= 'A' && input[i] <= 'F') {
							val += string(input[i])
							i++
							if i >= len(input) {
								break
							}
						}
						res = append(res, Item{Type: HEX, Val: val, Line: line})
						return true
					}
				}
				if input[i] >= '0' && input[i] <= '9' {
					val := ""
					for i < len(input) && input[i] >= '0' && input[i] <= '9' {
						val += string(input[i])
						i++
					}
					res = append(res, Item{Type: NUMBER, Val: val, Line: line})
					return true
				}
				return false
			}
			if succ := lexNumber(); succ {
				if i >= len(input) || input[i] != '.' {
					continue
				}
				// VERSION: NUMBER.NUMBER.NUMBER
				i++ // skip over "."
				isNumber := lexNumber()
				if !isNumber {
					i--
					continue
				}
				if i >= len(input) || input[i] != '.' {
					continue
				}
				i++ // skip over "."
				isNumber = lexNumber()
				if !isNumber {
					panic("Unexpected char in Lexer: " + string(input[i]))
				}
				continue
			}

			// Whitespace (ignore)
			if len(strings.TrimSpace(string(input[i]))) == 0 {
				i++
				continue
			}

			panic("Unexpected char in Lexer: " + string(input[i]))
		}
	}

	res = append(res, Item{Type: EOL}, Item{Type: EOF})

	return res
}
