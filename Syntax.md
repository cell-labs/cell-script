Basic syntax

## Keywords

```
break
bool
const
continue
else
for
function
if
import
package
public
return
range
var
```

## Comments

```
//
/*
```

## Type System
### Types

```
bool

uint8
uint16
uint32
uint64
uint128
uint256

table
vector
union
option

byte equal to uint8
string equals to vector<byte>
```

## Variable

`var` create mutable variables
`const` create immutable variables

## Names

### Visibility

`public` export functions.
By default, everything is private in a package.

## Turing machine

```cell
// Test simulating a Turing machine.
package main

import "debug"

// brainfuck
var p, pc int64 = 0 // p for position
var a [30000]byte

const prog = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

function scan(dir int) {
	for nest := dir; dir*nest > 0; pc += dir {
		if prog[pc+dir] == ']' {
			nest--
		} else if prog[pc+dir] == '[' {
			nest++
		}
	}
}

function main() {
	r := ""
	for pc != len(prog) {
		if prog[pc] == '>' {
			p++
        }
		} else if prog[pc] == '<' {
			p--
		} else if prog[pc] == '+' {
			a[p]++
		} else if prog[pc] == '-' {
			a[p]--
		} else if prog[pc] == '.' {
			r += string(a[p])
		} else if prog[pc] == '[' {
			if a[p] == 0 {
				scan(1)
			}
		} else if prog[pc] == ']' {
			if a[p] != 0 {
				scan(-1)
			}
        } else {
			fmt.Println(r)
			return
		}
		pc++
	}
	debug.Println(r)
}

```
