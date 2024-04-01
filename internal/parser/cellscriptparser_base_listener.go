// Code generated from CellScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CellScriptParser
import "github.com/antlr4-go/antlr/v4"

// BaseCellScriptParserListener is a complete listener for a parse tree produced by CellScriptParser.
type BaseCellScriptParserListener struct{}

var _ CellScriptParserListener = &BaseCellScriptParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCellScriptParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCellScriptParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCellScriptParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCellScriptParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *BaseCellScriptParserListener) EnterSourceFile(ctx *SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *BaseCellScriptParserListener) ExitSourceFile(ctx *SourceFileContext) {}

// EnterImportStmt is called when production importStmt is entered.
func (s *BaseCellScriptParserListener) EnterImportStmt(ctx *ImportStmtContext) {}

// ExitImportStmt is called when production importStmt is exited.
func (s *BaseCellScriptParserListener) ExitImportStmt(ctx *ImportStmtContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseCellScriptParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseCellScriptParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseCellScriptParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseCellScriptParserListener) ExitEos(ctx *EosContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseCellScriptParserListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseCellScriptParserListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseCellScriptParserListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseCellScriptParserListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *BaseCellScriptParserListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production functionDecl is exited.
func (s *BaseCellScriptParserListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseCellScriptParserListener) EnterTypeParameters(ctx *TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseCellScriptParserListener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterSignature is called when production signature is entered.
func (s *BaseCellScriptParserListener) EnterSignature(ctx *SignatureContext) {}

// ExitSignature is called when production signature is exited.
func (s *BaseCellScriptParserListener) ExitSignature(ctx *SignatureContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseCellScriptParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseCellScriptParserListener) ExitBlock(ctx *BlockContext) {}

// EnterTypeParameterDecl is called when production typeParameterDecl is entered.
func (s *BaseCellScriptParserListener) EnterTypeParameterDecl(ctx *TypeParameterDeclContext) {}

// ExitTypeParameterDecl is called when production typeParameterDecl is exited.
func (s *BaseCellScriptParserListener) ExitTypeParameterDecl(ctx *TypeParameterDeclContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCellScriptParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCellScriptParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterArithmeticExpr is called when production arithmeticExpr is entered.
func (s *BaseCellScriptParserListener) EnterArithmeticExpr(ctx *ArithmeticExprContext) {}

// ExitArithmeticExpr is called when production arithmeticExpr is exited.
func (s *BaseCellScriptParserListener) ExitArithmeticExpr(ctx *ArithmeticExprContext) {}

// EnterReturnExpr is called when production returnExpr is entered.
func (s *BaseCellScriptParserListener) EnterReturnExpr(ctx *ReturnExprContext) {}

// ExitReturnExpr is called when production returnExpr is exited.
func (s *BaseCellScriptParserListener) ExitReturnExpr(ctx *ReturnExprContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCellScriptParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCellScriptParserListener) ExitStatement(ctx *StatementContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BaseCellScriptParserListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BaseCellScriptParserListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterExpressionStmt is called when production expressionStmt is entered.
func (s *BaseCellScriptParserListener) EnterExpressionStmt(ctx *ExpressionStmtContext) {}

// ExitExpressionStmt is called when production expressionStmt is exited.
func (s *BaseCellScriptParserListener) ExitExpressionStmt(ctx *ExpressionStmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseCellScriptParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseCellScriptParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssign_op is called when production assign_op is entered.
func (s *BaseCellScriptParserListener) EnterAssign_op(ctx *Assign_opContext) {}

// ExitAssign_op is called when production assign_op is exited.
func (s *BaseCellScriptParserListener) ExitAssign_op(ctx *Assign_opContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseCellScriptParserListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseCellScriptParserListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseCellScriptParserListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseCellScriptParserListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseCellScriptParserListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseCellScriptParserListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseCellScriptParserListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseCellScriptParserListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseCellScriptParserListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseCellScriptParserListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterForClause is called when production forClause is entered.
func (s *BaseCellScriptParserListener) EnterForClause(ctx *ForClauseContext) {}

// ExitForClause is called when production forClause is exited.
func (s *BaseCellScriptParserListener) ExitForClause(ctx *ForClauseContext) {}
