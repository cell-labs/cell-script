// Code generated from CellScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CellScriptParser
import "github.com/antlr4-go/antlr/v4"

// CellScriptParserListener is a complete listener for a parse tree produced by CellScriptParser.
type CellScriptParserListener interface {
	antlr.ParseTreeListener

	// EnterSourceFile is called when entering the sourceFile production.
	EnterSourceFile(c *SourceFileContext)

	// EnterImportStmt is called when entering the importStmt production.
	EnterImportStmt(c *ImportStmtContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// EnterImportDecl is called when entering the importDecl production.
	EnterImportDecl(c *ImportDeclContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterFunctionDecl is called when entering the functionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterSignature is called when entering the signature production.
	EnterSignature(c *SignatureContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterTypeParameterDecl is called when entering the typeParameterDecl production.
	EnterTypeParameterDecl(c *TypeParameterDeclContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterArithmeticExpr is called when entering the arithmeticExpr production.
	EnterArithmeticExpr(c *ArithmeticExprContext)

	// EnterReturnExpr is called when entering the returnExpr production.
	EnterReturnExpr(c *ReturnExprContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSimpleStmt is called when entering the simpleStmt production.
	EnterSimpleStmt(c *SimpleStmtContext)

	// EnterExpressionStmt is called when entering the expressionStmt production.
	EnterExpressionStmt(c *ExpressionStmtContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterAssign_op is called when entering the assign_op production.
	EnterAssign_op(c *Assign_opContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the continueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterForClause is called when entering the forClause production.
	EnterForClause(c *ForClauseContext)

	// ExitSourceFile is called when exiting the sourceFile production.
	ExitSourceFile(c *SourceFileContext)

	// ExitImportStmt is called when exiting the importStmt production.
	ExitImportStmt(c *ImportStmtContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)

	// ExitImportDecl is called when exiting the importDecl production.
	ExitImportDecl(c *ImportDeclContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitFunctionDecl is called when exiting the functionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitSignature is called when exiting the signature production.
	ExitSignature(c *SignatureContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitTypeParameterDecl is called when exiting the typeParameterDecl production.
	ExitTypeParameterDecl(c *TypeParameterDeclContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitArithmeticExpr is called when exiting the arithmeticExpr production.
	ExitArithmeticExpr(c *ArithmeticExprContext)

	// ExitReturnExpr is called when exiting the returnExpr production.
	ExitReturnExpr(c *ReturnExprContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSimpleStmt is called when exiting the simpleStmt production.
	ExitSimpleStmt(c *SimpleStmtContext)

	// ExitExpressionStmt is called when exiting the expressionStmt production.
	ExitExpressionStmt(c *ExpressionStmtContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitAssign_op is called when exiting the assign_op production.
	ExitAssign_op(c *Assign_opContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the continueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitForClause is called when exiting the forClause production.
	ExitForClause(c *ForClauseContext)
}
