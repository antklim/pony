// Pony is a simple static site generator and server.
//
// Usage
//
// To start pony as a server, simply run:
//
//  pony
//
// For more specific usage information, refer to the help doc `pony -h`:
//  Usage:
//    pony [command] [flags]
//
//  Available commands:
//    build    Build static pages
//    run      Run server to preview pages
//    verify   Verifies that site metadata complies with template
//
//  Flags:
//		-s, --strict		Metadata and templates should be strictly matched [build|run]
//    -v, --version		Print version info and exit
package main

import (
	"fmt"
	"os"

	"github.com/antklim/pony/commands"
)

func main() {
	if err := commands.Execute(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
