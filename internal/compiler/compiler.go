package compiler

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	_ "path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cell-labs/cell-script/internal/lexer"
	"github.com/cell-labs/cell-script/internal/parser"
)

func ReadSource() {}

func compilePackage(path, name, output string, debug bool) error {
	// todo: name output and debug
	p, err := os.Stat(path)
	if err != nil {
		return err
	}

	var parsedFiles []parser.ISourceFileContext

	// Parse all files in the folder
	if p.IsDir() {
		err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if strings.HasSuffix(d.Name(), ".cell") {
				parsedFiles = append(parsedFiles, parseFile(filepath.Join(path, d.Name())))
			}
			return nil
		})
		if err != nil {
			log.Printf("walk directories for source files error: %s", err)
		}
	} else {
		// Parse a single file
		parsedFiles = append(parsedFiles, parseFile(path))
	}
	return nil
}

func parseFile(name string) parser.ISourceFileContext {
	// Read specified input file
	fileContents, err := os.ReadFile(name)
	if err != nil {
		log.Printf("read file error: %s\n", err)
	}

	// generate tokens using lexer
	lexer := lexer.NewCellScriptLexer(antlr.NewInputStream(string(fileContents)))

	// generate AST using parser
	parser := parser.NewCellScriptParser(antlr.NewCommonTokenStream(lexer, 0))
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	return parser.SourceFile()
}

func Run(path, output string, debug bool) error {
	return compilePackage(path, "main", output, debug)
}
