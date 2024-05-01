package bigint

import (
	"fmt"
	"testing"

	"github.com/cell-labs/cell-script/compiler/lexer"
	"github.com/cell-labs/cell-script/compiler/parser"
	"github.com/stretchr/testify/assert"
)

func TestBigIntAlloc(t *testing.T) {
	// Run input code through the lexer. A list of tokens is returned.
	lexed := lexer.Lex(`
package main

var (
	v1 bigint
	v2 uint128 = 1
	v3 uint256 = 2
)

function main() {
	var b1 bigint = 3
	var b2 uint128
	var b3 uint256
	c := b3
}

`)

	lexedSugar := lexer.Lex(`
package main

var (
	v1 bigint
	v2 bigint = "1"
	v3 bigint = "2"
)

function main() {
	var b1 bigint = "3"
	var b2 bigint
	var b3 bigint
	c := b3
}

`)
	parsed := parser.Parse(lexed, false)
	res := BigInt(parsed)
	fmt.Println(res)

	parsedSugar := parser.Parse(lexedSugar, false)
	fmt.Println(parsedSugar)
	assert.Equal(t, parsedSugar, res)
}
