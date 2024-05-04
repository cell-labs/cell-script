package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerSimpleAdd(t *testing.T) {
	r := Lex("aa + b")

	expected := []Item{
		{Type: IDENTIFIER, Val: "aa", Line: 1},
		{Type: OPERATOR, Val: "+", Line: 1},
		{Type: IDENTIFIER, Val: "b", Line: 1},
		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)
}

func TestLexerSimpleAddWithNewlines(t *testing.T) {
	r := Lex("aa + b\naa + b")

	expected := []Item{
		{Type: IDENTIFIER, Val: "aa", Line: 1},
		{Type: OPERATOR, Val: "+", Line: 1},
		{Type: IDENTIFIER, Val: "b", Line: 1},

		{Type: EOL, Val: "", Line: 1},

		{Type: IDENTIFIER, Val: "aa", Line: 2},
		{Type: OPERATOR, Val: "+", Line: 2},
		{Type: IDENTIFIER, Val: "b", Line: 2},

		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)
}

func TestLexerSimpleAddNumber(t *testing.T) {
	r := Lex("aa + 14")

	expected := []Item{
		{Type: IDENTIFIER, Val: "aa", Line: 1},
		{Type: OPERATOR, Val: "+", Line: 1},
		{Type: NUMBER, Val: "14", Line: 1},

		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)
}

func TestLexerSimpleCall(t *testing.T) {
	r := Lex("foo(bar)")

	expected := []Item{
		{Type: IDENTIFIER, Val: "foo", Line: 1},
		{Type: OPERATOR, Val: "(", Line: 1},
		{Type: IDENTIFIER, Val: "bar", Line: 1},
		{Type: OPERATOR, Val: ")", Line: 1},

		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)
}

func TestLexerSimpleCallWithString(t *testing.T) {
	r := Lex("foo(\"bar\")")

	expected := []Item{
		{Type: IDENTIFIER, Val: "foo", Line: 1},
		{Type: OPERATOR, Val: "(", Line: 1},
		{Type: STRING, Val: "bar", Line: 1},
		{Type: OPERATOR, Val: ")", Line: 1},

		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)
}

func TestString(t *testing.T) {
	r := Lex(`"bar"`)

	expected := []Item{
		{Type: STRING, Val: "bar", Line: 1},
		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)
}

func TestEscapedString(t *testing.T) {
	r := Lex(`"bar\""`)

	expected := []Item{
		{Type: STRING, Val: "bar\"", Line: 1},
		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)
}

func TestChar(t *testing.T) {
	r := Lex(`'a'`)

	expected := []Item{
		{Type: BYTE, Val: "a", Line: 1},
		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)
}

func TestEscapedChar(t *testing.T) {
	r := Lex(`'\''`)

	expected := []Item{
		{Type: BYTE, Val: "'", Line: 1},
		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)
}

func TestLexerSimpleCallWithTwoStrings(t *testing.T) {
	r := Lex(`foo("bar", "baz")`)

	expected := []Item{
		{Type: IDENTIFIER, Val: "foo", Line: 1},
		{Type: OPERATOR, Val: "(", Line: 1},
		{Type: STRING, Val: "bar", Line: 1},
		{Type: OPERATOR, Val: ",", Line: 1},
		{Type: STRING, Val: "baz", Line: 1},
		{Type: OPERATOR, Val: ")", Line: 1},
		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)
}

func TestLexerSimpleCallWithStringNum(t *testing.T) {
	r := Lex(`printf("%d\n", 123)`)

	expected := []Item{
		{Type: IDENTIFIER, Val: "printf", Line: 1},
		{Type: OPERATOR, Val: "(", Line: 1},
		{Type: STRING, Val: "%d\n", Line: 1},
		{Type: OPERATOR, Val: ",", Line: 1},
		{Type: NUMBER, Val: "123", Line: 1},
		{Type: OPERATOR, Val: ")", Line: 1},
		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)
}

func TestLexerConstantIntegerWithSuffix(t *testing.T) {
	r := Lex(`printf("%d\n", 123u128)`)
	expected := []Item{
		{Type: IDENTIFIER, Val: "printf", Line: 1},
		{Type: OPERATOR, Val: "(", Line: 1},
		{Type: STRING, Val: "%d\n", Line: 1},
		{Type: OPERATOR, Val: ",", Line: 1},
		{Type: BIGNUMBER, Val: "123", Line: 1, Suffix: "u128"},
		{Type: OPERATOR, Val: ")", Line: 1},
		{Type: EOL},
		{Type: EOF},
	}
	assert.Equal(t, expected, r)

	r = Lex(`a := 100u128`)
	expected = []Item{
		{Type: IDENTIFIER, Val: "a", Line: 1},
		{Type: OPERATOR, Val: ":=", Line: 1},
		{Type: BIGNUMBER, Val: "100", Line: 1, Suffix: "u128"},
		{Type: EOL},
		{Type: EOF},
	}

	assert.Equal(t, expected, r)

	r = Lex(`a = 100u256`)
	expected = []Item{
		{Type: IDENTIFIER, Val: "a", Line: 1},
		{Type: OPERATOR, Val: "=", Line: 1},
		{Type: BIGNUMBER, Val: "100", Line: 1, Suffix: "u256"},
		{Type: EOL},
		{Type: EOF},
	}
	assert.Equal(t, expected, r)

	r = Lex(`a = 100u64`)
	expected = []Item{
		{Type: IDENTIFIER, Val: "a", Line: 1},
		{Type: OPERATOR, Val: "=", Line: 1},
		{Type: NUMBER, Val: "100", Line: 1, Suffix: "u64"},
		{Type: EOL},
		{Type: EOF},
	}
	assert.Equal(t, expected, r)
}
