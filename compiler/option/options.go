package option

type VersionScheme struct{
	Major int
	Minor int
	Patch int
}

type Options struct {
	Debug    bool
	Verbose  bool
	Optimize bool
	Path     string
	Package  string
	Output   string
	Target   string
	Root     string
	Version  VersionScheme
}
