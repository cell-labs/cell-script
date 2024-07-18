package compiler

import (
	"fmt"

	"github.com/llir/llvm/ir"

	"github.com/cell-labs/cell-script/compiler/compiler/internal/pointer"
	"github.com/cell-labs/cell-script/compiler/compiler/name"

	"github.com/cell-labs/cell-script/compiler/compiler/internal"
	"github.com/cell-labs/cell-script/compiler/compiler/value"

	"github.com/cell-labs/cell-script/compiler/compiler/types"
	"github.com/cell-labs/cell-script/compiler/parser"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	llvmTypes "github.com/llir/llvm/ir/types"
	llvmValue "github.com/llir/llvm/ir/value"
)

// Is used in interfaces to keep track of the backing data type
var typeIDs = map[string]int64{}
var nextTypeID int64

func getTypeID(typeName string) int64 {
	if id, ok := typeIDs[typeName]; ok {
		return id
	}

	nextTypeID++
	typeIDs[typeName] = nextTypeID
	return nextTypeID
}

func (c *Compiler) parserTypeToType(typeNode parser.TypeNode) types.Type {
	switch t := typeNode.(type) {
	case *parser.SingleTypeNode:
		if len(t.PackageName) > 0 {
			tp, ok := c.packages[t.PackageName].GetPkgType(t.TypeName, false)
			if !ok {
				panic("unknown type: " + t.PackageName + "." + t.TypeName)
			}
			return tp
		}

		if res, ok := c.currentPackage.GetPkgType(t.TypeName, true); ok {
			return res
		}

		// TODO: Find a better way to organize builtin types
		if res, ok := c.packages["global"].GetPkgType(t.TypeName, true); ok {
			return res
		}

		panic("unknown type: " + t.TypeName)

	case *parser.ArrayTypeNode:
		itemType := c.parserTypeToType(t.ItemType)
		return &types.Array{
			Type:     itemType,
			LlvmType: llvmTypes.NewArray(uint64(t.Len), itemType.LLVM()),
		}

	case *parser.StructTypeNode:
		var structTypes []llvmTypes.Type
		members := make(map[string]types.Type)
		memberIndexes := t.Names

		inverseNamesIndex := make(map[int]string)
		for name, index := range memberIndexes {
			inverseNamesIndex[index] = name
		}

		for i, tt := range t.Types {
			ty := c.parserTypeToType(tt)
			members[inverseNamesIndex[i]] = ty
			structTypes = append(structTypes, ty.LLVM())
		}

		return &types.Struct{
			SourceName:    t.GetName(),
			Members:       members,
			MemberIndexes: memberIndexes,
			Type:          llvmTypes.NewStruct(structTypes...),
		}

	case *parser.SliceTypeNode:
		itemType := c.parserTypeToType(t.ItemType)
		return &types.Slice{
			Type:     itemType,
			LlvmType: internal.Slice(itemType.LLVM()),
		}

	case *parser.InterfaceTypeNode:
		requiredMethods := make(map[string]types.InterfaceMethod)

		for name, def := range t.Methods {
			ifaceMethod := types.InterfaceMethod{
				ArgumentTypes: make([]types.Type, 0),
				ReturnTypes:   make([]types.Type, 0),
			}

			for _, arg := range def.ArgumentTypes {
				ifaceMethod.ArgumentTypes = append(ifaceMethod.ArgumentTypes, c.parserTypeToType(arg))
			}
			for _, ret := range def.ReturnTypes {
				ifaceMethod.ReturnTypes = append(ifaceMethod.ReturnTypes, c.parserTypeToType(ret))
			}

			requiredMethods[name] = ifaceMethod
		}

		return &types.Interface{RequiredMethods: requiredMethods}

	case *parser.PointerTypeNode:
		return &types.Pointer{
			Type: c.parserTypeToType(t.ValueType),
		}

	case *parser.FuncTypeNode:
		retType, treReturnTypes, llvmArgTypes, treParams, _, _ := c.funcType(t.ArgTypes, t.RetTypes)

		fn := ir.NewFunc("UNNAMEDFUNC", retType.LLVM(), llvmArgTypes...)

		return &types.Function{
			ArgumentTypes:  treParams,
			ReturnTypes:    treReturnTypes,
			FuncType:       fn.Type(),
			LlvmReturnType: retType,
		}
	}

	panic(fmt.Sprintf("unknown typeNode: %T", typeNode))
}

func (c *Compiler) compileTypeCastNode(v *parser.TypeCastNode) value.Value {
	val := c.compileValue(v.Val)

	var current *llvmTypes.IntType
	var ok bool

	current, ok = val.Type.LLVM().(*llvmTypes.IntType)
	if !ok {
		panic("TypeCast origin must be int type")
	}

	targetType := c.parserTypeToType(v.Type)
	target, ok := targetType.LLVM().(*llvmTypes.IntType)
	if !ok {
		panic("TypeCast target must be int type")
	}

	llvmVal := internal.LoadIfVariable(c.contextBlock, val)

	// Same size, nothing to do here
	if current.BitSize == target.BitSize {
		return val
	}

	res := c.contextBlock.NewAlloca(target)

	var changedSize llvmValue.Value

	if current.BitSize < target.BitSize {
		changedSize = c.contextBlock.NewSExt(llvmVal, target)
	} else {
		changedSize = c.contextBlock.NewTrunc(llvmVal, target)
	}

	c.contextBlock.NewStore(changedSize, res)

	return value.Value{
		Value:      res,
		Type:       targetType,
		IsVariable: true,
	}
}

func (c *Compiler) compileTypeCastInterfaceNode(v *parser.TypeCastInterfaceNode) value.Value {
	tryCastToType := c.parserTypeToType(v.Type)

	// Allocate the OK variable
	okVal := c.contextBlock.NewAlloca(types.Bool.LLVM())
	types.Bool.Zero(c.contextBlock, okVal)
	okVal.SetName(name.Var("ok"))

	resCastedVal := c.contextBlock.NewAlloca(tryCastToType.LLVM())
	tryCastToType.Zero(c.contextBlock, resCastedVal)
	resCastedVal.SetName(name.Var("rescastedval"))

	interfaceVal := c.compileValue(v.Item)

	interfaceDataType := c.contextBlock.NewGetElementPtr(pointer.ElemType(interfaceVal.Value), interfaceVal.Value, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, 1))
	loadedInterfaceDataType := c.contextBlock.NewLoad(pointer.ElemType(interfaceDataType), interfaceDataType)

	trueBlock := c.contextBlock.Parent.NewBlock(name.Block() + "-was-correct-type")
	falseBlock := c.contextBlock.Parent.NewBlock(name.Block() + "-was-other-type")
	afterBlock := c.contextBlock.Parent.NewBlock(name.Block() + "-after-type-check")

	trueBlock.NewBr(afterBlock)
	falseBlock.NewBr(afterBlock)

	backingTypeID := getTypeID(tryCastToType.Name())
	cmp := c.contextBlock.NewICmp(enum.IPredEQ, loadedInterfaceDataType, constant.NewInt(llvmTypes.I32, backingTypeID))
	c.contextBlock.NewCondBr(cmp, trueBlock, falseBlock)

	trueBlock.NewStore(constant.NewInt(llvmTypes.I1, 1), okVal)

	backingDataPtr := trueBlock.NewGetElementPtr(pointer.ElemType(interfaceVal.Value), interfaceVal.Value, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, 0))
	loadedBackingDataPtr := trueBlock.NewLoad(pointer.ElemType(backingDataPtr), backingDataPtr)
	casted := trueBlock.NewBitCast(loadedBackingDataPtr, llvmTypes.NewPointer(tryCastToType.LLVM()))
	loadedCasted := trueBlock.NewLoad(pointer.ElemType(casted), casted)
	trueBlock.NewStore(loadedCasted, resCastedVal)

	c.contextBlock = afterBlock

	return value.Value{
		Type: &types.MultiValue{
			Types: []types.Type{
				tryCastToType,
				types.Bool,
			},
		},
		MultiValues: []value.Value{
			{Type: tryCastToType, Value: resCastedVal, IsVariable: true},
			{Type: types.Bool, Value: okVal, IsVariable: true},
		},
	}
}
