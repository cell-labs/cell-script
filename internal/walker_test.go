package test

import (
	"testing"

	"github.com/cell-labs/cell-script/internal/lex"
	"github.com/cell-labs/cell-script/internal/parse"

	"github.com/antlr4-go/antlr/v4"
)

func parseAndCompare(src string) string {
	lexer := lex.NewCellScriptLexer(antlr.NewInputStream(src))
	// generate AST using parser
	parser := parse.NewCellScriptParser(antlr.NewCommonTokenStream(lexer, 0))
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	walker := parse.NewWalker()
	walker.Visit(parser.SourceFile())
	return ""
}

func TestWalker(t *testing.T) {
	parseAndCompare(`import "tx" function foo() bool { return a+b }`)
}
