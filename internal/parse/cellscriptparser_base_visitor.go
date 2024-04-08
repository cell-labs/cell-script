// Code generated from CellScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parse // CellScriptParser
import "github.com/antlr4-go/antlr/v4"

type BaseCellScriptParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCellScriptParserVisitor) VisitSourceFile(ctx *SourceFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitImportStmt(ctx *ImportStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitTypeParameters(ctx *TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitSignature(ctx *SignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitTypeParameterDecl(ctx *TypeParameterDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitArithmeticExpr(ctx *ArithmeticExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitReturnExpr(ctx *ReturnExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitSimpleStmt(ctx *SimpleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitExpressionStmt(ctx *ExpressionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitAssign_op(ctx *Assign_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitForClause(ctx *ForClauseContext) interface{} {
	return v.VisitChildren(ctx)
}
