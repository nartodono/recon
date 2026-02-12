package ui

import (
	"fmt"
	"time"
	"recon/internal/input"
	"recon/internal/modules/host"
	"recon/internal/modules/port"
)

func RunCommand(cmd string, args []string) bool {
	switch cmd {

	case "help", "?":
		PrintBanner()
		return false

	case "clear", "cls":
		ClearScreen()
		PrintBanner()
		return false

	case "exit", "quit", "q":
		fmt.Println("Bye...")
		return true

	case "host":
		return runHost(args)

	case "port":
		return runPort(args)

	default:
		fmt.Println(Red("[!] Unknown command.") + " Type '" + Cyan("help") + "' to see available commands.\n")
		return false
	}
}

func runHost(args []string) bool {
	if len(args) == 0 {
		fmt.Println(Yellow("[!] Usage: host <ip>  OR  host -f <file.txt>\n"))
		return false
	}

	// file mode
	if args[0] == "-f" {
		if len(args) != 2 {
			fmt.Println(Yellow("[!] Usage: host -f <file.txt>\n"))
			return false
		}

		targets, err := input.LoadTargetsFromFile(args[1])
		if err != nil {
			PrintError(err)
			return false
		}
		if len(targets) == 0 {
			fmt.Println(Yellow("[!] No valid targets found in file.\n"))
			return false
		}

		counts := HostCounts{}
		totalStart := time.Now()

		for i, t := range targets {
			start := time.Now()
			sp := NewSpinner()
			sp.Start(fmt.Sprintf("[*] Checking (%d/%d) %s ...", i+1, len(targets), t))

			res, err := host.Check(t)

			sp.Stop()
			elapsed := time.Since(start)

			if err != nil {
				PrintError(err)
				continue
			}

			RenderHostResult(res)
			fmt.Printf("    Time  : %.2fs\n\n", elapsed.Seconds())
			CountHostStatus(res, &counts)
		}

		PrintHostSummary(counts)
		fmt.Printf("Total Time: %.2fs\n\n", time.Since(totalStart).Seconds())
		return false
	}

	// single mode
	if len(args) != 1 {
		fmt.Println(Yellow("[!] Usage: host <ip-or-hostname>\n"))
		return false
	}

	start := time.Now()
	sp := NewSpinner()
	sp.Start(fmt.Sprintf("[*] Checking %s ...", args[0]))

	res, err := host.Check(args[0])

	sp.Stop()
	elapsed := time.Since(start)

	if err != nil {
		PrintError(err)
		return false
	}

	RenderHostResult(res)
	fmt.Printf("    Time  : %.2fs\n\n", elapsed.Seconds())
	return false
}

func runPort(args []string) bool {
	if len(args) == 0 {
		fmt.Println(Yellow("[!] Usage: port <ip>  OR  port <profile> <ip>  OR  port -f <file.txt>\n"))
		return false
	}

	// file mode (default profile)
	if args[0] == "-f" {
		if len(args) != 2 {
			fmt.Println(Yellow("[!] Usage: port -f <file.txt>\n"))
			return false
		}

		targets, err := input.LoadTargetsFromFile(args[1])
		if err != nil {
			PrintError(err)
			return false
		}
		if len(targets) == 0 {
			fmt.Println(Yellow("[!] No valid targets found in file.\n"))
			return false
		}

		totalStart := time.Now()

		for i, t := range targets {
			start := time.Now()
			sp := NewSpinner()
			sp.Start(fmt.Sprintf("[*] Port scan (%d/%d) %s ...", i+1, len(targets), t))

			// default profile for now
			res, err := port.Scan(t, []string{"-sC", "-sV"})

			sp.Stop()
			elapsed := time.Since(start)

			if err != nil {
				PrintError(err)
				continue
			}

			fmt.Println(Cyan("========================================"))
			fmt.Printf(Cyan("Target: %s\n\n"), t)

			RenderPortResult(res)
			fmt.Printf("    Time  : %.2fs\n\n", elapsed.Seconds())
		}

		fmt.Printf(Green("All scans completed in %.2fs\n\n"), time.Since(totalStart).Seconds())
		return false
	}

	// single/profile mode
	profile := "default"
	target := ""

	if len(args) == 1 {
		target = args[0] // port <ip>
	} else if len(args) == 2 {
		profile = args[0] // port <profile> <ip>
		target = args[1]
	} else {
		fmt.Println(Yellow("[!] Usage: port <ip>  OR  port <profile> <ip>\n"))
		return false
	}

	var extraArgs []string
	switch profile {
	case "default":
		extraArgs = []string{"-sC", "-sV"}
	case "aggr":
		extraArgs = []string{"-A"}
	default:
		fmt.Println(Yellow("[!] Unknown port profile. Available: default, aggr\n"))
		return false
	}

	start := time.Now()
	sp := NewSpinner()
	sp.Start(fmt.Sprintf("[*] Port scan (%s) %s ...", profile, target))

	res, err := port.Scan(target, extraArgs)

	sp.Stop()
	elapsed := time.Since(start)

	if err != nil {
		PrintError(err)
		return false
	}

	RenderPortResult(res)
	fmt.Printf("    Time  : %.2fs\n\n", elapsed.Seconds())
	return false
}
