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
	Strcat  value.Value
	Strcpy  value.Value
	Strncpy value.Value
	Strndup value.Value
	Exit    value.Value
}

type BigIntFuncs struct {
	New         value.Value
	Clone       value.Value
	FromString  value.Value
	Free        value.Value
	Assign      value.Value
	AssignInt64 value.Value
	Print       value.Value
	ToString    value.Value
	EQUAL       value.Value
	GT          value.Value
	GTE         value.Value
	LT          value.Value
	LTE         value.Value
	ADD         value.Value
	SUB         value.Value
	MUL         value.Value
	DIV         value.Value
	MOD         value.Value
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
		ir.NewParam("", i64.LLVM()),
	), false)

	c.osFuncs.Realloc = osPkg.setExternal("realloc", c.module.NewFunc("realloc",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", i64.LLVM()),
	), false)

	c.osFuncs.Memcpy = osPkg.setExternal("memcpy", c.module.NewFunc("memcpy",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("dest", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("src", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("n", i64.LLVM()),
	), false)

	c.osFuncs.Strcat = osPkg.setExternal("strcat", c.module.NewFunc("strcat",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
	), false)

	c.osFuncs.Strcpy = osPkg.setExternal("strcpy", c.module.NewFunc("strcpy",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
	), false)

	c.osFuncs.Strncpy = osPkg.setExternal("strncpy", c.module.NewFunc("strncpy",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", i64.LLVM()),
	), false)

	c.osFuncs.Strndup = osPkg.setExternal("strndup", c.module.NewFunc("strndup",
		llvmTypes.NewPointer(i8.LLVM()),
		ir.NewParam("", llvmTypes.NewPointer(i8.LLVM())),
		ir.NewParam("", i64.LLVM()),
	), false)

	c.osFuncs.Exit = osPkg.setExternal("exit", c.module.NewFunc("syscall_exit",
		llvmTypes.Void,
		ir.NewParam("", i8.LLVM()),
	), false)

	c.packages["os"] = osPkg
}

func (c *Compiler) createBigInt() {
	globalPkg := c.packages["global"]
	newBigIntType := func() *llvmTypes.StructType {
		return llvmTypes.NewStruct(Uintptr.Type, i32.Type, i32.Type, i8.Type)
	}
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntNew", c.module.NewFunc("big_init_new",
		newBigIntType(),
	), false)
	c.bigIntFuncs.Clone = globalPkg.setExternal("bigIntClone", c.module.NewFunc("big_init_clone",
		newBigIntType(),
		ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntFree", c.module.NewFunc("big_init_free",
		llvmTypes.Void,
		ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntFromString", c.module.NewFunc("big_init_from_string",
		newBigIntType(),
		ir.NewParam("", strTy.Type),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntPrint", c.module.NewFunc("big_init_print",
		llvmTypes.Void,
		ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntLen", c.module.NewFunc("big_init_len",
		u32.Type,
		ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntToString", c.module.NewFunc("big_init_to_string",
		strTy.Type,
		ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntAssign", c.module.NewFunc("big_init_assign",
		boolean.LLVM(),
		ir.NewParam("", newBigIntType()), ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntGT", c.module.NewFunc("big_init_gt",
		boolean.LLVM(),
		ir.NewParam("", newBigIntType()), ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntGTE", c.module.NewFunc("big_init_gte",
		boolean.LLVM(),
		ir.NewParam("", newBigIntType()), ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntLT", c.module.NewFunc("big_init_lt",
		boolean.LLVM(),
		ir.NewParam("", newBigIntType()), ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntLTE", c.module.NewFunc("big_init_lte",
		boolean.LLVM(),
		ir.NewParam("", newBigIntType()), ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntEqual", c.module.NewFunc("big_init_equal",
		boolean.LLVM(),
		ir.NewParam("", newBigIntType()), ir.NewParam("", newBigIntType()),
	), false)

	c.bigIntFuncs.New = globalPkg.setExternal("bigIntAdd", c.module.NewFunc("big_init_add",
		newBigIntType(),
		ir.NewParam("", newBigIntType()), ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntSub", c.module.NewFunc("big_init_sub",
		newBigIntType(),
		ir.NewParam("", newBigIntType()), ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntMul", c.module.NewFunc("big_init_mul",
		newBigIntType(),
		ir.NewParam("", newBigIntType()), ir.NewParam("", newBigIntType()),
	), false)
	c.bigIntFuncs.New = globalPkg.setExternal("bigIntDiv", c.module.NewFunc("big_init_div",
		newBigIntType(),
		ir.NewParam("", newBigIntType()), ir.NewParam("", newBigIntType()),
	), false)
}
