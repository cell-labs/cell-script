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
	"github.com/cell-labs/cell-script/internal/lex"
	"github.com/cell-labs/cell-script/internal/parse"
)

type compiler struct {
	Lexer  *lex.CellScriptLexer
	Parser *parse.CellScriptParser
}

func compilePackage(options *Options) error {
	// todo: name output and debug
	p, err := os.Stat(options.Path)
	if err != nil {
		return err
	}

	var parsedFiles []parse.ISourceFileContext

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

func compileFile(options *Options) parse.ISourceFileContext {
	// Read specified input file
	fileContents, err := os.ReadFile(options.Path)
	if err != nil {
		log.Printf("read file error: %s\n", err)
	}

	// generate tokens using lexer
	lexer := lex.NewCellScriptLexer(antlr.NewInputStream(string(fileContents)))
	DumpTokens(options, lexer)

	// generate AST using parser
	parser := parse.NewCellScriptParser(antlr.NewCommonTokenStream(lexer, 0))
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	DumpAST(options, parser)
	return parser.SourceFile()
}

func DumpTokens(options *Options, lexer *lex.CellScriptLexer) {
	if options.Stage != STAGE_LEXED {
		return
	}
	for t := lexer.NextToken(); t.GetTokenType() == antlr.TokenEOF; t = lexer.NextToken() {
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

func DumpAST(options *Options, parser *parse.CellScriptParser) {
	if options.Stage != STAGE_PARSED {
		return
	}
	fmt.Println(ast.Dump(parser))
}

func Run(options *Options) error {
	return compilePackage(options)
}
