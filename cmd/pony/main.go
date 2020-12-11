package main

// TODO: add version command
// TODO: implement strict verification

import (
	"os"

	"github.com/antklim/pony/commands"
)

func main() {
	if err := commands.Execute(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
