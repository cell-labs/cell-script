package utils

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Errorln(prefix string, v any) {
	fmt.Fprint(os.Stderr, prefix)
	fmt.Fprintln(os.Stderr, v)
	fmt.Fprintln(os.Stderr)
}

func Ice(message any) {
	Errorln("error: ", "internal compiler error")
	Errorln("error: ", message)
	PrintStack()
	os.Exit(1)
}

func PrintStack() {
	Errorln("stack backtrace:\n\n", string(debug.Stack()))
}
