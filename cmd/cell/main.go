package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cell-labs/cell-script/cmd/cell/build"
	"github.com/cell-labs/cell-script/compiler/option"
	flag "github.com/spf13/pflag"
)

var (
	debug    bool
	optimize bool
	verbose  bool
	output   string
	target   string
	help     bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <filename>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVarP(&debug, "debug", "d", false, "Emit debug information during compile time")
	flag.BoolVarP(&verbose, "verbose", "v", false, "Emit verbose command during compiling")
	flag.BoolVarP(&optimize, "optimize", "O", false, "Enable clang optimization")
	flag.BoolVarP(&help, "help", "h", false, "Show help message")
	flag.StringVarP(&target, "target", "t", "native", "Compile to this target")
	flag.StringVarP(&output, "output", "o", "", "Output binary filename")
}

func parseOptions() *option.Options {
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}
	if len(flag.Args()) < 1 {
		log.Printf("No file specified. Usage: %s path/to/file.cell", os.Args[0])
		os.Exit(1)
	}

	binaryPath, _ := os.Executable()
	root := filepath.Clean(binaryPath + "/../pkg/")

	if output == "" {
		basename := filepath.Base(flag.Arg(0))
		output = strings.TrimSuffix(basename, filepath.Ext(basename))
	}
	return &option.Options{
		Debug:    debug,
		Verbose:  verbose,
		Optimize: optimize,
		Path:     flag.Arg(0),
		Package:  "",
		Output:   output,
		Target:   target,
		Root:     root,
		Version:  option.VersionScheme{Major: 0, Minor: 1, Patch: 0},
	}
}

func main() {
	options := parseOptions()
	err := build.Build(options)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
