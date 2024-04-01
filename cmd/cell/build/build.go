package build

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/cell-labs/cell-script/compiler/compiler"
	"github.com/cell-labs/cell-script/compiler/lexer"
	"github.com/cell-labs/cell-script/compiler/parser"
	"github.com/cell-labs/cell-script/compiler/passes/const_iota"
	"github.com/cell-labs/cell-script/compiler/passes/escape"
)

var debug bool

type Options struct {
	Debug    bool
	Optimize bool
	Path     string
	Package  string
	Output   string
	Target   string
	Root     string
}

func Build(options *Options) error {
	path := options.Path
	outputBinaryPath := options.Output
	optimize := options.Optimize
	root := options.Root
	c := compiler.NewCompiler()
	debug = options.Debug

	err := compilePackage(c, path, root, "main")
	if err != nil {
		return err
	}

	compiled := c.GetIR()

	if debug {
		fmt.Println(compiled)
	}

	// Get dir to save temporary dirs in
	tmpDir, err := os.MkdirTemp("", ".tpircsllec")
	if err != nil {
		panic(err)
	}

	// Write LLVM IR to disk
	err = os.WriteFile(tmpDir+"/main.ll", []byte(compiled), 0666)
	if err != nil {
		panic(err)
	}

	if outputBinaryPath == "" {
		outputBinaryPath = "output-binary"
	}

	clangArgs := []string{
		"-Wno-override-module", // Disable override target triple warnings
		tmpDir + "/main.ll",    // Path to LLVM IR
		"-o", outputBinaryPath, // Output path
	}

	if optimize {
		clangArgs = append(clangArgs, "-O3")
	}

	if options.Target == "riscv" {
		crossCompileArgs := []string{
			"--target=riscv64",
			"-march=rv64imac",
			"--sysroot=/opt/homebrew/opt/riscv-gnu-toolchain/riscv64-unknown-elf",
			"--gcc-toolchain=/opt/homebrew/opt/riscv-gnu-toolchain",
		}
		clangArgs = append(clangArgs, crossCompileArgs...)
	}

	// Invoke clang compiler to compile LLVM IR to a binary executable
	cmd := exec.Command("clang", clangArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return err
	}

	if len(output) > 0 {
		fmt.Println(string(output))
		return errors.New("Clang failure")
	}

	return nil
}

func compilePackage(c *compiler.Compiler, path, goroot, name string) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	}

	var parsedFiles []parser.FileNode

	// Parse all files in the folder
	if f.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(path + ": " + err.Error())
		}

		for _, file := range files {
			if !file.IsDir() {
				// Tre files doesn't have to contain valid Go code, and is used to prevent issues
				// with some of the go tools (like vgo)
				if strings.HasSuffix(file.Name(), ".cell") || strings.HasSuffix(file.Name(), ".go") {
					parsedFiles = append(parsedFiles, parseFile(path+"/"+file.Name()))
				}
			}
		}
	} else {
		// Parse a single file
		parsedFiles = append(parsedFiles, parseFile(path))
	}

	// Scan for ImportNodes
	// Use importNodes to import more packages
	for _, file := range parsedFiles {
		for _, i := range file.Instructions {
			if _, ok := i.(*parser.DeclarePackageNode); ok {
				continue
			}

			if importNode, ok := i.(*parser.ImportNode); ok {

				for _, packagePath := range importNode.PackagePaths {

					// Is built in to the compiler
					if packagePath == "debug" || packagePath == "tx" || packagePath == "cell" {
						continue
					}

					searchPaths := []string{
						path + "/vendor/" + packagePath,
						goroot + "/" + packagePath,
					}

					importSuccessful := false

					for _, sp := range searchPaths {
						fp, err := os.Stat(sp)
						if err != nil || !fp.IsDir() {
							continue
						}

						if debug {
							log.Printf("Loading %s from %s", packagePath, sp)
						}

						err = compilePackage(c, sp, goroot, packagePath)
						if err != nil {
							return err
						}

						importSuccessful = true
					}

					if !importSuccessful {
						return fmt.Errorf("Unable to import: %s", packagePath)
					}
				}

				continue
			}

			break
		}
	}

	return c.Compile(parser.PackageNode{
		Files: parsedFiles,
		Name:  name,
	})
}

func parseFile(path string) parser.FileNode {
	// Read specified input file
	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// Run input code through the lexer. A list of tokens is returned.
	lexed := lexer.Lex(string(fileContents))

	// Run lexed source through the parser. A syntax tree is returned.
	parsed := parser.Parse(lexed, debug)

	// List of passes to run on the AST
	passes := []func(*parser.FileNode) *parser.FileNode{
		const_iota.Iota,
		escape.Escape,
	}
	for _, pass := range passes {
		parsed = pass(parsed)
	}

	return *parsed
}
