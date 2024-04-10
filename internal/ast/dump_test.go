package ast

import (
	"testing"

	"github.com/cell-labs/cell-script/internal/ast"
	"github.com/cell-labs/cell-script/internal/lex"
	"github.com/cell-labs/cell-script/internal/parse"
	"github.com/stretchr/testify/assert"

	"github.com/antlr4-go/antlr/v4"
)

func parseAndCompare(t *testing.T, src, expected string) {
	lexer := lex.NewCellScriptLexer(antlr.NewInputStream(src))
	// generate AST using parser
	parser := parse.NewCellScriptParser(antlr.NewCommonTokenStream(lexer, 0))
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	assert.Equal(t, ast.Dump(parser), expected)
}

func TestDumper(t *testing.T) {
	parseAndCompare(t, `import "tx" function foo() bool { return a+b }`, `(sourceFile (importStmt (importDecl import "tx") eos) (declaration (functionDecl function foo (typeParameters ( )) (signature bool) (block { (expression (returnExpr return a)) + b })) (eos <EOF>) <EOF>) <EOF>)`)
}
