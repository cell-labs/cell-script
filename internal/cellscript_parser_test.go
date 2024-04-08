package parser

import (
	"testing"

	"github.com/cell-labs/cell-script/internal/lex"
	"github.com/cell-labs/cell-script/internal/parse"

	"github.com/antlr4-go/antlr/v4"
)

func TestParser(t *testing.T) {

	lexer := lex.NewCellScriptLexer(antlr.NewInputStream("a+b"))
	// generate AST using parser
	parser := parse.NewCellScriptParser(antlr.NewCommonTokenStream(lexer, 0))
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	parser.SourceFile().Accept(&parse.BaseCellScriptParserVisitor{})
}
