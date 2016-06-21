package main

import (
	"os"

	"github.com/solher/kit-crud/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
