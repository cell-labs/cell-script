package intrisinc

import (
	"fmt"

	"github.com/cell-labs/cell-script/compiler/parser"
)

var collector = map[string]parser.TypeNode{}

func GetTypeNodeByName(name string) parser.TypeNode {
	v, ok := collector[name]
	if !ok {
		panic("the node is not exist")
	}
	return v
}

// Escape performs variable escape analysis on variables allocated in functions
func BigIntCollect(input *parser.FileNode) *parser.FileNode {
	for _, ins := range input.Instructions {
		if td, ok := ins.(*parser.DefineTypeNode); ok {
			if s, ok := td.Type.(*parser.StructTypeNode); ok {
				if s.GetName() == "bigint" {
					collector["bigint"] = s
					fmt.Println("gotcha")
				}
			}
		}
	}
	return input
}
