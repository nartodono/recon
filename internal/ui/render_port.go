package ui

import (
	"fmt"
	"strings"

	"github.com/nartodono/recon/internal/modules/port"
)

func RenderPortResult(r port.Result) {
	fmt.Printf("Target: %s\n", r.Target)

	printedMeta := false
	if r.HostUp {
		if r.LatencySec > 0 {
			fmt.Printf("Host is up (%.1fs latency).\n", r.LatencySec)
		} else {
			fmt.Println("Host is up.")
		}
		printedMeta = true
	}
	if strings.TrimSpace(r.NotShown) != "" {
		fmt.Println(r.NotShown)
		printedMeta = true
	}
	if printedMeta {
		fmt.Println()
	}

	if strings.TrimSpace(r.Warning) != "" {
		for _, line := range strings.Split(r.Warning, "\n") {
			line = strings.TrimRight(line, " \t")
			if line == "" {
				continue
			}
			fmt.Println(Yellow(line))
		}
		fmt.Println()
	}

	if len(r.Findings) == 0 {
		if strings.TrimSpace(r.Warning) == "" {
			fmt.Println(Yellow("[!] No ports found (or host did not respond)."))
			fmt.Println()
		}
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
		} else {
			fmt.Printf("    Version : UNKNOWN\n")
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

	if strings.TrimSpace(r.ServiceInfo) != "" {
		s := strings.TrimSpace(r.ServiceInfo)

		if strings.HasPrefix(s, "Service Info:") {
			body := strings.TrimSpace(strings.TrimPrefix(s, "Service Info:"))
			parts := strings.Split(body, ";")

			osLine := ""
			cpeLine := ""
			for _, p := range parts {
				p = strings.TrimSpace(p)
				if strings.HasPrefix(p, "OS:") {
					osLine = strings.TrimSpace(p)
				} else if strings.HasPrefix(p, "CPE:") {
					cpeLine = strings.TrimSpace(p)
				}
			}

			if osLine != "" {
				fmt.Printf("Service Info: %s\n", osLine)
			} else {
				fmt.Println(s)
			}

			if cpeLine != "" {
				cpeBody := strings.TrimSpace(strings.TrimPrefix(cpeLine, "CPE:"))
				cpes := strings.Split(cpeBody, ",")
				clean := make([]string, 0, len(cpes))
				for _, c := range cpes {
					c = strings.TrimSpace(c)
					if c != "" {
						clean = append(clean, c)
					}
				}
				if len(clean) > 0 {
					fmt.Println("CPE:")
					for _, c := range clean {
						fmt.Println("  - " + c)
					}
				}
			}
		} else {
			fmt.Println(s)
		}
		fmt.Println()
	}
}
