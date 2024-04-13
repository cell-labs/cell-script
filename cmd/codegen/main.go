package main

import (
	"log"
	"os"

	"github.com/cell-labs/cell-script/internal/compiler"
)

func init() {
	compiler.SetupOptions()
}

func main() {
	options := compiler.ParseOptions()
	options.Stage = compiler.STAGE_CODEGENED
	err := compiler.Run(options)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
