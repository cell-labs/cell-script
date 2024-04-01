package compiler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

type Options struct {
	Path    string
	Package string
	Stage   int
	Output  string
	Debug   bool
}

const (
	STAGE_LEXER = iota
	STAGE_PARSER
	STAGE_FINAL
	STAGE_EXIT
)
var (
	debug  bool
	output string
)

func NewOptions(path string) *Options {
	return &Options{
		Path:    path,
		Package: "main",
		Stage:   STAGE_FINAL,
		Output:  "main",
		Debug:   false,
	}
}

func SetupOptions() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <filename>\n", os.Args[0])
		flag.PrintDefaults()
	}

	// todo: add verbose compile option
	flag.BoolVarP(&debug, "debug", "d", false, "Emit debug information during compile time")
	flag.StringVarP(&output, "output", "o", "", "Output binary filename")
}

func ParseOptions() *Options {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Printf("No file specified. Usage: %s path/to/file.cell", os.Args[0])
		os.Exit(1)
	}

	if output == "" {
		basename := filepath.Base(flag.Arg(0))
		output = strings.TrimSuffix(basename, filepath.Ext(basename))
	}

	options := NewOptions(flag.Arg(0))
	options.Output = output
	options.Debug = debug
	return options
}
