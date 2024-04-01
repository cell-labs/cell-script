package escape

import (
	"github.com/cell-labs/cell-script/compiler/parser"
)

// Escape performs variable escape analysis on variables allocated in functions
func Escape(input *parser.FileNode) *parser.FileNode {
	for _, ins := range input.Instructions {
		if defFunc, ok := ins.(*parser.DefineFuncNode); ok {

			// Name of the var mapped to their allocNode instruction index
			allocatedVars := map[string]int{}
			escapingVars := map[string]struct{}{}

			for insIndex, ins := range defFunc.Body {

				// Find all variables allocated in this function
				if allocIns, ok := ins.(*parser.AllocNode); ok {
					for _, name := range allocIns.Name {
						allocatedVars[name] = insIndex
					}
				}

				// Find all variables returned from this function
				if retIns, ok := ins.(*parser.ReturnNode); ok {
					for _, val := range retIns.Vals {
						findEscaping(escapingVars, val)
					}
				}
			}

			// Mark as escaping in the AST
			for escapingName := range escapingVars {
				if allocIndex, ok := allocatedVars[escapingName]; ok {
					allocIns := defFunc.Body[allocIndex].(*parser.AllocNode)
					allocIns.Escapes = true
					defFunc.Body[allocIndex] = allocIns
				}
			}
		}
	}

	return input
}

func findEscaping(escapingVars map[string]struct{}, ins parser.Node) {
	if retVariable, ok := ins.(*parser.NameNode); ok {
		escapingVars[retVariable.Name] = struct{}{}
		return
	}

	if retPtr, ok := ins.(*parser.GetReferenceNode); ok {
		findEscaping(escapingVars, retPtr.Item)
		return
	}

	/*if initStruct, ok := ins.(*parser.InitializeStructNode); ok {
		// initStruct.Items
	}*/
}
