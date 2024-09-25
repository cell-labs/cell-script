package build

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cell-labs/cell-script/compiler/compiler"
	"github.com/cell-labs/cell-script/compiler/lexer"
	"github.com/cell-labs/cell-script/compiler/option"
	"github.com/cell-labs/cell-script/compiler/parser"
	"github.com/cell-labs/cell-script/compiler/passes/const_iota"
	"github.com/cell-labs/cell-script/compiler/passes/escape"
)

func Build(options *option.Options) error {
	path := options.Path
	outputBinaryPath := options.Output + ".out"
	optimize := options.Optimize
	root := options.Root
	c := compiler.NewCompiler(&compiler.Options{Target: options.Target})
	debug := options.Debug

	err := compilePackage(c, root+"/"+"builtin", "global", options)
	if err != nil {
		return err
	}

	err = compilePackage(c, path, "main", options)
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

func compilePackage(c *compiler.Compiler, path, name string, options *option.Options) error {
	stdroot := options.Root
	f, err := os.Stat(path)
	if err != nil {
		return err
	}

	var parsedFiles []parser.FileNode

	// Parse all files in the folder
	parseCurrentPackage := func() {
		if f.IsDir() {
			files, err := os.ReadDir(path)
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
	}
	parseCurrentPackage()

	// Scan for ImportNodes
	// Use importNodes to import more packages
	for _, file := range parsedFiles {
		for _, i := range file.Instructions {
			if pragma, ok := i.(*parser.PragmaNode); ok {
				version := pragma.Version
				if !(version.Major == options.Version.Major &&
					version.Minor == options.Version.Minor &&
					version.Patch == options.Version.Patch) {
					panic("unsupported compiler version")
				}
				continue
			}

			if _, ok := i.(*parser.DeclarePackageNode); ok {
				continue
			}

			if importNode, ok := i.(*parser.ImportNode); ok {

				for _, packagePath := range importNode.PackagePaths {

					// Is built in to the compiler
					if packagePath == "os" {
						continue
					}
					packageName := filepath.Base(packagePath)
					if c.IsPackageImported(packageName) && packageName != "debug" {
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

	// Parse again after importing all types needed
	clear(parsedFiles)
	parseCurrentPackage()

	return c.Compile(parser.PackageNode{
		Files: parsedFiles,
		Name:  name,
	})
}

func parseFile(path string, options *option.Options) parser.FileNode {
	// Read specified input file
	fileContents, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// Run input code through the lexer. A list of tokens is returned.
	lexed := lexer.Lex(string(fileContents))

	// Run lexed source through the parser. A syntax tree is returned.
	parsed := parser.Parse(lexed, options)

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
