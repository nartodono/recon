package ui

import (
	"fmt"
	"strings"
	"recon/internal/modules/port"
)

func RenderPortResult(r port.Result) {
	if len(r.Findings) == 0 {
		fmt.Println(Yellow("[!] No ports found (or host did not respond)."))
		fmt.Println()
		return
	}

	for _, f := range r.Findings {

		prefix := Yellow("[?]")
		if f.State == "OPEN" {
			prefix = Green("[+]")
		} else if f.State == "CLOSED" {
			prefix = Red("[-]")
		}

		fmt.Printf("%s Port %d ---------------------------\n", prefix, f.Port)
		fmt.Printf("    Protocol: %s\n", f.Proto)
		fmt.Printf("    Service : %s\n", f.Service)
		fmt.Printf("    Status  : %s\n", f.State)

		if strings.TrimSpace(f.Version) != "" {
		    fmt.Printf("    Version : %s\n", f.Version)
		}

		if len(f.Scripts) > 0 {
			fmt.Println()
			for _, s := range f.Scripts {
				fmt.Printf("    %s:\n", Cyan(s.ID))

				lines := strings.Split(s.Output, "\n")
				for _, line := range lines {
					line = strings.TrimRight(line, " \t")
					if line == "" {
						continue
					}
					fmt.Printf("      %s\n", line)
				}
				fmt.Println()
			}
		} else {
			fmt.Println()
		}
	}
}
