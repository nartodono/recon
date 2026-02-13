package ui

import (
	"fmt"
	"github.com/nartodono/recon/internal/system"
)

func PrintError(err error) {
	switch err.(type) {
	case system.ResolveError:
		fmt.Println(Red("[!] " + err.Error()))
		fmt.Println("    Hint: check spelling or DNS settings (try an IP address).")
		fmt.Println()
	default:
		fmt.Println(Red("[!] " + err.Error()))
		fmt.Println()
	}
}
