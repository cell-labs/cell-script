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

type Options struct {
	Debug    bool
	Verbose  bool
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
	c := compiler.NewCompiler(&compiler.Options{Target: options.Target})
	debug := options.Debug

	err := compilePackage(c, path, "main", options)
	if err != nil {
		return err
	}

	compiled := c.GetIR()

	if options.Verbose {
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
			"-march=rv64imc",
			"-ffunction-sections", "-fdata-sections",
			"-nostdlib",
			"-L" + root,
		}
		if debug {
			crossCompileArgs = append(crossCompileArgs, "-ldummylibc-debug")
		} else {
			crossCompileArgs = append(crossCompileArgs, "-ldummylibc")

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

func compilePackage(c *compiler.Compiler, path, name string, options *Options) error {
	stdroot := options.Root
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
					parsedFiles = append(parsedFiles, parseFile(path+"/"+file.Name(), options))
				}
			}
		}
	} else {
		// Parse a single file
		parsedFiles = append(parsedFiles, parseFile(path, options))
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
					if packagePath == "debug" {
						continue
					}
					if c.IsPackageImported(packagePath) {
						continue
					}

					searchPaths := []string{
						stdroot + "/" + packagePath,
						path + "/" + packagePath,
					}

					importSuccessful := false

					for _, sp := range searchPaths {
						fp, err := os.Stat(sp)
						if err != nil || !fp.IsDir() {
							continue
						}

						if options.Verbose {
							log.Printf("Loading %s from %s", packagePath, sp)
						}

						err = compilePackage(c, sp, packagePath, options)
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

func parseFile(path string, options *Options) parser.FileNode {
	// Read specified input file
	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// Run input code through the lexer. A list of tokens is returned.
	lexed := lexer.Lex(string(fileContents))

	// Run lexed source through the parser. A syntax tree is returned.
	parsed := parser.Parse(lexed, options.Debug)

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
