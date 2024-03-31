package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"

	"github.com/cell-labs/cell-script/internal/lexer"
	"github.com/cell-labs/cell-script/internal/parser"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream(`function a() { int a = 1 }`)

	// Create the Lexer
	lexer := lexer.NewCellScriptLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}

	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewCellScriptParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
}