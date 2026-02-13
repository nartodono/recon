package main

import (
	"fmt"
	"os"
	"github.com/nartodono/recon/internal/system"
	"github.com/nartodono/recon/internal/ui"
)

func main() {
	// Dependency Check
	if err := system.CheckDeps(); err != nil {
		fmt.Fprintf(os.Stderr, "[!] %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
	cmd := os.Args[1]

		if cmd == "-h" || cmd == "--help" || cmd == "help" {
			ui.PrintHelp()
			return
		}
	
		args := os.Args[2:]
		ui.RunCommand(cmd, args)
		return
	}

ui.RunShell()
}
