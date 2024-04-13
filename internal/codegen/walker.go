package codegen

import (
	"github.com/antlr4-go/antlr/v4"
)

// overwrite ParseTreeVisitor interface for Visitor
func (g *Generator) Visit(ctx antlr.ParseTree) interface{} {
	return ctx.Accept(g)
}

func (g *Generator) VisitChildren(ctx antlr.RuleNode) interface{} {
	for _, child := range ctx.GetChildren() {
		child.(antlr.ParseTree).Accept(g)
	}
	return nil
}
