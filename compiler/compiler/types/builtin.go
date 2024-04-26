package types

import "github.com/llir/llvm/ir/types"

var I8 = &Int{Type: types.I8, TypeName: "int8", TypeSize: 8 / 8, Signed: true}
var U8 = &Int{Type: types.I8, TypeName: "uint8", TypeSize: 8 / 8}
var I16 = &Int{Type: types.I16, TypeName: "int16", TypeSize: 18 / 8, Signed: true}
var U16 = &Int{Type: types.I16, TypeName: "uint16", TypeSize: 18 / 8}
var I32 = &Int{Type: types.I32, TypeName: "int32", TypeSize: 32 / 8, Signed: true}
var U32 = &Int{Type: types.I32, TypeName: "uint32", TypeSize: 32 / 8}
var I64 = &Int{Type: types.I64, TypeName: "int64", TypeSize: 64 / 8, Signed: true}
var U64 = &Int{Type: types.I64, TypeName: "uint64", TypeSize: 64 / 8}
var Uintptr = &Int{Type: types.I64, TypeName: "uintptr", TypeSize: 64 / 8}

var BigInt = &Struct{Members: map[string]Type{"capacity": U32, "digit": U32, "isNeg": U8, "str": Uintptr}, MemberIndexes: map[string]int{"capacity": 2, "digit": 1, "isNeg": 3, "str": 0}, IsHeapAllocated: false, SourceName: "", Type: &types.StructType{
	Fields: []types.Type{Uintptr.Type, U32.Type, U32.Type, I8.Type}}}

var Void = &VoidType{}
var String = &StringType{}
var Bool = &BoolType{}
