package lexer

import (
	"fmt"
	"testing"

	"github.com/cell-labs/cell-script/internal/lex"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
)

func tokensToString(tokens []antlr.Token) string {
	return fmt.Sprint(tokens)
}

func TestLexer(t *testing.T) {
	lexer := lex.NewCellScriptLexer(antlr.NewInputStream("a+b"))
	assert.Equal(t, "[[@-1,0:0='a',<19>,1:0] [@-1,1:1='+',<23>,1:1] [@-1,2:2='b',<19>,1:2]]", tokensToString(lexer.GetAllTokens()))
}
