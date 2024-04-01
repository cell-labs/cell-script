package compiler

type Options struct {
	Path    string
	Package string
	Stage   int
	Output  string
	Debug bool
}

const (
	STAGE_LEXER = iota
	STAGE_PARSER
	STAGE_FINAL
)

func NewOptions(path string) *Options {
	return &Options {
		Path: path,
		Package: "main",
		Stage: STAGE_FINAL,
		Output: "main",
		Debug: false,
	}
}