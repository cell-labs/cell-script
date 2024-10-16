package types

import (
	"fmt"
	"math/big"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	llvmValue "github.com/llir/llvm/ir/value"

	"github.com/cell-labs/cell-script/compiler/compiler/internal/pointer"
	"github.com/cell-labs/cell-script/compiler/compiler/name"

	"github.com/cell-labs/cell-script/compiler/compiler/strings"
)

type Type interface {
	LLVM() types.Type
	Name() string

	// Size of type in bytes
	Size() int64

	AddMethod(string, *Method)
	GetMethod(string) (*Method, bool)

	Zero(*ir.Block, llvmValue.Value)

	IsSigned() bool
}

type backingType struct {
	methods map[string]*Method
}

func (b *backingType) AddMethod(name string, method *Method) {
	if b.methods == nil {
		b.methods = make(map[string]*Method)
	}
	b.methods[name] = method
}

func (b *backingType) GetMethod(name string) (*Method, bool) {
	m, ok := b.methods[name]
	return m, ok
}

func (backingType) Size() int64 {
	panic("Type does not have size set")
}

func (backingType) Zero(*ir.Block, llvmValue.Value) {
	// NOOP
}

func (backingType) IsSigned() bool {
	return false
}

type Struct struct {
	backingType

	Members       map[string]Type
	MemberIndexes map[string]int

	IsHeapAllocated bool

	SourceName string
	Type       types.Type
}

func (s Struct) LLVM() types.Type {
	return s.Type
}

func (s Struct) Name() string {
	return fmt.Sprintf("Any")
}

func (s Struct) Zero(block *ir.Block, alloca llvmValue.Value) {
	for key, valType := range s.Members {
		ptr := block.NewGetElementPtr(pointer.ElemType(alloca), alloca,
			constant.NewInt(types.I32, 0),
			constant.NewInt(types.I32, int64(s.MemberIndexes[key])),
		)
		valType.Zero(block, ptr)
	}
}

func (s Struct) Size() int64 {
	var sum int64
	for _, valType := range s.Members {
		sum += valType.Size()
	}
	return sum
}

type Method struct {
	backingType

	Function        *Function
	LlvmFunction    llvmValue.Named
	PointerReceiver bool
	MethodName      string
}

func (m Method) LLVM() types.Type {
	return m.Function.LLVM()
}

func (m Method) Name() string {
	return m.MethodName
}

type Function struct {
	backingType

	// LlvmFunction llvmValue.Named
	FuncType types.Type

	// The return type of the LLVM function (is always 1)
	LlvmReturnType Type
	// Return types of the Tre function
	ReturnTypes []Type

	IsVariadic    bool
	IsExtern 	  bool
	ArgumentTypes []Type
	IsBuiltin    bool

	// Is used when calling an interface method
	JumpFunction *ir.Func
}

func (f Function) LLVM() types.Type {
	return f.FuncType
}

func (f Function) Name() string {
	return "func"
}

type BoolType struct {
	backingType
}

func (BoolType) LLVM() types.Type {
	return types.I1
}

func (BoolType) Name() string {
	return "bool"
}

func (BoolType) Size() int64 {
	return 1
}

func (b BoolType) Zero(block *ir.Block, alloca llvmValue.Value) {
	block.NewStore(constant.NewInt(types.I1, 0), alloca)
}

type VoidType struct {
	backingType
}

func (VoidType) LLVM() types.Type {
	return types.Void
}

func (VoidType) Name() string {
	return "void"
}

func (VoidType) Size() int64 {
	return 0
}

type Int struct {
	backingType

	Type     *types.IntType
	TypeName string
	TypeSize int64
	Signed   bool
}

func (i Int) LLVM() types.Type {
	return i.Type
}

func (i Int) Name() string {
	return i.TypeName
}

func (i Int) Size() int64 {
	return i.TypeSize
}

func (i Int) Zero(block *ir.Block, alloca llvmValue.Value) {
	b := big.NewInt(0)
	if !i.IsSigned() {
		b.SetUint64(0)
	}

	c := &constant.Int{
		Typ: i.Type,
		X:   b,
	}

	block.NewStore(c, alloca)
}

func (i Int) IsSigned() bool {
	return i.Signed
}

type StringType struct {
	backingType
	Type types.Type
}

// Populated by compiler.go
var ModuleStringType types.Type
var EmptyStringConstant *ir.Global

func (StringType) LLVM() types.Type {
	return ModuleStringType
}

func (StringType) Name() string {
	return "string"
}

func (StringType) Size() int64 {
	return 16
}

func (s StringType) Zero(block *ir.Block, alloca llvmValue.Value) {
	lenPtr := block.NewGetElementPtr(pointer.ElemType(alloca), alloca, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	backingDataPtr := block.NewGetElementPtr(pointer.ElemType(alloca), alloca, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	block.NewStore(constant.NewInt(types.I64, 0), lenPtr)
	block.NewStore(strings.Toi8Ptr(block, EmptyStringConstant), backingDataPtr)
}

type Array struct {
	backingType
	Type     Type
	Len      uint64
	LlvmType types.Type
}

func (a Array) LLVM() types.Type {
	return a.LlvmType
}

func (a Array) Name() string {
	return "array"
}

func (a Array) Size() int64 {
	return int64(a.Len) * a.Type.Size()
}

func (a Array) Zero(block *ir.Block, alloca llvmValue.Value) {
	for i := uint64(0); i < a.Len; i++ {
		ptr := block.NewGetElementPtr(pointer.ElemType(alloca), alloca, constant.NewInt(types.I64, 0), constant.NewInt(types.I64, int64(i)))
		a.Type.Zero(block, ptr)
	}
}

type Slice struct {
	backingType
	Type     Type // type of the items in the slice []int => int
	LlvmType types.Type
}

func (s Slice) LLVM() types.Type {
	return s.LlvmType
}

func (s Slice) Name() string {
	return "slice" + s.Type.Name()
}

func (Slice) Size() int64 {
	return 3*4 + 8 // 3 int32s and a pointer
}

func (s Slice) SliceZero(block *ir.Block, mallocFunc llvmValue.Named, memsetFunc llvmValue.Named, initLen, initCap llvmValue.Value, emptySlice llvmValue.Value) {
	// The cap must always be larger than 0
	// Use 2 as the default value
	// Todo: check initCap
	// if initCap < 2 {
	// 	initCap = 2
	// }

	len := block.NewGetElementPtr(pointer.ElemType(emptySlice), emptySlice, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	len.SetName(name.Var("len"))
	cap := block.NewGetElementPtr(pointer.ElemType(emptySlice), emptySlice, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	cap.SetName(name.Var("cap"))
	offset := block.NewGetElementPtr(pointer.ElemType(emptySlice), emptySlice, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 2))
	offset.SetName(name.Var("offset"))
	backingArray := block.NewGetElementPtr(pointer.ElemType(emptySlice), emptySlice, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 3))
	backingArray.SetName(name.Var("backing"))

	initLen32 := initLen
	if initLen32.Type() != types.I32 {
		initLen32 = block.NewTrunc(initLen32, types.I32)
	}
	initCap32 := initCap
	if initCap32.Type() != types.I32 {
		initCap32 = block.NewTrunc(initCap32, types.I32)
	}
	block.NewStore(initLen32, len)
	block.NewStore(initCap32, cap)
	block.NewStore(constant.NewInt(types.I32, 0), offset)

	size := block.NewMul(initCap32, constant.NewInt(types.I32, s.Type.Size()))
	size64 := llvmValue.Value(size)
	if size64.Type() != types.I64 {
		size64 = block.NewSExt(size64, types.I64)
	}
	mallocatedSpaceRaw := block.NewCall(mallocFunc, size64)
	// todo: memset cause data error
	// block.NewCall(memsetFunc, mallocatedSpaceRaw, constant.NewInt(types.I32, 0), size64)
	mallocatedSpaceRaw.SetName(name.Var("slicezero"))
	bitcasted := block.NewBitCast(mallocatedSpaceRaw, types.NewPointer(s.Type.LLVM()))
	block.NewStore(bitcasted, backingArray)
}

type Pointer struct {
	backingType

	Type                  Type
	IsNonAllocDereference bool

	LlvmType types.Type
}

func (p Pointer) LLVM() types.Type {
	return types.NewPointer(p.Type.LLVM())
}

func (p Pointer) Name() string {
	return fmt.Sprintf("pointer(%s)", p.Type.Name())
}

func (p Pointer) Size() int64 {
	return 8
}

func (p Pointer) Zero(block *ir.Block, alloca llvmValue.Value) {
	i8PtrPtr := block.NewBitCast(alloca, types.NewPointer(types.NewPointer(types.I8)))
	block.NewStore(constant.NewIntToPtr(constant.NewInt(types.I64, 0), types.NewPointer(types.I8)), i8PtrPtr)
}

// MultiValue is used when returning multiple values from a function
type MultiValue struct {
	backingType
	Types []Type
}

func (m MultiValue) Name() string {
	return "multivalue"
}

func (m MultiValue) LLVM() types.Type {
	panic("MutliValue has no LLVM type")
}

type UntypedConstantNumber struct {
	backingType
}

func (m UntypedConstantNumber) Name() string {
	return "UntypedConstantNumber"
}

func (m UntypedConstantNumber) LLVM() types.Type {
	panic("UntypedConstantNumber has no LLVM type")
}
