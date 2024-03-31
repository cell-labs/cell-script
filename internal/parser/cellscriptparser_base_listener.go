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

// EnterFunctionStmt is called when production functionStmt is entered.
func (s *BaseCellScriptParserListener) EnterFunctionStmt(ctx *FunctionStmtContext) {}

// ExitFunctionStmt is called when production functionStmt is exited.
func (s *BaseCellScriptParserListener) ExitFunctionStmt(ctx *FunctionStmtContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseCellScriptParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseCellScriptParserListener) ExitEos(ctx *EosContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseCellScriptParserListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseCellScriptParserListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseCellScriptParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseCellScriptParserListener) ExitDeclaration(ctx *DeclarationContext) {}

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

// EnterBody is called when production body is entered.
func (s *BaseCellScriptParserListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseCellScriptParserListener) ExitBody(ctx *BodyContext) {}

// EnterTypeParameterDecl is called when production typeParameterDecl is entered.
func (s *BaseCellScriptParserListener) EnterTypeParameterDecl(ctx *TypeParameterDeclContext) {}

// ExitTypeParameterDecl is called when production typeParameterDecl is exited.
func (s *BaseCellScriptParserListener) ExitTypeParameterDecl(ctx *TypeParameterDeclContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCellScriptParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCellScriptParserListener) ExitExpression(ctx *ExpressionContext) {}
