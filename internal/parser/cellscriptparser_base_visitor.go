// Code generated from CellScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CellScriptParser
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

func (v *BaseCellScriptParserVisitor) VisitFunctionStmt(ctx *FunctionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
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

func (v *BaseCellScriptParserVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitTypeParameterDecl(ctx *TypeParameterDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCellScriptParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
