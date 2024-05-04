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
	v2 bigint = bigint{str:"1",digits:1,capacity:0,isNeg:0}
	v3 bigint = bigint{str:"2",digits:1,capacity:0,isNeg:0}
)

function main() {
	var b1 bigint = bigint{str:"3",digits:1,capacity:0,isNeg:0}
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

func TestBigIntCondition(t *testing.T) {
	// Run input code through the lexer. A list of tokens is returned.
	lexed := lexer.Lex(`
function main() {
	v := 1u128
	if v >= 1u128 {
	}
}

`)
	lexedSugar := lexer.Lex(`
function main() {
	v := bigint{str:"1",digits:1,capacity:0,isNeg:0}
	if bigIntGTE(v, bigint{str:"1",digits:1,capacity:0,isNeg:0}) {
	}
}
`)
	parsed := parser.Parse(lexed, false)
	res := BigInt(parsed)

	parsedSugar := parser.Parse(lexedSugar, false)
	assert.Equal(t, parsedSugar, res)
}



// if v >= 1u128 {
// }
// if v == 1u128 {
// }
// if v < 2u128 {
// }
// if v <= 1u128 {
// }

// if bigIntGT(v, 0u128) {
// }
// if bigIntGTE(v, 1u128) {
// }
// if bigIntEqual(v, 1u128) {
// }
// if bigIntLT(v, 2u128) {
// }
// if bigIntLTE(v, 1u128) {
// }