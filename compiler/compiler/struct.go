package compiler

import (
	"fmt"

	"github.com/llir/llvm/ir/constant"
	llvmTypes "github.com/llir/llvm/ir/types"
	llvmValue "github.com/llir/llvm/ir/value"

	"github.com/cell-labs/cell-script/compiler/compiler/internal/pointer"
	"github.com/cell-labs/cell-script/compiler/compiler/name"
	"github.com/cell-labs/cell-script/compiler/compiler/types"
	"github.com/cell-labs/cell-script/compiler/compiler/value"
	"github.com/cell-labs/cell-script/compiler/parser"
)

func (c *Compiler) compileStructLoadElementNode(v *parser.StructLoadElementNode) value.Value {
	src := c.compileValue(v.Struct)

	// Use this type, or the type behind the pointer
	targetType := src.Type
	var isPointer bool
	var isPointerNonAllocDereference bool
	if pointerType, ok := src.Type.(*types.Pointer); ok {
		targetType = pointerType.Type
		isPointerNonAllocDereference = pointerType.IsNonAllocDereference
		isPointer = true
	}

	if !src.IsVariable && !isPointer {
		// GetElementPtr only works on pointer types, and we don't have a pointer to our object.
		// Allocate it and use the pointer instead
		dst := c.contextBlock.NewAlloca(src.Type.LLVM())
		c.contextBlock.NewStore(src.Value, dst)
		src = value.Value{
			Value:      dst,
			Type:       src.Type,
			IsVariable: true,
		}
	}

	// Check if it is a struct member
	if structType, ok := targetType.(*types.Struct); ok {
		if memberIndex, ok := structType.MemberIndexes[v.ElementName]; ok {
			val := src.Value

			if isPointer && !isPointerNonAllocDereference && src.IsVariable {
				val = c.contextBlock.NewLoad(pointer.ElemType(val), val)
			}

			retVal := c.contextBlock.NewGetElementPtr(pointer.ElemType(val), val, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, int64(memberIndex)))

			return value.Value{
				Type:       structType.Members[v.ElementName],
				Value:      retVal,
				IsVariable: true,
			}
		}
	}

	// Check if it's a method
	if method, ok := targetType.GetMethod(v.ElementName); ok {
		return value.Value{
			Type:       method,
			Value:      src.Value,
			IsVariable: false,
		}
	}

	// Check if it's a interface method
	if iface, ok := src.Type.(*types.Interface); ok {
		if ifaceMethod, ok := iface.RequiredMethods[v.ElementName]; ok {
			// Find method index
			// TODO: This can be much smarter
			var methodIndex int64
			for i, name := range iface.SortedRequiredMethods() {
				if name == v.ElementName {
					methodIndex = int64(i)
					break
				}
			}

			// Load jump function
			jumpTable := c.contextBlock.NewGetElementPtr(pointer.ElemType(src.Value), src.Value, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, 2))
			jumpLoad := c.contextBlock.NewLoad(pointer.ElemType(jumpTable), jumpTable)
			jumpFunc := c.contextBlock.NewGetElementPtr(pointer.ElemType(jumpLoad), jumpLoad, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, methodIndex))
			jumpFuncLoad := c.contextBlock.NewLoad(pointer.ElemType(jumpFunc), jumpFunc)

			// Set jump function
			ifaceMethod.LlvmJumpFunction = jumpFuncLoad

			return value.Value{
				Type:       &ifaceMethod,
				Value:      src.Value,
				IsVariable: false,
			}
		}
	}

	panic(fmt.Sprintf("%T internal error: no such type map indexing: %s", src, v.ElementName))
}

func (c *Compiler) compileInitStructWithValues(v *parser.InitializeStructNode) value.Value {
	treType := c.parserTypeToType(v.Type)

	structType, ok := treType.(*types.Struct)
	if !ok {
		panic("Expected struct type in compileInitStructWithValues")
	}

	var alloc llvmValue.Value

	// Allocate on the heap or on the stack
	if len(c.contextAlloc) > 0 && c.contextAlloc[len(c.contextAlloc)-1].Escapes {
		mallocatedSpaceRaw := c.contextBlock.NewCall(c.osFuncs.Malloc.Value.(llvmValue.Named), constant.NewInt(llvmTypes.I64, structType.Size()))
		alloc = c.contextBlock.NewBitCast(mallocatedSpaceRaw, llvmTypes.NewPointer(structType.LLVM()))
		structType.IsHeapAllocated = true
	} else {
		alloc = c.contextBlock.NewAlloca(structType.LLVM())
	}

	treType.Zero(c.contextBlock, alloc)

	for key, val := range v.Items {
		keyIndex, ok := structType.MemberIndexes[key]
		if !ok {
			panic("Unknown struct key: " + key)
		}

		itemPtr := c.contextBlock.NewGetElementPtr(pointer.ElemType(alloc), alloc, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, int64(keyIndex)))
		itemPtr.SetName(name.Var(key))

		compiledVal := c.compileValue(val)
		if compiledVal.IsVariable {
			loaded := c.contextBlock.NewLoad(compiledVal.Type.LLVM(), compiledVal.Value)
			c.contextBlock.NewStore(loaded, itemPtr)
		} else {
			c.contextBlock.NewStore(compiledVal.Value, itemPtr)
		}
	}

	return value.Value{
		Type:       structType,
		Value:      alloc,
		IsVariable: true,
	}
}
