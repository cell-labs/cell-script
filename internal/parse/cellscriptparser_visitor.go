// Code generated from CellScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parse // CellScriptParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CellScriptParser.
type CellScriptParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CellScriptParser#sourceFile.
	VisitSourceFile(ctx *SourceFileContext) interface{}

	// Visit a parse tree produced by CellScriptParser#importStmt.
	VisitImportStmt(ctx *ImportStmtContext) interface{}

	// Visit a parse tree produced by CellScriptParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by CellScriptParser#eos.
	VisitEos(ctx *EosContext) interface{}

	// Visit a parse tree produced by CellScriptParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by CellScriptParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by CellScriptParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by CellScriptParser#typeParameters.
	VisitTypeParameters(ctx *TypeParametersContext) interface{}

	// Visit a parse tree produced by CellScriptParser#signature.
	VisitSignature(ctx *SignatureContext) interface{}

	// Visit a parse tree produced by CellScriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by CellScriptParser#typeParameterDecl.
	VisitTypeParameterDecl(ctx *TypeParameterDeclContext) interface{}

	// Visit a parse tree produced by CellScriptParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by CellScriptParser#arithmeticExpr.
	VisitArithmeticExpr(ctx *ArithmeticExprContext) interface{}

	// Visit a parse tree produced by CellScriptParser#returnExpr.
	VisitReturnExpr(ctx *ReturnExprContext) interface{}

	// Visit a parse tree produced by CellScriptParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by CellScriptParser#simpleStmt.
	VisitSimpleStmt(ctx *SimpleStmtContext) interface{}

	// Visit a parse tree produced by CellScriptParser#expressionStmt.
	VisitExpressionStmt(ctx *ExpressionStmtContext) interface{}

	// Visit a parse tree produced by CellScriptParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by CellScriptParser#assign_op.
	VisitAssign_op(ctx *Assign_opContext) interface{}

	// Visit a parse tree produced by CellScriptParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by CellScriptParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by CellScriptParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by CellScriptParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by CellScriptParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by CellScriptParser#forClause.
	VisitForClause(ctx *ForClauseContext) interface{}
}
