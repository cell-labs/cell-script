package compiler

import (
	"github.com/cell-labs/cell-script/compiler/compiler/types"
	"github.com/cell-labs/cell-script/compiler/compiler/value"
	"github.com/llir/llvm/ir"
)

// Representation of a Go package
type pkg struct {
	name  string
	vars  map[string]value.Value
	types map[string]types.Type
}

func NewPkg(name string) *pkg {
	return &pkg{
		name:  name,
		vars:  make(map[string]value.Value),
		types: make(map[string]types.Type),
	}
}

func (p *pkg) DefinePkgVar(name string, val value.Value) {
	p.vars[name] = val
}

func (p *pkg) setExternal(internalName string, fn *ir.Func, variadic bool) value.Value {
	fn.Sig.Variadic = variadic
	val := value.Value{
		Type: &types.Function{
			LlvmReturnType: types.Void,
			FuncType:       fn.Type(),
			IsBuiltin:     true,
		},
		Value: fn,
	}
	p.DefinePkgVar(internalName, val)
	return val
}

func (p *pkg) GetPkgVar(name string, inSamePackage bool) (value.Value, bool) {

	v, ok := p.vars[name]
	return v, ok
}

func (p *pkg) DefinePkgType(name string, ty types.Type) {
	p.types[name] = ty
}

func (p *pkg) GetPkgType(name string, inSamePackage bool) (types.Type, bool) {

	v, ok := p.types[name]
	return v, ok
}
