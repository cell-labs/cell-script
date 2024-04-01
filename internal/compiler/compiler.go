package compiler

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cell-labs/cell-script/internal/lexer"
	"github.com/cell-labs/cell-script/internal/parser"
)

type compiler struct {
	Lexer  *lexer.CellScriptLexer
	Parser *parser.CellScriptParser
}

func compilePackage(options *Options) error {
	// todo: name output and debug
	p, err := os.Stat(options.Path)
	if err != nil {
		return err
	}

	var parsedFiles []parser.ISourceFileContext

	// Parse all files in the folder
	if p.IsDir() {
		root := options.Path
		err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if strings.HasSuffix(d.Name(), ".cell") {
				return errors.New("Unknown source file format")
			}
			options.Path = filepath.Join(options.Path, d.Name())
			parsedFiles = append(parsedFiles, compileFile(options))
			return nil
		})
		if err != nil {
			log.Printf("walk directories for source files error: %s", err)
		}
	} else {
		// Parse a single file
		parsedFiles = append(parsedFiles, compileFile(options))
	}
	return nil
}

func compileFile(options *Options) parser.ISourceFileContext {
	// Read specified input file
	fileContents, err := os.ReadFile(options.Path)
	if err != nil {
		log.Printf("read file error: %s\n", err)
	}

	// generate tokens using lexer
	lexer := lexer.NewCellScriptLexer(antlr.NewInputStream(string(fileContents)))
	checkStage(options.Stage, compiler{Lexer: lexer})

	// generate AST using parser
	parser := parser.NewCellScriptParser(antlr.NewCommonTokenStream(lexer, 0))
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	checkStage(options.Stage, compiler{Parser: parser})
	return parser.SourceFile()
}

func checkStage(i int, c compiler) {
	switch i {
	case STAGE_LEXER:
		for {
			t := c.Lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
			fmt.Printf("%s (%q)\n",
				c.Lexer.SymbolicNames[t.GetTokenType()], t.GetText())
		}
	}
}

func Run(options *Options) error {
	return compilePackage(options)
}
