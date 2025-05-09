// Package main implements the main program functions.
package main

import (
	"fmt"
	"os"

	"github.com/stvmln86/glint/glint/comms"
	"github.com/stvmln86/glint/glint/items/book"
	"github.com/stvmln86/glint/glint/tools/clui"
)

var Stdout = os.Stdout

func try(err error) {
	if err != nil {
		fmt.Fprintf(Stdout, "Error: %s.\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	dire, err := clui.Env("GLINT_DIR")
	try(err)

	extn, err := clui.Env("GLINT_EXT")
	try(err)

	book := book.New(dire, extn, 0666)
	try(comms.Run(Stdout, book, os.Args[1:]))
}
