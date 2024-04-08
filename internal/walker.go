package parse

import (
	"github.com/antlr4-go/antlr/v4"
)

type Walker struct {
	*BaseCellScriptParserVisitor
}

func NewWalker() *Walker {
	return &Walker{&BaseCellScriptParserVisitor{}}
}

// overwrite ParseTreeVisitor interface for BaseCellScriptParserVisitor
func (v *BaseCellScriptParserVisitor) Visit(ctx antlr.ParseTree) interface{} {
	return ctx.Accept(v)
}

func (v *BaseCellScriptParserVisitor) VisitChildren(ctx antlr.RuleNode) interface{} {
	for _, child := range ctx.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
