// Test simulating a Turing machine, sort of.
package main

import "fmt"

// brainfuck

var p, pc int
var a [30000]byte

const prog = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

func scan(dir int) {
	for nest := dir; dir*nest > 0; pc += dir {
		switch prog[pc+dir] {
		case ']':
			nest--
		case '[':
			nest++
		}
	}
}

func main() {
	r := ""
	for pc != len(prog) {
		switch prog[pc] {
		case '>':
			p++
		case '<':
			p--
		case '+':
			a[p]++
		case '-':
			a[p]--
		case '.':
			r += string(a[p])
		case '[':
			if a[p] == 0 {
				scan(1)
			}
		case ']':
			if a[p] != 0 {
				scan(-1)
			}
		default:
			fmt.Println(r)
			return
		}
		pc++
	}
	fmt.Println(r)
}
