package codegen

import (
	"github.com/antlr4-go/antlr/v4"
	. "github.com/cell-labs/cell-script/internal/parse"
)

type Generator struct {
	*antlr.BaseParseTreeVisitor
	builder *IrBuilder
}

func NewGenerator(parser *CellScriptParser) *Generator {
	g := &Generator{&antlr.BaseParseTreeVisitor{}, NewIrBuilder()}
	g.Visit(parser.SourceFile())
	return g
}

func (g *Generator) VisitSourceFile(ctx *SourceFileContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitImportStmt(ctx *ImportStmtContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitEos(ctx *EosContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	GenFunc()
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitTypeParameters(ctx *TypeParametersContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitSignature(ctx *SignatureContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitBlock(ctx *BlockContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitTypeParameterDecl(ctx *TypeParameterDeclContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitExpression(ctx *ExpressionContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitArithmeticExpr(ctx *ArithmeticExprContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitReturnExpr(ctx *ReturnExprContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitStatement(ctx *StatementContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitSimpleStmt(ctx *SimpleStmtContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitExpressionStmt(ctx *ExpressionStmtContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitAssignment(ctx *AssignmentContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitAssign_op(ctx *Assign_opContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitForStmt(ctx *ForStmtContext) interface{} {
	return g.VisitChildren(ctx)
}

func (g *Generator) VisitForClause(ctx *ForClauseContext) interface{} {
	return g.VisitChildren(ctx)
}
