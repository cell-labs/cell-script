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
	"github.com/cell-labs/cell-script/internal/ast"
	"github.com/cell-labs/cell-script/internal/codegen"
	"github.com/cell-labs/cell-script/internal/lex"
	"github.com/cell-labs/cell-script/internal/parse"
)

func compilePackage(options *Options) error {
	// todo: name output and debug
	p, err := os.Stat(options.Path)
	if err != nil {
		return err
	}

	// Parse all files in the folder
	if p.IsDir() {
		var files []string
		root := options.Path
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if strings.HasSuffix(d.Name(), ".cell") {
				return errors.New("Unknown source file format")
			}
			f := filepath.Join(options.Path, d.Name())
			files = append(files, f)
			return nil
		})
		if err != nil {
			log.Printf("walk directories for source files error: %s", err)
		}
		return CompilePackage(files, options)
	} else {
		// Parse a single file
		f := options.Path
		return CompileFile(f, options)
	}

	return nil
}

func CompilePackage(files []string, options *Options) error {
	for _, f := range files {
		err := CompileFile(f, options)
		if err != nil {
			return err
		}
	}
	return nil
}

func CompileFile(f string, options *Options) error {
	// Read specified input file
	fileContents, err := os.ReadFile(f)
	if err != nil {
		log.Printf("read file error: %s\n", err)
	}

	// generate tokens using lexer
	lexer := lex.NewCellScriptLexer(antlr.NewInputStream(string(fileContents)))
	lexer.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	if lexer.HasError() {
		return errors.New("failed to compile due to lexer errors")
	}
	if options.Stage == STAGE_LEXED {
		DumpTokens(options, lexer)
		return nil
	}

	// generate AST using parser
	parser := parse.NewCellScriptParser(antlr.NewCommonTokenStream(lexer, 0))
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	if parser.HasError() {
		return errors.New("failed to compile due to parser errors")
	}
	if options.Stage == STAGE_PARSED {
		DumpAST(options, parser)
		return nil
	}

	// generate llvm ir
	ir := codegen.NewGenerator(parser)
	if options.Stage == STAGE_CODEGENED {
		EmitIR(options, ir)
		return nil
	}
	return nil
}

func DumpTokens(options *Options, lexer *lex.CellScriptLexer) {
	fmt.Println("dump tokens:")
	for t := lexer.NextToken(); t.GetTokenType() != antlr.TokenEOF; t = lexer.NextToken() {
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

func DumpAST(options *Options, parser *parse.CellScriptParser) {
	fmt.Println("dump ast:")
	fmt.Println(ast.Dump(parser))
}

func EmitIR(options *Options, codegen *codegen.Generator) {
	fmt.Println("emit ir:")
}

func Run(options *Options) error {
	return compilePackage(options)
}
