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

	// EnterFunctionStmt is called when entering the functionStmt production.
	EnterFunctionStmt(c *FunctionStmtContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// EnterImportDecl is called when entering the importDecl production.
	EnterImportDecl(c *ImportDeclContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterFunctionDecl is called when entering the functionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterSignature is called when entering the signature production.
	EnterSignature(c *SignatureContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterTypeParameterDecl is called when entering the typeParameterDecl production.
	EnterTypeParameterDecl(c *TypeParameterDeclContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitSourceFile is called when exiting the sourceFile production.
	ExitSourceFile(c *SourceFileContext)

	// ExitImportStmt is called when exiting the importStmt production.
	ExitImportStmt(c *ImportStmtContext)

	// ExitFunctionStmt is called when exiting the functionStmt production.
	ExitFunctionStmt(c *FunctionStmtContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)

	// ExitImportDecl is called when exiting the importDecl production.
	ExitImportDecl(c *ImportDeclContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitFunctionDecl is called when exiting the functionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitSignature is called when exiting the signature production.
	ExitSignature(c *SignatureContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitTypeParameterDecl is called when exiting the typeParameterDecl production.
	ExitTypeParameterDecl(c *TypeParameterDeclContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
