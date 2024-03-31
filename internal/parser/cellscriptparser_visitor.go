// Code generated from CellScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CellScriptParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CellScriptParser.
type CellScriptParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CellScriptParser#sourceFile.
	VisitSourceFile(ctx *SourceFileContext) interface{}

	// Visit a parse tree produced by CellScriptParser#importStmt.
	VisitImportStmt(ctx *ImportStmtContext) interface{}

	// Visit a parse tree produced by CellScriptParser#functionStmt.
	VisitFunctionStmt(ctx *FunctionStmtContext) interface{}

	// Visit a parse tree produced by CellScriptParser#eos.
	VisitEos(ctx *EosContext) interface{}

	// Visit a parse tree produced by CellScriptParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by CellScriptParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by CellScriptParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by CellScriptParser#typeParameters.
	VisitTypeParameters(ctx *TypeParametersContext) interface{}

	// Visit a parse tree produced by CellScriptParser#signature.
	VisitSignature(ctx *SignatureContext) interface{}

	// Visit a parse tree produced by CellScriptParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by CellScriptParser#typeParameterDecl.
	VisitTypeParameterDecl(ctx *TypeParameterDeclContext) interface{}

	// Visit a parse tree produced by CellScriptParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}
}
