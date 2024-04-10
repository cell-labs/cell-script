package ast

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/cell-labs/cell-script/internal/parse"
)

// dump ast node
func Dump(parser *parse.CellScriptParser) string {
	return antlr.TreesStringTree(parser.SourceFile(), parser.RuleNames, parser)
}
