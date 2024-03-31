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
		"'break'", "'import'", "'return'", "", "'{'", "'}'", "'('", "')'", "','",
		"", "", "", "'int'", "'bool'", "'+'", "'-'", "'*'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "FUNC", "PACKAGE", "IF", "ELSE", "FOR", "CONTINUE", "BREAK", "IMPORT",
		"RETURN", "WHITESPACE", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET",
		"COMMA", "NUMBER", "IDENTIFIER", "STRING", "TYPE", "BOOL", "ADD", "SUB",
		"MUL", "DIV", "EOS",
	}
	staticData.RuleNames = []string{
		"FUNC", "PACKAGE", "IF", "ELSE", "FOR", "CONTINUE", "BREAK", "IMPORT",
		"RETURN", "WHITESPACE", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET",
		"COMMA", "NUMBER", "IDENTIFIER", "STRING", "TYPE", "BOOL", "ADD", "SUB",
		"MUL", "DIV", "EOS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 179, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 4, 9, 111, 8, 9,
		11, 9, 12, 9, 112, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 4, 15, 128, 8, 15, 11, 15, 12, 15, 129,
		1, 16, 1, 16, 5, 16, 134, 8, 16, 10, 16, 12, 16, 137, 9, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 4,
		24, 161, 8, 24, 11, 24, 12, 24, 162, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		5, 24, 170, 8, 24, 10, 24, 12, 24, 173, 9, 24, 1, 24, 1, 24, 1, 24, 3,
		24, 178, 8, 24, 1, 171, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		1, 0, 5, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 186,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 60, 1, 0, 0, 0,
		5, 68, 1, 0, 0, 0, 7, 71, 1, 0, 0, 0, 9, 76, 1, 0, 0, 0, 11, 80, 1, 0,
		0, 0, 13, 89, 1, 0, 0, 0, 15, 95, 1, 0, 0, 0, 17, 102, 1, 0, 0, 0, 19,
		110, 1, 0, 0, 0, 21, 116, 1, 0, 0, 0, 23, 118, 1, 0, 0, 0, 25, 120, 1,
		0, 0, 0, 27, 122, 1, 0, 0, 0, 29, 124, 1, 0, 0, 0, 31, 127, 1, 0, 0, 0,
		33, 131, 1, 0, 0, 0, 35, 138, 1, 0, 0, 0, 37, 142, 1, 0, 0, 0, 39, 146,
		1, 0, 0, 0, 41, 151, 1, 0, 0, 0, 43, 153, 1, 0, 0, 0, 45, 155, 1, 0, 0,
		0, 47, 157, 1, 0, 0, 0, 49, 177, 1, 0, 0, 0, 51, 52, 5, 102, 0, 0, 52,
		53, 5, 117, 0, 0, 53, 54, 5, 110, 0, 0, 54, 55, 5, 99, 0, 0, 55, 56, 5,
		116, 0, 0, 56, 57, 5, 105, 0, 0, 57, 58, 5, 111, 0, 0, 58, 59, 5, 110,
		0, 0, 59, 2, 1, 0, 0, 0, 60, 61, 5, 112, 0, 0, 61, 62, 5, 97, 0, 0, 62,
		63, 5, 99, 0, 0, 63, 64, 5, 107, 0, 0, 64, 65, 5, 97, 0, 0, 65, 66, 5,
		103, 0, 0, 66, 67, 5, 101, 0, 0, 67, 4, 1, 0, 0, 0, 68, 69, 5, 105, 0,
		0, 69, 70, 5, 102, 0, 0, 70, 6, 1, 0, 0, 0, 71, 72, 5, 101, 0, 0, 72, 73,
		5, 108, 0, 0, 73, 74, 5, 115, 0, 0, 74, 75, 5, 101, 0, 0, 75, 8, 1, 0,
		0, 0, 76, 77, 5, 102, 0, 0, 77, 78, 5, 111, 0, 0, 78, 79, 5, 114, 0, 0,
		79, 10, 1, 0, 0, 0, 80, 81, 5, 99, 0, 0, 81, 82, 5, 111, 0, 0, 82, 83,
		5, 110, 0, 0, 83, 84, 5, 116, 0, 0, 84, 85, 5, 105, 0, 0, 85, 86, 5, 110,
		0, 0, 86, 87, 5, 117, 0, 0, 87, 88, 5, 101, 0, 0, 88, 12, 1, 0, 0, 0, 89,
		90, 5, 98, 0, 0, 90, 91, 5, 114, 0, 0, 91, 92, 5, 101, 0, 0, 92, 93, 5,
		97, 0, 0, 93, 94, 5, 107, 0, 0, 94, 14, 1, 0, 0, 0, 95, 96, 5, 105, 0,
		0, 96, 97, 5, 109, 0, 0, 97, 98, 5, 112, 0, 0, 98, 99, 5, 111, 0, 0, 99,
		100, 5, 114, 0, 0, 100, 101, 5, 116, 0, 0, 101, 16, 1, 0, 0, 0, 102, 103,
		5, 114, 0, 0, 103, 104, 5, 101, 0, 0, 104, 105, 5, 116, 0, 0, 105, 106,
		5, 117, 0, 0, 106, 107, 5, 114, 0, 0, 107, 108, 5, 110, 0, 0, 108, 18,
		1, 0, 0, 0, 109, 111, 7, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111, 112, 1, 0,
		0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0,
		114, 115, 6, 9, 0, 0, 115, 20, 1, 0, 0, 0, 116, 117, 5, 123, 0, 0, 117,
		22, 1, 0, 0, 0, 118, 119, 5, 125, 0, 0, 119, 24, 1, 0, 0, 0, 120, 121,
		5, 40, 0, 0, 121, 26, 1, 0, 0, 0, 122, 123, 5, 41, 0, 0, 123, 28, 1, 0,
		0, 0, 124, 125, 5, 44, 0, 0, 125, 30, 1, 0, 0, 0, 126, 128, 7, 1, 0, 0,
		127, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129,
		130, 1, 0, 0, 0, 130, 32, 1, 0, 0, 0, 131, 135, 7, 2, 0, 0, 132, 134, 7,
		3, 0, 0, 133, 132, 1, 0, 0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0,
		0, 135, 136, 1, 0, 0, 0, 136, 34, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138,
		139, 5, 34, 0, 0, 139, 140, 3, 33, 16, 0, 140, 141, 5, 34, 0, 0, 141, 36,
		1, 0, 0, 0, 142, 143, 5, 105, 0, 0, 143, 144, 5, 110, 0, 0, 144, 145, 5,
		116, 0, 0, 145, 38, 1, 0, 0, 0, 146, 147, 5, 98, 0, 0, 147, 148, 5, 111,
		0, 0, 148, 149, 5, 111, 0, 0, 149, 150, 5, 108, 0, 0, 150, 40, 1, 0, 0,
		0, 151, 152, 5, 43, 0, 0, 152, 42, 1, 0, 0, 0, 153, 154, 5, 45, 0, 0, 154,
		44, 1, 0, 0, 0, 155, 156, 5, 42, 0, 0, 156, 46, 1, 0, 0, 0, 157, 158, 5,
		47, 0, 0, 158, 48, 1, 0, 0, 0, 159, 161, 7, 4, 0, 0, 160, 159, 1, 0, 0,
		0, 161, 162, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163,
		178, 1, 0, 0, 0, 164, 178, 5, 59, 0, 0, 165, 166, 5, 47, 0, 0, 166, 167,
		5, 42, 0, 0, 167, 171, 1, 0, 0, 0, 168, 170, 9, 0, 0, 0, 169, 168, 1, 0,
		0, 0, 170, 173, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0,
		172, 174, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174, 175, 5, 42, 0, 0, 175,
		178, 5, 47, 0, 0, 176, 178, 5, 0, 0, 1, 177, 160, 1, 0, 0, 0, 177, 164,
		1, 0, 0, 0, 177, 165, 1, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178, 50, 1, 0,
		0, 0, 7, 0, 112, 129, 135, 162, 171, 177, 1, 6, 0, 0,
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
	CellScriptLexerWHITESPACE = 10
	CellScriptLexerL_CURLY    = 11
	CellScriptLexerR_CURLY    = 12
	CellScriptLexerL_BRACKET  = 13
	CellScriptLexerR_BRACKET  = 14
	CellScriptLexerCOMMA      = 15
	CellScriptLexerNUMBER     = 16
	CellScriptLexerIDENTIFIER = 17
	CellScriptLexerSTRING     = 18
	CellScriptLexerTYPE       = 19
	CellScriptLexerBOOL       = 20
	CellScriptLexerADD        = 21
	CellScriptLexerSUB        = 22
	CellScriptLexerMUL        = 23
	CellScriptLexerDIV        = 24
	CellScriptLexerEOS        = 25
)
