// Code generated from CellScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CellScriptParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CellScriptParser struct {
	*antlr.BaseParser
}

var CellScriptParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cellscriptparserParserInit() {
	staticData := &CellScriptParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'function'", "'package'", "'if'", "'else'", "'for'", "'continue'",
		"'break'", "'import'", "'return'", "'main'", "", "'{'", "'}'", "'('",
		"')'", "','", "", "", "'int'", "'bool'", "'+'", "'-'", "'*'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "FUNC", "PACKAGE", "IF", "ELSE", "FOR", "CONTINUE", "BREAK", "IMPORT",
		"RETURN", "MAIN", "WHITESPACE", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET",
		"COMMA", "IDENTIFIER", "TYPE", "TYPE_INT", "TYPE_BOOL", "ADD", "SUB",
		"MUL", "DIV", "LITERAL", "NUMBER", "STRING_LIT", "BOOL_LIT", "EOS",
	}
	staticData.RuleNames = []string{
		"sourceFile", "importStmt", "functionStmt", "eos", "importDecl", "declaration",
		"functionDecl", "typeParameters", "signature", "body", "typeParameterDecl",
		"expression", "arithmeticExpr", "returnExpr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 126, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 5, 1, 36, 8, 1, 10, 1, 12, 1, 39, 9, 1, 1, 2, 1, 2, 3, 2,
		43, 8, 2, 1, 2, 1, 2, 5, 2, 47, 8, 2, 10, 2, 12, 2, 50, 9, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 3,
		6, 65, 8, 6, 1, 6, 1, 6, 3, 6, 69, 8, 6, 1, 6, 1, 6, 1, 6, 3, 6, 74, 8,
		6, 1, 6, 3, 6, 77, 8, 6, 3, 6, 79, 8, 6, 1, 7, 1, 7, 3, 7, 83, 8, 7, 1,
		7, 1, 7, 5, 7, 87, 8, 7, 10, 7, 12, 7, 90, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 3, 9, 98, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 3, 11, 107, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 5, 12, 118, 8, 12, 10, 12, 12, 12, 121, 9, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 0, 1, 24, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 0, 4, 1, 1, 29, 29, 1, 0, 21, 22, 1, 0, 23, 24, 2, 0, 17, 17,
		25, 25, 125, 0, 28, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 48, 1, 0, 0, 0, 6,
		53, 1, 0, 0, 0, 8, 55, 1, 0, 0, 0, 10, 58, 1, 0, 0, 0, 12, 78, 1, 0, 0,
		0, 14, 80, 1, 0, 0, 0, 16, 93, 1, 0, 0, 0, 18, 95, 1, 0, 0, 0, 20, 101,
		1, 0, 0, 0, 22, 106, 1, 0, 0, 0, 24, 108, 1, 0, 0, 0, 26, 122, 1, 0, 0,
		0, 28, 29, 3, 2, 1, 0, 29, 30, 3, 4, 2, 0, 30, 31, 5, 0, 0, 1, 31, 1, 1,
		0, 0, 0, 32, 33, 3, 8, 4, 0, 33, 34, 3, 6, 3, 0, 34, 36, 1, 0, 0, 0, 35,
		32, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0,
		0, 38, 3, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 43, 3, 12, 6, 0, 41, 43,
		3, 10, 5, 0, 42, 40, 1, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0,
		44, 45, 3, 6, 3, 0, 45, 47, 1, 0, 0, 0, 46, 42, 1, 0, 0, 0, 47, 50, 1,
		0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50,
		48, 1, 0, 0, 0, 51, 52, 5, 0, 0, 1, 52, 5, 1, 0, 0, 0, 53, 54, 7, 0, 0,
		0, 54, 7, 1, 0, 0, 0, 55, 56, 5, 8, 0, 0, 56, 57, 5, 17, 0, 0, 57, 9, 1,
		0, 0, 0, 58, 59, 5, 18, 0, 0, 59, 60, 5, 17, 0, 0, 60, 11, 1, 0, 0, 0,
		61, 62, 5, 1, 0, 0, 62, 64, 5, 17, 0, 0, 63, 65, 3, 14, 7, 0, 64, 63, 1,
		0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 3, 16, 8, 0, 67,
		69, 3, 18, 9, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 79, 1, 0,
		0, 0, 70, 71, 5, 1, 0, 0, 71, 73, 5, 10, 0, 0, 72, 74, 3, 14, 7, 0, 73,
		72, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 77, 3, 18,
		9, 0, 76, 75, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 79, 1, 0, 0, 0, 78, 61,
		1, 0, 0, 0, 78, 70, 1, 0, 0, 0, 79, 13, 1, 0, 0, 0, 80, 82, 5, 14, 0, 0,
		81, 83, 3, 20, 10, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 88, 1,
		0, 0, 0, 84, 85, 5, 16, 0, 0, 85, 87, 3, 20, 10, 0, 86, 84, 1, 0, 0, 0,
		87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 1,
		0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 92, 5, 15, 0, 0, 92, 15, 1, 0, 0, 0, 93,
		94, 5, 18, 0, 0, 94, 17, 1, 0, 0, 0, 95, 97, 5, 12, 0, 0, 96, 98, 3, 22,
		11, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99,
		100, 5, 13, 0, 0, 100, 19, 1, 0, 0, 0, 101, 102, 5, 18, 0, 0, 102, 103,
		5, 17, 0, 0, 103, 21, 1, 0, 0, 0, 104, 107, 3, 24, 12, 0, 105, 107, 3,
		26, 13, 0, 106, 104, 1, 0, 0, 0, 106, 105, 1, 0, 0, 0, 107, 23, 1, 0, 0,
		0, 108, 109, 6, 12, -1, 0, 109, 110, 5, 26, 0, 0, 110, 119, 1, 0, 0, 0,
		111, 112, 10, 3, 0, 0, 112, 113, 7, 1, 0, 0, 113, 118, 3, 24, 12, 4, 114,
		115, 10, 2, 0, 0, 115, 116, 7, 2, 0, 0, 116, 118, 3, 24, 12, 3, 117, 111,
		1, 0, 0, 0, 117, 114, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0,
		0, 0, 119, 120, 1, 0, 0, 0, 120, 25, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0,
		122, 123, 5, 9, 0, 0, 123, 124, 7, 3, 0, 0, 124, 27, 1, 0, 0, 0, 14, 37,
		42, 48, 64, 68, 73, 76, 78, 82, 88, 97, 106, 117, 119,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CellScriptParserInit initializes any static state used to implement CellScriptParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCellScriptParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CellScriptParserInit() {
	staticData := &CellScriptParserParserStaticData
	staticData.once.Do(cellscriptparserParserInit)
}

// NewCellScriptParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCellScriptParser(input antlr.TokenStream) *CellScriptParser {
	CellScriptParserInit()
	this := new(CellScriptParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CellScriptParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CellScriptParser.g4"

	return this
}

// CellScriptParser tokens.
const (
	CellScriptParserEOF        = antlr.TokenEOF
	CellScriptParserFUNC       = 1
	CellScriptParserPACKAGE    = 2
	CellScriptParserIF         = 3
	CellScriptParserELSE       = 4
	CellScriptParserFOR        = 5
	CellScriptParserCONTINUE   = 6
	CellScriptParserBREAK      = 7
	CellScriptParserIMPORT     = 8
	CellScriptParserRETURN     = 9
	CellScriptParserMAIN       = 10
	CellScriptParserWHITESPACE = 11
	CellScriptParserL_CURLY    = 12
	CellScriptParserR_CURLY    = 13
	CellScriptParserL_BRACKET  = 14
	CellScriptParserR_BRACKET  = 15
	CellScriptParserCOMMA      = 16
	CellScriptParserIDENTIFIER = 17
	CellScriptParserTYPE       = 18
	CellScriptParserTYPE_INT   = 19
	CellScriptParserTYPE_BOOL  = 20
	CellScriptParserADD        = 21
	CellScriptParserSUB        = 22
	CellScriptParserMUL        = 23
	CellScriptParserDIV        = 24
	CellScriptParserLITERAL    = 25
	CellScriptParserNUMBER     = 26
	CellScriptParserSTRING_LIT = 27
	CellScriptParserBOOL_LIT   = 28
	CellScriptParserEOS        = 29
)

// CellScriptParser rules.
const (
	CellScriptParserRULE_sourceFile        = 0
	CellScriptParserRULE_importStmt        = 1
	CellScriptParserRULE_functionStmt      = 2
	CellScriptParserRULE_eos               = 3
	CellScriptParserRULE_importDecl        = 4
	CellScriptParserRULE_declaration       = 5
	CellScriptParserRULE_functionDecl      = 6
	CellScriptParserRULE_typeParameters    = 7
	CellScriptParserRULE_signature         = 8
	CellScriptParserRULE_body              = 9
	CellScriptParserRULE_typeParameterDecl = 10
	CellScriptParserRULE_expression        = 11
	CellScriptParserRULE_arithmeticExpr    = 12
	CellScriptParserRULE_returnExpr        = 13
)

// ISourceFileContext is an interface to support dynamic dispatch.
type ISourceFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ImportStmt() IImportStmtContext
	FunctionStmt() IFunctionStmtContext
	EOF() antlr.TerminalNode

	// IsSourceFileContext differentiates from other interfaces.
	IsSourceFileContext()
}

type SourceFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceFileContext() *SourceFileContext {
	var p = new(SourceFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_sourceFile
	return p
}

func InitEmptySourceFileContext(p *SourceFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_sourceFile
}

func (*SourceFileContext) IsSourceFileContext() {}

func NewSourceFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceFileContext {
	var p = new(SourceFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_sourceFile

	return p
}

func (s *SourceFileContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceFileContext) ImportStmt() IImportStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStmtContext)
}

func (s *SourceFileContext) FunctionStmt() IFunctionStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionStmtContext)
}

func (s *SourceFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(CellScriptParserEOF, 0)
}

func (s *SourceFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterSourceFile(s)
	}
}

func (s *SourceFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitSourceFile(s)
	}
}

func (s *SourceFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitSourceFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) SourceFile() (localctx ISourceFileContext) {
	localctx = NewSourceFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CellScriptParserRULE_sourceFile)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.ImportStmt()
	}
	{
		p.SetState(29)
		p.FunctionStmt()
	}
	{
		p.SetState(30)
		p.Match(CellScriptParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportStmtContext is an interface to support dynamic dispatch.
type IImportStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllImportDecl() []IImportDeclContext
	ImportDecl(i int) IImportDeclContext
	AllEos() []IEosContext
	Eos(i int) IEosContext

	// IsImportStmtContext differentiates from other interfaces.
	IsImportStmtContext()
}

type ImportStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStmtContext() *ImportStmtContext {
	var p = new(ImportStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_importStmt
	return p
}

func InitEmptyImportStmtContext(p *ImportStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_importStmt
}

func (*ImportStmtContext) IsImportStmtContext() {}

func NewImportStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStmtContext {
	var p = new(ImportStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_importStmt

	return p
}

func (s *ImportStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStmtContext) AllImportDecl() []IImportDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportDeclContext); ok {
			len++
		}
	}

	tst := make([]IImportDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportDeclContext); ok {
			tst[i] = t.(IImportDeclContext)
			i++
		}
	}

	return tst
}

func (s *ImportStmtContext) ImportDecl(i int) IImportDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *ImportStmtContext) AllEos() []IEosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEosContext); ok {
			len++
		}
	}

	tst := make([]IEosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEosContext); ok {
			tst[i] = t.(IEosContext)
			i++
		}
	}

	return tst
}

func (s *ImportStmtContext) Eos(i int) IEosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ImportStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterImportStmt(s)
	}
}

func (s *ImportStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitImportStmt(s)
	}
}

func (s *ImportStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitImportStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) ImportStmt() (localctx IImportStmtContext) {
	localctx = NewImportStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CellScriptParserRULE_importStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CellScriptParserIMPORT {
		{
			p.SetState(32)
			p.ImportDecl()
		}
		{
			p.SetState(33)
			p.Eos()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionStmtContext is an interface to support dynamic dispatch.
type IFunctionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllEos() []IEosContext
	Eos(i int) IEosContext
	AllFunctionDecl() []IFunctionDeclContext
	FunctionDecl(i int) IFunctionDeclContext
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsFunctionStmtContext differentiates from other interfaces.
	IsFunctionStmtContext()
}

type FunctionStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionStmtContext() *FunctionStmtContext {
	var p = new(FunctionStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_functionStmt
	return p
}

func InitEmptyFunctionStmtContext(p *FunctionStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_functionStmt
}

func (*FunctionStmtContext) IsFunctionStmtContext() {}

func NewFunctionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionStmtContext {
	var p = new(FunctionStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_functionStmt

	return p
}

func (s *FunctionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionStmtContext) EOF() antlr.TerminalNode {
	return s.GetToken(CellScriptParserEOF, 0)
}

func (s *FunctionStmtContext) AllEos() []IEosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEosContext); ok {
			len++
		}
	}

	tst := make([]IEosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEosContext); ok {
			tst[i] = t.(IEosContext)
			i++
		}
	}

	return tst
}

func (s *FunctionStmtContext) Eos(i int) IEosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *FunctionStmtContext) AllFunctionDecl() []IFunctionDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDeclContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDeclContext); ok {
			tst[i] = t.(IFunctionDeclContext)
			i++
		}
	}

	return tst
}

func (s *FunctionStmtContext) FunctionDecl(i int) IFunctionDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclContext)
}

func (s *FunctionStmtContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *FunctionStmtContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *FunctionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterFunctionStmt(s)
	}
}

func (s *FunctionStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitFunctionStmt(s)
	}
}

func (s *FunctionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitFunctionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) FunctionStmt() (localctx IFunctionStmtContext) {
	localctx = NewFunctionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CellScriptParserRULE_functionStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CellScriptParserFUNC || _la == CellScriptParserTYPE {
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case CellScriptParserFUNC:
			{
				p.SetState(40)
				p.FunctionDecl()
			}

		case CellScriptParserTYPE:
			{
				p.SetState(41)
				p.Declaration()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(44)
			p.Eos()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(51)
		p.Match(CellScriptParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	EOS() antlr.TerminalNode

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_eos
	return p
}

func InitEmptyEosContext(p *EosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_eos
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(CellScriptParserEOF, 0)
}

func (s *EosContext) EOS() antlr.TerminalNode {
	return s.GetToken(CellScriptParserEOS, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CellScriptParserRULE_eos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CellScriptParserEOF || _la == CellScriptParserEOS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportDeclContext is an interface to support dynamic dispatch.
type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}

type ImportDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclContext() *ImportDeclContext {
	var p = new(ImportDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_importDecl
	return p
}

func InitEmptyImportDeclContext(p *ImportDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_importDecl
}

func (*ImportDeclContext) IsImportDeclContext() {}

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext {
	var p = new(ImportDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_importDecl

	return p
}

func (s *ImportDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(CellScriptParserIMPORT, 0)
}

func (s *ImportDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CellScriptParserIDENTIFIER, 0)
}

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterImportDecl(s)
	}
}

func (s *ImportDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitImportDecl(s)
	}
}

func (s *ImportDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitImportDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) ImportDecl() (localctx IImportDeclContext) {
	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CellScriptParserRULE_importDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(CellScriptParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(CellScriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(CellScriptParserTYPE, 0)
}

func (s *DeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CellScriptParserIDENTIFIER, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CellScriptParserRULE_declaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(CellScriptParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(CellScriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDeclContext is an interface to support dynamic dispatch.
type IFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Signature() ISignatureContext
	TypeParameters() ITypeParametersContext
	Body() IBodyContext
	MAIN() antlr.TerminalNode

	// IsFunctionDeclContext differentiates from other interfaces.
	IsFunctionDeclContext()
}

type FunctionDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclContext() *FunctionDeclContext {
	var p = new(FunctionDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_functionDecl
	return p
}

func InitEmptyFunctionDeclContext(p *FunctionDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_functionDecl
}

func (*FunctionDeclContext) IsFunctionDeclContext() {}

func NewFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclContext {
	var p = new(FunctionDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_functionDecl

	return p
}

func (s *FunctionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclContext) FUNC() antlr.TerminalNode {
	return s.GetToken(CellScriptParserFUNC, 0)
}

func (s *FunctionDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CellScriptParserIDENTIFIER, 0)
}

func (s *FunctionDeclContext) Signature() ISignatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignatureContext)
}

func (s *FunctionDeclContext) TypeParameters() ITypeParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeParametersContext)
}

func (s *FunctionDeclContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *FunctionDeclContext) MAIN() antlr.TerminalNode {
	return s.GetToken(CellScriptParserMAIN, 0)
}

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterFunctionDecl(s)
	}
}

func (s *FunctionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitFunctionDecl(s)
	}
}

func (s *FunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitFunctionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) FunctionDecl() (localctx IFunctionDeclContext) {
	localctx = NewFunctionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CellScriptParserRULE_functionDecl)
	var _la int

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Match(CellScriptParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Match(CellScriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CellScriptParserL_BRACKET {
			{
				p.SetState(63)
				p.TypeParameters()
			}

		}
		{
			p.SetState(66)
			p.Signature()
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CellScriptParserL_CURLY {
			{
				p.SetState(67)
				p.Body()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Match(CellScriptParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(CellScriptParserMAIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CellScriptParserL_BRACKET {
			{
				p.SetState(72)
				p.TypeParameters()
			}

		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CellScriptParserL_CURLY {
			{
				p.SetState(75)
				p.Body()
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeParametersContext is an interface to support dynamic dispatch.
type ITypeParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BRACKET() antlr.TerminalNode
	R_BRACKET() antlr.TerminalNode
	AllTypeParameterDecl() []ITypeParameterDeclContext
	TypeParameterDecl(i int) ITypeParameterDeclContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsTypeParametersContext differentiates from other interfaces.
	IsTypeParametersContext()
}

type TypeParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParametersContext() *TypeParametersContext {
	var p = new(TypeParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_typeParameters
	return p
}

func InitEmptyTypeParametersContext(p *TypeParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_typeParameters
}

func (*TypeParametersContext) IsTypeParametersContext() {}

func NewTypeParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParametersContext {
	var p = new(TypeParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_typeParameters

	return p
}

func (s *TypeParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeParametersContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(CellScriptParserL_BRACKET, 0)
}

func (s *TypeParametersContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(CellScriptParserR_BRACKET, 0)
}

func (s *TypeParametersContext) AllTypeParameterDecl() []ITypeParameterDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeParameterDeclContext); ok {
			len++
		}
	}

	tst := make([]ITypeParameterDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeParameterDeclContext); ok {
			tst[i] = t.(ITypeParameterDeclContext)
			i++
		}
	}

	return tst
}

func (s *TypeParametersContext) TypeParameterDecl(i int) ITypeParameterDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeParameterDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeParameterDeclContext)
}

func (s *TypeParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CellScriptParserCOMMA)
}

func (s *TypeParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CellScriptParserCOMMA, i)
}

func (s *TypeParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterTypeParameters(s)
	}
}

func (s *TypeParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitTypeParameters(s)
	}
}

func (s *TypeParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitTypeParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) TypeParameters() (localctx ITypeParametersContext) {
	localctx = NewTypeParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CellScriptParserRULE_typeParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(CellScriptParserL_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserTYPE {
		{
			p.SetState(81)
			p.TypeParameterDecl()
		}

	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CellScriptParserCOMMA {
		{
			p.SetState(84)
			p.Match(CellScriptParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.TypeParameterDecl()
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
		p.Match(CellScriptParserR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISignatureContext is an interface to support dynamic dispatch.
type ISignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode

	// IsSignatureContext differentiates from other interfaces.
	IsSignatureContext()
}

type SignatureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignatureContext() *SignatureContext {
	var p = new(SignatureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_signature
	return p
}

func InitEmptySignatureContext(p *SignatureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_signature
}

func (*SignatureContext) IsSignatureContext() {}

func NewSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignatureContext {
	var p = new(SignatureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_signature

	return p
}

func (s *SignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *SignatureContext) TYPE() antlr.TerminalNode {
	return s.GetToken(CellScriptParserTYPE, 0)
}

func (s *SignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterSignature(s)
	}
}

func (s *SignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitSignature(s)
	}
}

func (s *SignatureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitSignature(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) Signature() (localctx ISignatureContext) {
	localctx = NewSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CellScriptParserRULE_signature)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(CellScriptParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	Expression() IExpressionContext

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_body
	return p
}

func InitEmptyBodyContext(p *BodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_body
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(CellScriptParserL_CURLY, 0)
}

func (s *BodyContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(CellScriptParserR_CURLY, 0)
}

func (s *BodyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitBody(s)
	}
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CellScriptParserRULE_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(CellScriptParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserRETURN || _la == CellScriptParserNUMBER {
		{
			p.SetState(96)
			p.Expression()
		}

	}
	{
		p.SetState(99)
		p.Match(CellScriptParserR_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeParameterDeclContext is an interface to support dynamic dispatch.
type ITypeParameterDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsTypeParameterDeclContext differentiates from other interfaces.
	IsTypeParameterDeclContext()
}

type TypeParameterDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParameterDeclContext() *TypeParameterDeclContext {
	var p = new(TypeParameterDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_typeParameterDecl
	return p
}

func InitEmptyTypeParameterDeclContext(p *TypeParameterDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_typeParameterDecl
}

func (*TypeParameterDeclContext) IsTypeParameterDeclContext() {}

func NewTypeParameterDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParameterDeclContext {
	var p = new(TypeParameterDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_typeParameterDecl

	return p
}

func (s *TypeParameterDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeParameterDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(CellScriptParserTYPE, 0)
}

func (s *TypeParameterDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CellScriptParserIDENTIFIER, 0)
}

func (s *TypeParameterDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParameterDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeParameterDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterTypeParameterDecl(s)
	}
}

func (s *TypeParameterDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitTypeParameterDecl(s)
	}
}

func (s *TypeParameterDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitTypeParameterDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) TypeParameterDecl() (localctx ITypeParameterDeclContext) {
	localctx = NewTypeParameterDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CellScriptParserRULE_typeParameterDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(CellScriptParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(CellScriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArithmeticExpr() IArithmeticExprContext
	ReturnExpr() IReturnExprContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ArithmeticExpr() IArithmeticExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticExprContext)
}

func (s *ExpressionContext) ReturnExpr() IReturnExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnExprContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CellScriptParserRULE_expression)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CellScriptParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.arithmeticExpr(0)
		}

	case CellScriptParserRETURN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.ReturnExpr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArithmeticExprContext is an interface to support dynamic dispatch.
type IArithmeticExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	NUMBER() antlr.TerminalNode
	AllArithmeticExpr() []IArithmeticExprContext
	ArithmeticExpr(i int) IArithmeticExprContext
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode

	// IsArithmeticExprContext differentiates from other interfaces.
	IsArithmeticExprContext()
}

type ArithmeticExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyArithmeticExprContext() *ArithmeticExprContext {
	var p = new(ArithmeticExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_arithmeticExpr
	return p
}

func InitEmptyArithmeticExprContext(p *ArithmeticExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_arithmeticExpr
}

func (*ArithmeticExprContext) IsArithmeticExprContext() {}

func NewArithmeticExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticExprContext {
	var p = new(ArithmeticExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_arithmeticExpr

	return p
}

func (s *ArithmeticExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticExprContext) GetOp() antlr.Token { return s.op }

func (s *ArithmeticExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ArithmeticExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CellScriptParserNUMBER, 0)
}

func (s *ArithmeticExprContext) AllArithmeticExpr() []IArithmeticExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmeticExprContext); ok {
			len++
		}
	}

	tst := make([]IArithmeticExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmeticExprContext); ok {
			tst[i] = t.(IArithmeticExprContext)
			i++
		}
	}

	return tst
}

func (s *ArithmeticExprContext) ArithmeticExpr(i int) IArithmeticExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticExprContext)
}

func (s *ArithmeticExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(CellScriptParserADD, 0)
}

func (s *ArithmeticExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(CellScriptParserSUB, 0)
}

func (s *ArithmeticExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(CellScriptParserMUL, 0)
}

func (s *ArithmeticExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(CellScriptParserDIV, 0)
}

func (s *ArithmeticExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithmeticExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterArithmeticExpr(s)
	}
}

func (s *ArithmeticExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitArithmeticExpr(s)
	}
}

func (s *ArithmeticExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitArithmeticExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) ArithmeticExpr() (localctx IArithmeticExprContext) {
	return p.arithmeticExpr(0)
}

func (p *CellScriptParser) arithmeticExpr(_p int) (localctx IArithmeticExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewArithmeticExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArithmeticExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, CellScriptParserRULE_arithmeticExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(CellScriptParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CellScriptParserRULE_arithmeticExpr)
				p.SetState(111)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(112)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CellScriptParserADD || _la == CellScriptParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(113)
					p.arithmeticExpr(4)
				}

			case 2:
				localctx = NewArithmeticExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CellScriptParserRULE_arithmeticExpr)
				p.SetState(114)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(115)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CellScriptParserMUL || _la == CellScriptParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(116)
					p.arithmeticExpr(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnExprContext is an interface to support dynamic dispatch.
type IReturnExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	LITERAL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsReturnExprContext differentiates from other interfaces.
	IsReturnExprContext()
}

type ReturnExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnExprContext() *ReturnExprContext {
	var p = new(ReturnExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_returnExpr
	return p
}

func InitEmptyReturnExprContext(p *ReturnExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_returnExpr
}

func (*ReturnExprContext) IsReturnExprContext() {}

func NewReturnExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnExprContext {
	var p = new(ReturnExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_returnExpr

	return p
}

func (s *ReturnExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnExprContext) RETURN() antlr.TerminalNode {
	return s.GetToken(CellScriptParserRETURN, 0)
}

func (s *ReturnExprContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(CellScriptParserLITERAL, 0)
}

func (s *ReturnExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CellScriptParserIDENTIFIER, 0)
}

func (s *ReturnExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterReturnExpr(s)
	}
}

func (s *ReturnExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitReturnExpr(s)
	}
}

func (s *ReturnExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitReturnExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) ReturnExpr() (localctx IReturnExprContext) {
	localctx = NewReturnExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CellScriptParserRULE_returnExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(CellScriptParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CellScriptParserIDENTIFIER || _la == CellScriptParserLITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *CellScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *ArithmeticExprContext = nil
		if localctx != nil {
			t = localctx.(*ArithmeticExprContext)
		}
		return p.ArithmeticExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CellScriptParser) ArithmeticExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
