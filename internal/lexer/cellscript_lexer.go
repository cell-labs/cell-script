// Code generated from CellScriptLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package lexer

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CellScriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CellScriptLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cellscriptlexerLexerInit() {
	staticData := &CellScriptLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"FUNC", "PACKAGE", "IF", "ELSE", "FOR", "CONTINUE", "BREAK", "IMPORT",
		"RETURN", "MAIN", "WHITESPACE", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET",
		"COMMA", "IDENTIFIER", "TYPE", "TYPE_INT", "TYPE_BOOL", "ADD", "SUB",
		"MUL", "DIV", "LITERAL", "NUMBER", "STRING_LIT", "BOOL_LIT", "EOS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 212, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 4, 10,
		124, 8, 10, 11, 10, 12, 10, 125, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 5, 16, 142,
		8, 16, 10, 16, 12, 16, 145, 9, 16, 1, 17, 1, 17, 3, 17, 149, 8, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 3, 24, 171,
		8, 24, 1, 25, 4, 25, 174, 8, 25, 11, 25, 12, 25, 175, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		3, 27, 191, 8, 27, 1, 28, 4, 28, 194, 8, 28, 11, 28, 12, 28, 195, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 203, 8, 28, 10, 28, 12, 28, 206, 9,
		28, 1, 28, 1, 28, 1, 28, 3, 28, 211, 8, 28, 1, 204, 0, 29, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 1, 0, 5, 3,
		0, 9, 10, 13, 13, 32, 32, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 223, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3, 68, 1, 0, 0, 0, 5,
		76, 1, 0, 0, 0, 7, 79, 1, 0, 0, 0, 9, 84, 1, 0, 0, 0, 11, 88, 1, 0, 0,
		0, 13, 97, 1, 0, 0, 0, 15, 103, 1, 0, 0, 0, 17, 110, 1, 0, 0, 0, 19, 117,
		1, 0, 0, 0, 21, 123, 1, 0, 0, 0, 23, 129, 1, 0, 0, 0, 25, 131, 1, 0, 0,
		0, 27, 133, 1, 0, 0, 0, 29, 135, 1, 0, 0, 0, 31, 137, 1, 0, 0, 0, 33, 139,
		1, 0, 0, 0, 35, 148, 1, 0, 0, 0, 37, 150, 1, 0, 0, 0, 39, 154, 1, 0, 0,
		0, 41, 159, 1, 0, 0, 0, 43, 161, 1, 0, 0, 0, 45, 163, 1, 0, 0, 0, 47, 165,
		1, 0, 0, 0, 49, 170, 1, 0, 0, 0, 51, 173, 1, 0, 0, 0, 53, 177, 1, 0, 0,
		0, 55, 190, 1, 0, 0, 0, 57, 210, 1, 0, 0, 0, 59, 60, 5, 102, 0, 0, 60,
		61, 5, 117, 0, 0, 61, 62, 5, 110, 0, 0, 62, 63, 5, 99, 0, 0, 63, 64, 5,
		116, 0, 0, 64, 65, 5, 105, 0, 0, 65, 66, 5, 111, 0, 0, 66, 67, 5, 110,
		0, 0, 67, 2, 1, 0, 0, 0, 68, 69, 5, 112, 0, 0, 69, 70, 5, 97, 0, 0, 70,
		71, 5, 99, 0, 0, 71, 72, 5, 107, 0, 0, 72, 73, 5, 97, 0, 0, 73, 74, 5,
		103, 0, 0, 74, 75, 5, 101, 0, 0, 75, 4, 1, 0, 0, 0, 76, 77, 5, 105, 0,
		0, 77, 78, 5, 102, 0, 0, 78, 6, 1, 0, 0, 0, 79, 80, 5, 101, 0, 0, 80, 81,
		5, 108, 0, 0, 81, 82, 5, 115, 0, 0, 82, 83, 5, 101, 0, 0, 83, 8, 1, 0,
		0, 0, 84, 85, 5, 102, 0, 0, 85, 86, 5, 111, 0, 0, 86, 87, 5, 114, 0, 0,
		87, 10, 1, 0, 0, 0, 88, 89, 5, 99, 0, 0, 89, 90, 5, 111, 0, 0, 90, 91,
		5, 110, 0, 0, 91, 92, 5, 116, 0, 0, 92, 93, 5, 105, 0, 0, 93, 94, 5, 110,
		0, 0, 94, 95, 5, 117, 0, 0, 95, 96, 5, 101, 0, 0, 96, 12, 1, 0, 0, 0, 97,
		98, 5, 98, 0, 0, 98, 99, 5, 114, 0, 0, 99, 100, 5, 101, 0, 0, 100, 101,
		5, 97, 0, 0, 101, 102, 5, 107, 0, 0, 102, 14, 1, 0, 0, 0, 103, 104, 5,
		105, 0, 0, 104, 105, 5, 109, 0, 0, 105, 106, 5, 112, 0, 0, 106, 107, 5,
		111, 0, 0, 107, 108, 5, 114, 0, 0, 108, 109, 5, 116, 0, 0, 109, 16, 1,
		0, 0, 0, 110, 111, 5, 114, 0, 0, 111, 112, 5, 101, 0, 0, 112, 113, 5, 116,
		0, 0, 113, 114, 5, 117, 0, 0, 114, 115, 5, 114, 0, 0, 115, 116, 5, 110,
		0, 0, 116, 18, 1, 0, 0, 0, 117, 118, 5, 109, 0, 0, 118, 119, 5, 97, 0,
		0, 119, 120, 5, 105, 0, 0, 120, 121, 5, 110, 0, 0, 121, 20, 1, 0, 0, 0,
		122, 124, 7, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125,
		123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128,
		6, 10, 0, 0, 128, 22, 1, 0, 0, 0, 129, 130, 5, 123, 0, 0, 130, 24, 1, 0,
		0, 0, 131, 132, 5, 125, 0, 0, 132, 26, 1, 0, 0, 0, 133, 134, 5, 40, 0,
		0, 134, 28, 1, 0, 0, 0, 135, 136, 5, 41, 0, 0, 136, 30, 1, 0, 0, 0, 137,
		138, 5, 44, 0, 0, 138, 32, 1, 0, 0, 0, 139, 143, 7, 1, 0, 0, 140, 142,
		7, 2, 0, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0,
		0, 0, 143, 144, 1, 0, 0, 0, 144, 34, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0,
		146, 149, 3, 37, 18, 0, 147, 149, 3, 39, 19, 0, 148, 146, 1, 0, 0, 0, 148,
		147, 1, 0, 0, 0, 149, 36, 1, 0, 0, 0, 150, 151, 5, 105, 0, 0, 151, 152,
		5, 110, 0, 0, 152, 153, 5, 116, 0, 0, 153, 38, 1, 0, 0, 0, 154, 155, 5,
		98, 0, 0, 155, 156, 5, 111, 0, 0, 156, 157, 5, 111, 0, 0, 157, 158, 5,
		108, 0, 0, 158, 40, 1, 0, 0, 0, 159, 160, 5, 43, 0, 0, 160, 42, 1, 0, 0,
		0, 161, 162, 5, 45, 0, 0, 162, 44, 1, 0, 0, 0, 163, 164, 5, 42, 0, 0, 164,
		46, 1, 0, 0, 0, 165, 166, 5, 47, 0, 0, 166, 48, 1, 0, 0, 0, 167, 171, 3,
		53, 26, 0, 168, 171, 3, 55, 27, 0, 169, 171, 3, 51, 25, 0, 170, 167, 1,
		0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 169, 1, 0, 0, 0, 171, 50, 1, 0, 0,
		0, 172, 174, 7, 3, 0, 0, 173, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175,
		173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 52, 1, 0, 0, 0, 177, 178, 5,
		34, 0, 0, 178, 179, 3, 33, 16, 0, 179, 180, 5, 34, 0, 0, 180, 54, 1, 0,
		0, 0, 181, 182, 5, 116, 0, 0, 182, 183, 5, 114, 0, 0, 183, 184, 5, 117,
		0, 0, 184, 191, 5, 101, 0, 0, 185, 186, 5, 102, 0, 0, 186, 187, 5, 97,
		0, 0, 187, 188, 5, 108, 0, 0, 188, 189, 5, 115, 0, 0, 189, 191, 5, 101,
		0, 0, 190, 181, 1, 0, 0, 0, 190, 185, 1, 0, 0, 0, 191, 56, 1, 0, 0, 0,
		192, 194, 7, 4, 0, 0, 193, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195,
		193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 211, 1, 0, 0, 0, 197, 211,
		5, 59, 0, 0, 198, 199, 5, 47, 0, 0, 199, 200, 5, 42, 0, 0, 200, 204, 1,
		0, 0, 0, 201, 203, 9, 0, 0, 0, 202, 201, 1, 0, 0, 0, 203, 206, 1, 0, 0,
		0, 204, 205, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 205, 207, 1, 0, 0, 0, 206,
		204, 1, 0, 0, 0, 207, 208, 5, 42, 0, 0, 208, 211, 5, 47, 0, 0, 209, 211,
		5, 0, 0, 1, 210, 193, 1, 0, 0, 0, 210, 197, 1, 0, 0, 0, 210, 198, 1, 0,
		0, 0, 210, 209, 1, 0, 0, 0, 211, 58, 1, 0, 0, 0, 10, 0, 125, 143, 148,
		170, 175, 190, 195, 204, 210, 1, 6, 0, 0,
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

// CellScriptLexerInit initializes any static state used to implement CellScriptLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCellScriptLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CellScriptLexerInit() {
	staticData := &CellScriptLexerLexerStaticData
	staticData.once.Do(cellscriptlexerLexerInit)
}

// NewCellScriptLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCellScriptLexer(input antlr.CharStream) *CellScriptLexer {
	CellScriptLexerInit()
	l := new(CellScriptLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CellScriptLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "CellScriptLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CellScriptLexer tokens.
const (
	CellScriptLexerFUNC       = 1
	CellScriptLexerPACKAGE    = 2
	CellScriptLexerIF         = 3
	CellScriptLexerELSE       = 4
	CellScriptLexerFOR        = 5
	CellScriptLexerCONTINUE   = 6
	CellScriptLexerBREAK      = 7
	CellScriptLexerIMPORT     = 8
	CellScriptLexerRETURN     = 9
	CellScriptLexerMAIN       = 10
	CellScriptLexerWHITESPACE = 11
	CellScriptLexerL_CURLY    = 12
	CellScriptLexerR_CURLY    = 13
	CellScriptLexerL_BRACKET  = 14
	CellScriptLexerR_BRACKET  = 15
	CellScriptLexerCOMMA      = 16
	CellScriptLexerIDENTIFIER = 17
	CellScriptLexerTYPE       = 18
	CellScriptLexerTYPE_INT   = 19
	CellScriptLexerTYPE_BOOL  = 20
	CellScriptLexerADD        = 21
	CellScriptLexerSUB        = 22
	CellScriptLexerMUL        = 23
	CellScriptLexerDIV        = 24
	CellScriptLexerLITERAL    = 25
	CellScriptLexerNUMBER     = 26
	CellScriptLexerSTRING_LIT = 27
	CellScriptLexerBOOL_LIT   = 28
	CellScriptLexerEOS        = 29
)
