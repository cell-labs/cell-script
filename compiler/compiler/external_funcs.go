package compiler

import (
	"github.com/llir/llvm/ir"
	llvmTypes "github.com/llir/llvm/ir/types"

	"github.com/cell-labs/cell-script/compiler/compiler/value"
)

// OSFuncs and the "os" package contains a mapping to glibc functions.
// These are used to make bootstrapping of the language easier. The end goal is to not depend on glibc.
type OSFuncs struct {
	Printf  value.Value
	Malloc  value.Value
	Realloc value.Value
	Memcpy  value.Value
	Memset  value.Value
	Strcat  value.Value
	Strcpy  value.Value
	Strncpy value.Value
	Strndup value.Value
	Strcmp value.Value
	Exit    value.Value
}

func (c *Compiler) createExternalPackage() {
	debugPkg := NewPkg("debug")
	c.osFuncs.Printf = debugPkg.setExternal("Printf", c.module.NewFunc("printf",
		i32.LLVM(),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
	), true)
	c.packages["debug"] = debugPkg

	osPkg := NewPkg("os")
	c.osFuncs.Malloc = osPkg.setExternal("malloc", c.module.NewFunc("malloc",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("size", i64.LLVM()),
	), false)

	c.osFuncs.Realloc = osPkg.setExternal("realloc", c.module.NewFunc("realloc",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("ptr", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("size", i64.LLVM()),
	), false)

	c.osFuncs.Memcpy = osPkg.setExternal("memcpy", c.module.NewFunc("memcpy",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("dest", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("src", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("n", i64.LLVM()),
	), false)

	c.osFuncs.Memset = osPkg.setExternal("memset", c.module.NewFunc("memset",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("dest", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("c", i32.LLVM()),
		ir.NewParam("n", i32.LLVM()),
	), false)

	c.osFuncs.Strcat = osPkg.setExternal("strcat", c.module.NewFunc("strcat",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
	), false)

	c.osFuncs.Strcpy = osPkg.setExternal("strcpy", c.module.NewFunc("strcpy",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("dst", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("src", llvmTypes.NewPointer(i8.LLVM())),
	), false)

	c.osFuncs.Strncpy = osPkg.setExternal("strncpy", c.module.NewFunc("strncpy",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("dst", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("src", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("n", i64.LLVM()),
	), false)

	c.osFuncs.Strndup = osPkg.setExternal("strndup", c.module.NewFunc("strndup",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("str", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("n", i64.LLVM()),
	), false)

	c.osFuncs.Strcmp = osPkg.setExternal("strcmp", c.module.NewFunc("strcmp",
		i64.LLVM(),
		ir.NewParam("lhs", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("rhs", llvmTypes.NewPointer(i8.LLVM())),
	), false)

	c.osFuncs.Exit = osPkg.setExternal("exit", c.module.NewFunc("syscall_exit",
		llvmTypes.Void,
		ir.NewParam("exit_code", i8.LLVM()),
	), false)

	c.packages["os"] = osPkg
}
