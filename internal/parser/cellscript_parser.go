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
		"'break'", "'import'", "'return'", "'vector'", "'table'", "'union'",
		"'var'", "'{'", "'}'", "'('", "')'", "','", "", "", "'int'", "'bool'",
		"'+'", "'-'", "'*'", "'/'", "'%'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "FUNC", "PACKAGE", "IF", "ELSE", "FOR", "CONTINUE", "BREAK", "IMPORT",
		"RETURN", "VECTOR", "TABLE", "UNION", "VAR", "L_CURLY", "R_CURLY", "L_BRACKET",
		"R_BRACKET", "COMMA", "IDENTIFIER", "TYPE", "TYPE_INT", "TYPE_BOOL",
		"ADD", "SUB", "MUL", "DIV", "MOD", "ASSIGN", "LITERAL", "NUMBER", "STRING_LIT",
		"BOOL_LIT", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT", "WS", "COMMENT",
		"TERMINATOR", "LINE_COMMENT", "WS_DEFAULT", "COMMENT_DEFAULT", "LINE_COMMENT_DEFAULT",
		"EOS",
	}
	staticData.RuleNames = []string{
		"sourceFile", "importStmt", "declaration", "eos", "importDecl", "varDecl",
		"functionDecl", "typeParameters", "signature", "block", "typeParameterDecl",
		"expression", "arithmeticExpr", "returnExpr", "statement", "simpleStmt",
		"expressionStmt", "assignment", "assign_op", "returnStmt", "breakStmt",
		"continueStmt", "ifStmt", "forStmt", "forClause",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 217, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 58, 8, 1, 10, 1, 12, 1, 61, 9, 1, 1, 2, 1,
		2, 3, 2, 65, 8, 2, 1, 2, 1, 2, 5, 2, 69, 8, 2, 10, 2, 12, 2, 72, 9, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 3, 6, 88, 8, 6, 1, 6, 1, 6, 3, 6, 92, 8, 6, 1, 7, 1, 7, 3,
		7, 96, 8, 7, 1, 7, 1, 7, 5, 7, 100, 8, 7, 10, 7, 12, 7, 103, 9, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 111, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 3, 11, 120, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 131, 8, 12, 10, 12, 12, 12, 134,
		9, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 3, 14, 147, 8, 14, 1, 15, 1, 15, 3, 15, 151, 8, 15, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 3, 18, 160, 8, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 3, 19, 166, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 3, 20, 172,
		8, 20, 1, 21, 1, 21, 3, 21, 176, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 187, 8, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 3, 22, 193, 8, 22, 3, 22, 195, 8, 22, 1, 23, 1, 23, 3, 23, 199,
		8, 23, 1, 23, 3, 23, 202, 8, 23, 1, 23, 1, 23, 1, 24, 3, 24, 207, 8, 24,
		1, 24, 1, 24, 3, 24, 211, 8, 24, 1, 24, 1, 24, 3, 24, 215, 8, 24, 1, 24,
		0, 1, 24, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 4, 1, 1, 42, 42, 1, 0, 23, 24, 1,
		0, 25, 26, 1, 0, 23, 27, 223, 0, 50, 1, 0, 0, 0, 2, 59, 1, 0, 0, 0, 4,
		70, 1, 0, 0, 0, 6, 75, 1, 0, 0, 0, 8, 77, 1, 0, 0, 0, 10, 80, 1, 0, 0,
		0, 12, 84, 1, 0, 0, 0, 14, 93, 1, 0, 0, 0, 16, 106, 1, 0, 0, 0, 18, 108,
		1, 0, 0, 0, 20, 114, 1, 0, 0, 0, 22, 119, 1, 0, 0, 0, 24, 121, 1, 0, 0,
		0, 26, 135, 1, 0, 0, 0, 28, 146, 1, 0, 0, 0, 30, 150, 1, 0, 0, 0, 32, 152,
		1, 0, 0, 0, 34, 154, 1, 0, 0, 0, 36, 159, 1, 0, 0, 0, 38, 163, 1, 0, 0,
		0, 40, 169, 1, 0, 0, 0, 42, 173, 1, 0, 0, 0, 44, 177, 1, 0, 0, 0, 46, 196,
		1, 0, 0, 0, 48, 206, 1, 0, 0, 0, 50, 51, 3, 2, 1, 0, 51, 52, 3, 4, 2, 0,
		52, 53, 5, 0, 0, 1, 53, 1, 1, 0, 0, 0, 54, 55, 3, 8, 4, 0, 55, 56, 3, 6,
		3, 0, 56, 58, 1, 0, 0, 0, 57, 54, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57,
		1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 3, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0,
		62, 65, 3, 12, 6, 0, 63, 65, 3, 10, 5, 0, 64, 62, 1, 0, 0, 0, 64, 63, 1,
		0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 3, 6, 3, 0, 67, 69, 1, 0, 0, 0, 68,
		64, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0,
		0, 71, 73, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 74, 5, 0, 0, 1, 74, 5, 1,
		0, 0, 0, 75, 76, 7, 0, 0, 0, 76, 7, 1, 0, 0, 0, 77, 78, 5, 8, 0, 0, 78,
		79, 5, 19, 0, 0, 79, 9, 1, 0, 0, 0, 80, 81, 5, 13, 0, 0, 81, 82, 5, 20,
		0, 0, 82, 83, 5, 19, 0, 0, 83, 11, 1, 0, 0, 0, 84, 85, 5, 1, 0, 0, 85,
		87, 5, 19, 0, 0, 86, 88, 3, 14, 7, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0,
		0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 3, 16, 8, 0, 90, 92, 3, 18, 9, 0, 91,
		90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 13, 1, 0, 0, 0, 93, 95, 5, 16,
		0, 0, 94, 96, 3, 20, 10, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96,
		101, 1, 0, 0, 0, 97, 98, 5, 18, 0, 0, 98, 100, 3, 20, 10, 0, 99, 97, 1,
		0, 0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0,
		0, 102, 104, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 105, 5, 17, 0, 0, 105,
		15, 1, 0, 0, 0, 106, 107, 5, 20, 0, 0, 107, 17, 1, 0, 0, 0, 108, 110, 5,
		14, 0, 0, 109, 111, 3, 22, 11, 0, 110, 109, 1, 0, 0, 0, 110, 111, 1, 0,
		0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 5, 15, 0, 0, 113, 19, 1, 0, 0, 0,
		114, 115, 5, 20, 0, 0, 115, 116, 5, 19, 0, 0, 116, 21, 1, 0, 0, 0, 117,
		120, 3, 24, 12, 0, 118, 120, 3, 26, 13, 0, 119, 117, 1, 0, 0, 0, 119, 118,
		1, 0, 0, 0, 120, 23, 1, 0, 0, 0, 121, 122, 6, 12, -1, 0, 122, 123, 5, 19,
		0, 0, 123, 132, 1, 0, 0, 0, 124, 125, 10, 3, 0, 0, 125, 126, 7, 1, 0, 0,
		126, 131, 3, 24, 12, 4, 127, 128, 10, 2, 0, 0, 128, 129, 7, 2, 0, 0, 129,
		131, 3, 24, 12, 3, 130, 124, 1, 0, 0, 0, 130, 127, 1, 0, 0, 0, 131, 134,
		1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 25, 1, 0,
		0, 0, 134, 132, 1, 0, 0, 0, 135, 136, 5, 9, 0, 0, 136, 137, 5, 19, 0, 0,
		137, 27, 1, 0, 0, 0, 138, 147, 3, 4, 2, 0, 139, 147, 3, 30, 15, 0, 140,
		147, 3, 38, 19, 0, 141, 147, 3, 40, 20, 0, 142, 147, 3, 42, 21, 0, 143,
		147, 3, 18, 9, 0, 144, 147, 3, 44, 22, 0, 145, 147, 3, 46, 23, 0, 146,
		138, 1, 0, 0, 0, 146, 139, 1, 0, 0, 0, 146, 140, 1, 0, 0, 0, 146, 141,
		1, 0, 0, 0, 146, 142, 1, 0, 0, 0, 146, 143, 1, 0, 0, 0, 146, 144, 1, 0,
		0, 0, 146, 145, 1, 0, 0, 0, 147, 29, 1, 0, 0, 0, 148, 151, 3, 34, 17, 0,
		149, 151, 3, 32, 16, 0, 150, 148, 1, 0, 0, 0, 150, 149, 1, 0, 0, 0, 151,
		31, 1, 0, 0, 0, 152, 153, 3, 22, 11, 0, 153, 33, 1, 0, 0, 0, 154, 155,
		3, 22, 11, 0, 155, 156, 3, 36, 18, 0, 156, 157, 3, 22, 11, 0, 157, 35,
		1, 0, 0, 0, 158, 160, 7, 3, 0, 0, 159, 158, 1, 0, 0, 0, 159, 160, 1, 0,
		0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 5, 28, 0, 0, 162, 37, 1, 0, 0, 0,
		163, 165, 5, 9, 0, 0, 164, 166, 3, 22, 11, 0, 165, 164, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 3, 6, 3, 0, 168, 39, 1,
		0, 0, 0, 169, 171, 5, 7, 0, 0, 170, 172, 5, 19, 0, 0, 171, 170, 1, 0, 0,
		0, 171, 172, 1, 0, 0, 0, 172, 41, 1, 0, 0, 0, 173, 175, 5, 6, 0, 0, 174,
		176, 5, 19, 0, 0, 175, 174, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 43,
		1, 0, 0, 0, 177, 186, 5, 3, 0, 0, 178, 187, 3, 22, 11, 0, 179, 180, 3,
		6, 3, 0, 180, 181, 3, 22, 11, 0, 181, 187, 1, 0, 0, 0, 182, 183, 3, 30,
		15, 0, 183, 184, 3, 6, 3, 0, 184, 185, 3, 22, 11, 0, 185, 187, 1, 0, 0,
		0, 186, 178, 1, 0, 0, 0, 186, 179, 1, 0, 0, 0, 186, 182, 1, 0, 0, 0, 187,
		188, 1, 0, 0, 0, 188, 194, 3, 18, 9, 0, 189, 192, 5, 4, 0, 0, 190, 193,
		3, 44, 22, 0, 191, 193, 3, 18, 9, 0, 192, 190, 1, 0, 0, 0, 192, 191, 1,
		0, 0, 0, 193, 195, 1, 0, 0, 0, 194, 189, 1, 0, 0, 0, 194, 195, 1, 0, 0,
		0, 195, 45, 1, 0, 0, 0, 196, 201, 5, 5, 0, 0, 197, 199, 3, 22, 11, 0, 198,
		197, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 202,
		3, 48, 24, 0, 201, 198, 1, 0, 0, 0, 201, 200, 1, 0, 0, 0, 202, 203, 1,
		0, 0, 0, 203, 204, 3, 18, 9, 0, 204, 47, 1, 0, 0, 0, 205, 207, 3, 30, 15,
		0, 206, 205, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208,
		210, 3, 6, 3, 0, 209, 211, 3, 22, 11, 0, 210, 209, 1, 0, 0, 0, 210, 211,
		1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 214, 3, 6, 3, 0, 213, 215, 3, 30,
		15, 0, 214, 213, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 49, 1, 0, 0, 0,
		25, 59, 64, 70, 87, 91, 95, 101, 110, 119, 130, 132, 146, 150, 159, 165,
		171, 175, 186, 192, 194, 198, 201, 206, 210, 214,
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
	CellScriptParserEOF                    = antlr.TokenEOF
	CellScriptParserFUNC                   = 1
	CellScriptParserPACKAGE                = 2
	CellScriptParserIF                     = 3
	CellScriptParserELSE                   = 4
	CellScriptParserFOR                    = 5
	CellScriptParserCONTINUE               = 6
	CellScriptParserBREAK                  = 7
	CellScriptParserIMPORT                 = 8
	CellScriptParserRETURN                 = 9
	CellScriptParserVECTOR                 = 10
	CellScriptParserTABLE                  = 11
	CellScriptParserUNION                  = 12
	CellScriptParserVAR                    = 13
	CellScriptParserL_CURLY                = 14
	CellScriptParserR_CURLY                = 15
	CellScriptParserL_BRACKET              = 16
	CellScriptParserR_BRACKET              = 17
	CellScriptParserCOMMA                  = 18
	CellScriptParserIDENTIFIER             = 19
	CellScriptParserTYPE                   = 20
	CellScriptParserTYPE_INT               = 21
	CellScriptParserTYPE_BOOL              = 22
	CellScriptParserADD                    = 23
	CellScriptParserSUB                    = 24
	CellScriptParserMUL                    = 25
	CellScriptParserDIV                    = 26
	CellScriptParserMOD                    = 27
	CellScriptParserASSIGN                 = 28
	CellScriptParserLITERAL                = 29
	CellScriptParserNUMBER                 = 30
	CellScriptParserSTRING_LIT             = 31
	CellScriptParserBOOL_LIT               = 32
	CellScriptParserRAW_STRING_LIT         = 33
	CellScriptParserINTERPRETED_STRING_LIT = 34
	CellScriptParserWS                     = 35
	CellScriptParserCOMMENT                = 36
	CellScriptParserTERMINATOR             = 37
	CellScriptParserLINE_COMMENT           = 38
	CellScriptParserWS_DEFAULT             = 39
	CellScriptParserCOMMENT_DEFAULT        = 40
	CellScriptParserLINE_COMMENT_DEFAULT   = 41
	CellScriptParserEOS                    = 42
)

// CellScriptParser rules.
const (
	CellScriptParserRULE_sourceFile        = 0
	CellScriptParserRULE_importStmt        = 1
	CellScriptParserRULE_declaration       = 2
	CellScriptParserRULE_eos               = 3
	CellScriptParserRULE_importDecl        = 4
	CellScriptParserRULE_varDecl           = 5
	CellScriptParserRULE_functionDecl      = 6
	CellScriptParserRULE_typeParameters    = 7
	CellScriptParserRULE_signature         = 8
	CellScriptParserRULE_block             = 9
	CellScriptParserRULE_typeParameterDecl = 10
	CellScriptParserRULE_expression        = 11
	CellScriptParserRULE_arithmeticExpr    = 12
	CellScriptParserRULE_returnExpr        = 13
	CellScriptParserRULE_statement         = 14
	CellScriptParserRULE_simpleStmt        = 15
	CellScriptParserRULE_expressionStmt    = 16
	CellScriptParserRULE_assignment        = 17
	CellScriptParserRULE_assign_op         = 18
	CellScriptParserRULE_returnStmt        = 19
	CellScriptParserRULE_breakStmt         = 20
	CellScriptParserRULE_continueStmt      = 21
	CellScriptParserRULE_ifStmt            = 22
	CellScriptParserRULE_forStmt           = 23
	CellScriptParserRULE_forClause         = 24
)

// ISourceFileContext is an interface to support dynamic dispatch.
type ISourceFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ImportStmt() IImportStmtContext
	Declaration() IDeclarationContext
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

func (s *SourceFileContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
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
		p.SetState(50)
		p.ImportStmt()
	}
	{
		p.SetState(51)
		p.Declaration()
	}
	{
		p.SetState(52)
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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CellScriptParserIMPORT {
		{
			p.SetState(54)
			p.ImportDecl()
		}
		{
			p.SetState(55)
			p.Eos()
		}

		p.SetState(61)
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

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllEos() []IEosContext
	Eos(i int) IEosContext
	AllFunctionDecl() []IFunctionDeclContext
	FunctionDecl(i int) IFunctionDeclContext
	AllVarDecl() []IVarDeclContext
	VarDecl(i int) IVarDeclContext

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

func (s *DeclarationContext) EOF() antlr.TerminalNode {
	return s.GetToken(CellScriptParserEOF, 0)
}

func (s *DeclarationContext) AllEos() []IEosContext {
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

func (s *DeclarationContext) Eos(i int) IEosContext {
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

func (s *DeclarationContext) AllFunctionDecl() []IFunctionDeclContext {
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

func (s *DeclarationContext) FunctionDecl(i int) IFunctionDeclContext {
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

func (s *DeclarationContext) AllVarDecl() []IVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclContext); ok {
			tst[i] = t.(IVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationContext) VarDecl(i int) IVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
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

	return t.(IVarDeclContext)
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
	p.EnterRule(localctx, 4, CellScriptParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CellScriptParserFUNC || _la == CellScriptParserVAR {
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case CellScriptParserFUNC:
			{
				p.SetState(62)
				p.FunctionDecl()
			}

		case CellScriptParserVAR:
			{
				p.SetState(63)
				p.VarDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(66)
			p.Eos()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(73)
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
		p.SetState(75)
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
		p.SetState(77)
		p.Match(CellScriptParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
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

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) VAR() antlr.TerminalNode {
	return s.GetToken(CellScriptParserVAR, 0)
}

func (s *VarDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(CellScriptParserTYPE, 0)
}

func (s *VarDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CellScriptParserIDENTIFIER, 0)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CellScriptParserRULE_varDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(CellScriptParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(CellScriptParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
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
	Block() IBlockContext

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

func (s *FunctionDeclContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(CellScriptParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(CellScriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserL_BRACKET {
		{
			p.SetState(86)
			p.TypeParameters()
		}

	}
	{
		p.SetState(89)
		p.Signature()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserL_CURLY {
		{
			p.SetState(90)
			p.Block()
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
		p.SetState(93)
		p.Match(CellScriptParserL_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserTYPE {
		{
			p.SetState(94)
			p.TypeParameterDecl()
		}

	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CellScriptParserCOMMA {
		{
			p.SetState(97)
			p.Match(CellScriptParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.TypeParameterDecl()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(104)
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
		p.SetState(106)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	Expression() IExpressionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(CellScriptParserL_CURLY, 0)
}

func (s *BlockContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(CellScriptParserR_CURLY, 0)
}

func (s *BlockContext) Expression() IExpressionContext {
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

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CellScriptParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(CellScriptParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserRETURN || _la == CellScriptParserIDENTIFIER {
		{
			p.SetState(109)
			p.Expression()
		}

	}
	{
		p.SetState(112)
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
		p.SetState(114)
		p.Match(CellScriptParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CellScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.arithmeticExpr(0)
		}

	case CellScriptParserRETURN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
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
	IDENTIFIER() antlr.TerminalNode
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

func (s *ArithmeticExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CellScriptParserIDENTIFIER, 0)
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
		p.SetState(122)
		p.Match(CellScriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CellScriptParserRULE_arithmeticExpr)
				p.SetState(124)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(125)

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
					p.SetState(126)
					p.arithmeticExpr(4)
				}

			case 2:
				localctx = NewArithmeticExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CellScriptParserRULE_arithmeticExpr)
				p.SetState(127)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(128)

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
					p.SetState(129)
					p.arithmeticExpr(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(CellScriptParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(136)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaration() IDeclarationContext
	SimpleStmt() ISimpleStmtContext
	ReturnStmt() IReturnStmtContext
	BreakStmt() IBreakStmtContext
	ContinueStmt() IContinueStmtContext
	Block() IBlockContext
	IfStmt() IIfStmtContext
	ForStmt() IForStmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) SimpleStmt() ISimpleStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *StatementContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StatementContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CellScriptParserRULE_statement)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.SimpleStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.ReturnStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(141)
			p.BreakStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(142)
			p.ContinueStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(143)
			p.Block()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(144)
			p.IfStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(145)
			p.ForStmt()
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

// ISimpleStmtContext is an interface to support dynamic dispatch.
type ISimpleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	ExpressionStmt() IExpressionStmtContext

	// IsSimpleStmtContext differentiates from other interfaces.
	IsSimpleStmtContext()
}

type SimpleStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStmtContext() *SimpleStmtContext {
	var p = new(SimpleStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_simpleStmt
	return p
}

func InitEmptySimpleStmtContext(p *SimpleStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_simpleStmt
}

func (*SimpleStmtContext) IsSimpleStmtContext() {}

func NewSimpleStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStmtContext {
	var p = new(SimpleStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_simpleStmt

	return p
}

func (s *SimpleStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStmtContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *SimpleStmtContext) ExpressionStmt() IExpressionStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStmtContext)
}

func (s *SimpleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitSimpleStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) SimpleStmt() (localctx ISimpleStmtContext) {
	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CellScriptParserRULE_simpleStmt)
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.ExpressionStmt()
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

// IExpressionStmtContext is an interface to support dynamic dispatch.
type IExpressionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsExpressionStmtContext differentiates from other interfaces.
	IsExpressionStmtContext()
}

type ExpressionStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStmtContext() *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_expressionStmt
	return p
}

func InitEmptyExpressionStmtContext(p *ExpressionStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_expressionStmt
}

func (*ExpressionStmtContext) IsExpressionStmtContext() {}

func NewExpressionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_expressionStmt

	return p
}

func (s *ExpressionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStmtContext) Expression() IExpressionContext {
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

func (s *ExpressionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterExpressionStmt(s)
	}
}

func (s *ExpressionStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitExpressionStmt(s)
	}
}

func (s *ExpressionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitExpressionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) ExpressionStmt() (localctx IExpressionStmtContext) {
	localctx = NewExpressionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CellScriptParserRULE_expressionStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Expression()
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	Assign_op() IAssign_opContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *AssignmentContext) Assign_op() IAssign_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_opContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CellScriptParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Expression()
	}
	{
		p.SetState(155)
		p.Assign_op()
	}
	{
		p.SetState(156)
		p.Expression()
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

// IAssign_opContext is an interface to support dynamic dispatch.
type IAssign_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode

	// IsAssign_opContext differentiates from other interfaces.
	IsAssign_opContext()
}

type Assign_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_opContext() *Assign_opContext {
	var p = new(Assign_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_assign_op
	return p
}

func InitEmptyAssign_opContext(p *Assign_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_assign_op
}

func (*Assign_opContext) IsAssign_opContext() {}

func NewAssign_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_opContext {
	var p = new(Assign_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_assign_op

	return p
}

func (s *Assign_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_opContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CellScriptParserASSIGN, 0)
}

func (s *Assign_opContext) ADD() antlr.TerminalNode {
	return s.GetToken(CellScriptParserADD, 0)
}

func (s *Assign_opContext) SUB() antlr.TerminalNode {
	return s.GetToken(CellScriptParserSUB, 0)
}

func (s *Assign_opContext) MUL() antlr.TerminalNode {
	return s.GetToken(CellScriptParserMUL, 0)
}

func (s *Assign_opContext) DIV() antlr.TerminalNode {
	return s.GetToken(CellScriptParserDIV, 0)
}

func (s *Assign_opContext) MOD() antlr.TerminalNode {
	return s.GetToken(CellScriptParserMOD, 0)
}

func (s *Assign_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterAssign_op(s)
	}
}

func (s *Assign_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitAssign_op(s)
	}
}

func (s *Assign_opContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitAssign_op(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) Assign_op() (localctx IAssign_opContext) {
	localctx = NewAssign_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CellScriptParserRULE_assign_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&260046848) != 0 {
		{
			p.SetState(158)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&260046848) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(161)
		p.Match(CellScriptParserASSIGN)
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

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Eos() IEosContext
	Expression() IExpressionContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(CellScriptParserRETURN, 0)
}

func (s *ReturnStmtContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ReturnStmtContext) Expression() IExpressionContext {
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

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CellScriptParserRULE_returnStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(CellScriptParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserRETURN || _la == CellScriptParserIDENTIFIER {
		{
			p.SetState(164)
			p.Expression()
		}

	}
	{
		p.SetState(167)
		p.Eos()
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

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(CellScriptParserBREAK, 0)
}

func (s *BreakStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CellScriptParserIDENTIFIER, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CellScriptParserRULE_breakStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(CellScriptParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserIDENTIFIER {
		{
			p.SetState(170)
			p.Match(CellScriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(CellScriptParserCONTINUE, 0)
}

func (s *ContinueStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CellScriptParserIDENTIFIER, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CellScriptParserRULE_continueStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(CellScriptParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserIDENTIFIER {
		{
			p.SetState(174)
			p.Match(CellScriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	Expression() IExpressionContext
	Eos() IEosContext
	SimpleStmt() ISimpleStmtContext
	ELSE() antlr.TerminalNode
	IfStmt() IIfStmtContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(CellScriptParserIF, 0)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfStmtContext) Expression() IExpressionContext {
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

func (s *IfStmtContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *IfStmtContext) SimpleStmt() ISimpleStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(CellScriptParserELSE, 0)
}

func (s *IfStmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CellScriptParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(CellScriptParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(178)
			p.Expression()
		}

	case 2:
		{
			p.SetState(179)
			p.Eos()
		}
		{
			p.SetState(180)
			p.Expression()
		}

	case 3:
		{
			p.SetState(182)
			p.SimpleStmt()
		}
		{
			p.SetState(183)
			p.Eos()
		}
		{
			p.SetState(184)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(188)
		p.Block()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserELSE {
		{
			p.SetState(189)
			p.Match(CellScriptParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case CellScriptParserIF:
			{
				p.SetState(190)
				p.IfStmt()
			}

		case CellScriptParserL_CURLY:
			{
				p.SetState(191)
				p.Block()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	Block() IBlockContext
	ForClause() IForClauseContext
	Expression() IExpressionContext

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(CellScriptParserFOR, 0)
}

func (s *ForStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStmtContext) ForClause() IForClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForClauseContext)
}

func (s *ForStmtContext) Expression() IExpressionContext {
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

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CellScriptParserRULE_forStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(CellScriptParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CellScriptParserRETURN || _la == CellScriptParserIDENTIFIER {
			{
				p.SetState(197)
				p.Expression()
			}

		}

	case 2:
		{
			p.SetState(200)
			p.ForClause()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(203)
		p.Block()
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

// IForClauseContext is an interface to support dynamic dispatch.
type IForClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInitStmt returns the initStmt rule contexts.
	GetInitStmt() ISimpleStmtContext

	// GetPostStmt returns the postStmt rule contexts.
	GetPostStmt() ISimpleStmtContext

	// SetInitStmt sets the initStmt rule contexts.
	SetInitStmt(ISimpleStmtContext)

	// SetPostStmt sets the postStmt rule contexts.
	SetPostStmt(ISimpleStmtContext)

	// Getter signatures
	AllEos() []IEosContext
	Eos(i int) IEosContext
	Expression() IExpressionContext
	AllSimpleStmt() []ISimpleStmtContext
	SimpleStmt(i int) ISimpleStmtContext

	// IsForClauseContext differentiates from other interfaces.
	IsForClauseContext()
}

type ForClauseContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	initStmt ISimpleStmtContext
	postStmt ISimpleStmtContext
}

func NewEmptyForClauseContext() *ForClauseContext {
	var p = new(ForClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_forClause
	return p
}

func InitEmptyForClauseContext(p *ForClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CellScriptParserRULE_forClause
}

func (*ForClauseContext) IsForClauseContext() {}

func NewForClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForClauseContext {
	var p = new(ForClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CellScriptParserRULE_forClause

	return p
}

func (s *ForClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ForClauseContext) GetInitStmt() ISimpleStmtContext { return s.initStmt }

func (s *ForClauseContext) GetPostStmt() ISimpleStmtContext { return s.postStmt }

func (s *ForClauseContext) SetInitStmt(v ISimpleStmtContext) { s.initStmt = v }

func (s *ForClauseContext) SetPostStmt(v ISimpleStmtContext) { s.postStmt = v }

func (s *ForClauseContext) AllEos() []IEosContext {
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

func (s *ForClauseContext) Eos(i int) IEosContext {
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

func (s *ForClauseContext) Expression() IExpressionContext {
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

func (s *ForClauseContext) AllSimpleStmt() []ISimpleStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStmtContext); ok {
			tst[i] = t.(ISimpleStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForClauseContext) SimpleStmt(i int) ISimpleStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
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

	return t.(ISimpleStmtContext)
}

func (s *ForClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.EnterForClause(s)
	}
}

func (s *ForClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CellScriptParserListener); ok {
		listenerT.ExitForClause(s)
	}
}

func (s *ForClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CellScriptParserVisitor:
		return t.VisitForClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CellScriptParser) ForClause() (localctx IForClauseContext) {
	localctx = NewForClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CellScriptParserRULE_forClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserRETURN || _la == CellScriptParserIDENTIFIER {
		{
			p.SetState(205)

			var _x = p.SimpleStmt()

			localctx.(*ForClauseContext).initStmt = _x
		}

	}
	{
		p.SetState(208)
		p.Eos()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserRETURN || _la == CellScriptParserIDENTIFIER {
		{
			p.SetState(209)
			p.Expression()
		}

	}
	{
		p.SetState(212)
		p.Eos()
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CellScriptParserRETURN || _la == CellScriptParserIDENTIFIER {
		{
			p.SetState(213)

			var _x = p.SimpleStmt()

			localctx.(*ForClauseContext).postStmt = _x
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
