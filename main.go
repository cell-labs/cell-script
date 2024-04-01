package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/cell-labs/cell-script/internal/compiler"
)

var (
	debug  bool
	output string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <filename>\n", os.Args[0])
		flag.PrintDefaults()
	}

	// todo: add verbose compile option
	flag.BoolVarP(&debug, "debug", "d", false, "Emit debug information during compile time")
	flag.StringVarP(&output, "output", "o", "", "Output binary filename")
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Printf("No file specified. Usage: %s path/to/file.cell", os.Args[0])
		os.Exit(1)
	}

	if output == "" {
		basename := filepath.Base(flag.Arg(0))
		output = strings.TrimSuffix(basename, filepath.Ext(basename))
	}

	options := compiler.NewOptions(flag.Arg(0))
	options.Output = output
	options.Debug = debug
	err := compiler.Run(options)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
