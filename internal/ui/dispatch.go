package ui

import (
	"fmt"
	"github.com/nartodono/recon/internal/export"
	"github.com/nartodono/recon/internal/input"
	"github.com/nartodono/recon/internal/modules/host"
	"github.com/nartodono/recon/internal/modules/port"
	"path/filepath"
	"strings"
	"time"
)

func RunCommand(cmd string, args []string) bool {
	switch cmd {

	case "profile", "list":
		PrintProfile()
		return false

	case "help", "-h", "?":
		PrintBannerHelp()
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

	case "info":
		return runInfo(args)

	default:
		fmt.Println(Red("[!] Unknown command.") + " Type '" + Cyan("help") + "' to see available commands.\n")
		return false
	}
}

func runInfo(args []string) bool {
	if len(args) == 0 {
		fmt.Println(Yellow("[!] Usage: info <smb|ssh|snmp|ldap|rdp|web|smtp|mssql>"))
		return false
	}

	switch args[0] {
	case "smb":
		infoSmb()
		return false

	case "ssh":
		infoSsh()
		return false

	case "snmp":
		infoSnmp()
		return false

	case "ldap":
		infoLdap()
		return false

	case "rdp":
		infoRdp()
		return false

	case "web":
		infoWebService()
		return false

	case "smtp":
		infoSmtp()
		return false

	case "mssql":
		infoMssql()
		return false

	case "kerberos":
		infoKerberos()
		return false

	default:
		fmt.Println(Red("[!] Unknown info topic.") + " Try: " + Cyan("info smb") + ", " + Cyan("info ssh") + ", ...")
		return false
	}
}

// ---------------------------
// HOST
func runHost(args []string) bool {
	args, wantJSON, wantTXT := parseExportFlags(args)

	if len(args) == 0 {
		fmt.Println(Yellow("[!] Usage: host <ip/hostname>  OR  host -f <file.txt>\n"))
		return false
	}

	// file mode
	if args[0] == "-f" {
		if len(args) != 2 {
			fmt.Println(Yellow("[!] Usage: host -f <file.txt>\n"))
			return false
		}
		return MultiHost(args[1], wantJSON, wantTXT)
	}

	// single mode
	if len(args) != 1 {
		fmt.Println(Yellow("[!] Usage: host <ip-or-hostname>\n"))
		return false
	}

	return SingleHost(args[0], wantJSON, wantTXT)
}

func SingleHost(target string, wantJSON, wantTXT bool) bool {
	start := time.Now()
	sp := NewSpinner()
	sp.Start(fmt.Sprintf("[*] Checking %s ...", target))

	res, err := host.Check(target)

	sp.Stop()
	elapsed := time.Since(start)

	if err != nil {
		PrintError(err)
		return false
	}

	RenderHostResult(res)
	fmt.Printf("    Time  : %.2fs\n\n", elapsed.Seconds())

	if wantJSON || wantTXT {
		dir, derr := export.DefaultDir()
		if derr != nil {
			PrintError(derr)
			return false
		}
		if derr := export.EnsureDir(dir); derr != nil {
			PrintError(derr)
			return false
		}

		now := time.Now()
		jsonPayload := map[string]any{
			"module":          "host",
			"timestamp":       now.Format(time.RFC3339),
			"mode":            "single",
			"target":          res.Target,
			"result":          res,
			"elapsed_seconds": elapsed.Seconds(),
		}

		if wantJSON {
			p := filepath.Join(dir, export.Filename("host", "json", now))
			if e := export.WriteJSON(p, jsonPayload); e != nil {
				PrintError(e)
			} else {
				PrintSaved(p)
			}
		}

		if wantTXT {
			p := filepath.Join(dir, export.Filename("host", "txt", now))
			txt := export.HostSingleTXT(res, elapsed.Seconds(), now)
			if e := export.WriteFile(p, []byte(txt)); e != nil {
				PrintError(e)
			} else {
				PrintSaved(p)
			}
		}

		fmt.Println()
	}

	return true
}

func MultiHost(filePath string, wantJSON, wantTXT bool) bool {
	targets, err := input.LoadTargetsFromFile(filePath)
	if err != nil {
		PrintError(err)
		return false
	}
	if len(targets) == 0 {
		fmt.Println(Yellow("[!] No targets found in file.\n"))
		return false
	}

	counts := HostCounts{}
	totalStart := time.Now()
	all := []host.Result{}

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
		all = append(all, res)

		fmt.Printf("    Time  : %.2fs\n\n", elapsed.Seconds())
		CountHostStatus(res, &counts)
	}

	PrintHostSummary(counts)
	totalElapsed := time.Since(totalStart).Seconds()
	fmt.Printf("Total Time: %.2fs\n\n", totalElapsed)

	if wantJSON || wantTXT {
		dir, derr := export.DefaultDir()
		if derr != nil {
			PrintError(derr)
			return false
		}
		if derr := export.EnsureDir(dir); derr != nil {
			PrintError(derr)
			return false
		}

		now := time.Now()
		jsonPayload := map[string]any{
			"module":    "host",
			"timestamp": now.Format(time.RFC3339),
			"mode":      "file",
			"results":   all,
			"summary": map[string]int{
				"up":      counts.Up,
				"down":    counts.Down,
				"unknown": counts.Unknown,
				"total":   counts.Total,
			},
			"elapsed_seconds": totalElapsed,
		}

		if wantJSON {
			p := filepath.Join(dir, export.Filename("host", "json", now))
			if e := export.WriteJSON(p, jsonPayload); e != nil {
				PrintError(e)
			} else {
				PrintSaved(p)
			}
		}

		if wantTXT {
			p := filepath.Join(dir, export.Filename("host", "txt", now))
			txt := export.HostFileTXT(all, counts.Up, counts.Down, counts.Unknown, counts.Total, totalElapsed, now)
			if e := export.WriteFile(p, []byte(txt)); e != nil {
				PrintError(e)
			} else {
				PrintSaved(p)
			}
		}

		fmt.Println()
	}

	return true
}

// ---------------------------
// PORT

func runPort(args []string) bool {
	args, wantJSON, wantTXT := parseExportFlags(args)

	if len(args) == 0 {
		fmt.Println(Yellow("[!] Usage: port (service) <ip/host> (-f <file>) (-p <ports>)\n"))
		return false
	}

	// Service/profile is optional but must be the first token if present.
	service := "default"
	if isPortService(args[0]) {
		service = args[0]
		args = args[1:]
	}

	var (
		target string
		file   string
		ports  string
	)

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-f":
			if i+1 >= len(args) || strings.HasPrefix(args[i+1], "-") {
				fmt.Println(Yellow("[!] Usage: port ... -f <file>\n"))
				return false
			}
			file = args[i+1]
			i++

		case "-p":
			if i+1 >= len(args) || strings.HasPrefix(args[i+1], "-") {
				fmt.Println(Yellow("[!] Usage: port ... -p <ports>\n"))
				return false
			}
			ports = args[i+1]
			i++

		default:
			if strings.HasPrefix(args[i], "-") {
				fmt.Println(Yellow("[!] Unknown flag: " + args[i] + "\n"))
				return false
			}
			if target == "" {
				target = args[i]
			} else {
				fmt.Println(Yellow("[!] Unexpected argument: " + args[i] + "\n"))
				return false
			}
		}
	}

	// Mode selection:
	// - If -f is present => file mode; do not allow an additional target token.
	// - If -f is absent  => single mode; target is required.
	if file != "" {
		if target != "" {
			fmt.Println(Yellow("[!] Usage: port (service) -f <file> (-p <ports>)\n"))
			return false
		}
		return MultiPort(service, file, ports, wantJSON, wantTXT)
	}

	if target == "" {
		fmt.Println(Yellow("[!] Usage: port (service) <ip/host> (-p <ports>)\n"))
		return false
	}

	return SinglePort(service, target, ports, wantJSON, wantTXT)
}

func SinglePort(profile, target, portOverride string, wantJSON, wantTXT bool) bool {
	baseArgs, defaultPorts, ok := portExtraArgs(profile)
	if !ok {
		fmt.Println(Yellow("[!] Unknown port profile. Type 'help' to see available commands.\n"))
		return false
	}

	// Port override wins; otherwise use profile default ports (if any).
	effectivePorts := portOverride
	if effectivePorts == "" {
		effectivePorts = defaultPorts
	}

	extraArgs := append([]string{}, baseArgs...)
	if effectivePorts != "" {
		extraArgs = append(extraArgs, "-p", effectivePorts)
	}

	if strings.Contains(profile, "deep") {
		fmt.Println(Yellow("[!] Deep profile selected. This may take a long time..."))
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

	if wantJSON || wantTXT {
		dir, derr := export.DefaultDir()
		if derr != nil {
			PrintError(derr)
			return false
		}
		if derr := export.EnsureDir(dir); derr != nil {
			PrintError(derr)
			return false
		}

		now := time.Now()

		jsonPayload := map[string]any{
			"module":          "port",
			"timestamp":       now.Format(time.RFC3339),
			"mode":            "single",
			"target":          res.Target,
			"profile":         profile,
			"result":          res,
			"elapsed_seconds": elapsed.Seconds(),
		}

		if wantJSON {
			p := filepath.Join(dir, export.Filename("port", "json", now))
			if e := export.WriteJSON(p, jsonPayload); e != nil {
				PrintError(e)
			} else {
				PrintSaved(p)
			}
		}

		if wantTXT {
			p := filepath.Join(dir, export.Filename("port", "txt", now))
			txt := export.PortSingleTXT(res, profile, elapsed.Seconds(), now)
			if e := export.WriteFile(p, []byte(txt)); e != nil {
				PrintError(e)
			} else {
				PrintSaved(p)
			}
		}

		fmt.Println()
	}

	return true
}

func MultiPort(profile, filePath, portOverride string, wantJSON, wantTXT bool) bool {
	targets, err := input.LoadTargetsFromFile(filePath)
	if err != nil {
		PrintError(err)
		return false
	}
	if len(targets) == 0 {
		fmt.Println(Yellow("[!] No targets found in file.\n"))
		return false
	}

	baseArgs, defaultPorts, ok := portExtraArgs(profile)
	if !ok {
		fmt.Println(Yellow("[!] Unknown port profile. Type 'help' to see available commands.\n"))
		return false
	}

	effectivePorts := portOverride
	if effectivePorts == "" {
		effectivePorts = defaultPorts
	}

	extraArgs := append([]string{}, baseArgs...)
	if effectivePorts != "" {
		extraArgs = append(extraArgs, "-p", effectivePorts)
	}

	isDeep := strings.Contains(profile, "deep")
	if isDeep {
		fmt.Println(Yellow("[!] Deep profile selected in file mode. This may take a long time per target."))
	}

	if len(targets) > 10 {
		if isDeep {
			fmt.Printf(Yellow(
				"[!] %d hosts detected.\n"+
					"[!] Deep profile is enabled.\n"+
					"[!] Expect substantially longer execution time.\n\n",
			), len(targets))
		} else {
			fmt.Printf(Yellow(
				"[!] %d hosts detected.\n"+
					"[!] Scan time will scale with the number of hosts.\n\n",
			), len(targets))
		}
	}

	totalStart := time.Now()
	items := []export.PortFileItem{}

	for i, t := range targets {
		start := time.Now()
		sp := NewSpinner()
		sp.Start(fmt.Sprintf("[*] Port scan (%s) (%d/%d) %s ...", profile, i+1, len(targets), t))

		res, err := port.Scan(t, extraArgs)

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

		items = append(items, export.PortFileItem{
			Target:         t,
			Findings:       res.Findings,
			ElapsedSeconds: elapsed.Seconds(),
		})
	}

	fmt.Printf(Green("All scans completed in %.2fs\n\n"), time.Since(totalStart).Seconds())

	if wantJSON || wantTXT {
		dir, derr := export.DefaultDir()
		if derr != nil {
			PrintError(derr)
			return false
		}
		if derr := export.EnsureDir(dir); derr != nil {
			PrintError(derr)
			return false
		}

		now := time.Now()
		totalElapsed := time.Since(totalStart).Seconds()

		jsonPayload := map[string]any{
			"module":          "port",
			"timestamp":       now.Format(time.RFC3339),
			"mode":            "file",
			"profile":         profile,
			"results":         items,
			"elapsed_seconds": totalElapsed,
		}

		if wantJSON {
			p := filepath.Join(dir, export.Filename("port", "json", now))
			if e := export.WriteJSON(p, jsonPayload); e != nil {
				PrintError(e)
			} else {
				PrintSaved(p)
			}
		}

		if wantTXT {
			p := filepath.Join(dir, export.Filename("port", "txt", now))
			txt := export.PortFileTXT(items, profile, totalElapsed, now)
			if e := export.WriteFile(p, []byte(txt)); e != nil {
				PrintError(e)
			} else {
				PrintSaved(p)
			}
		}

		fmt.Println()
	}

	return true
}

// ---------------------------
// Port TXT formatter (no color)
type portFileItemForTXT struct {
	Target         string
	Findings       []port.PortFinding
	ElapsedSeconds float64
}

func portSingleTXT(r port.Result, profile string, elapsedSeconds float64, t time.Time) string {
	var sb strings.Builder
	sb.WriteString("=== recon port ===\n")
	sb.WriteString(fmt.Sprintf("Time    : %s\n", t.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("Target  : %s\n", r.Target))
	sb.WriteString(fmt.Sprintf("Profile : %s\n\n", profile))

	sb.WriteString(renderPortFindingsTXT(r.Findings))
	sb.WriteString(fmt.Sprintf("\nTime  : %.2fs\n\n", elapsedSeconds))
	return sb.String()
}

func portFileTXT(items any, profile string, totalElapsedSeconds float64, t time.Time) string {
	typed := []portFileItemForTXT{}

	switch v := items.(type) {
	case []struct {
		Target         string
		Findings       []port.PortFinding
		ElapsedSeconds float64
	}:
		for _, it := range v {
			typed = append(typed, portFileItemForTXT{
				Target:         it.Target,
				Findings:       it.Findings,
				ElapsedSeconds: it.ElapsedSeconds,
			})
		}
	default:

	}

	var sb strings.Builder
	sb.WriteString("=== recon port ===\n")
	sb.WriteString(fmt.Sprintf("Time    : %s\n", t.Format(time.RFC3339)))
	sb.WriteString("Mode    : file\n")
	sb.WriteString(fmt.Sprintf("Profile : %s\n\n", profile))

	for _, it := range typed {
		sb.WriteString("========================================\n")
		sb.WriteString(fmt.Sprintf("Target: %s\n\n", it.Target))
		sb.WriteString(renderPortFindingsTXT(it.Findings))
		sb.WriteString(fmt.Sprintf("\nTime  : %.2fs\n\n", it.ElapsedSeconds))
	}

	sb.WriteString(fmt.Sprintf("Total Time: %.2fs\n\n", totalElapsedSeconds))
	return sb.String()
}

func renderPortFindingsTXT(findings []port.PortFinding) string {
	var sb strings.Builder
	if len(findings) == 0 {
		sb.WriteString("[!] No ports found (or host did not respond).\n")
		return sb.String()
	}

	for _, f := range findings {
		prefix := "?"
		if f.State == "OPEN" {
			prefix = "+"
		} else if f.State == "CLOSED" {
			prefix = "-"
		}

		sb.WriteString(fmt.Sprintf("[%s] Port %d ---------------------------\n", prefix, f.Port))
		sb.WriteString(fmt.Sprintf("    %s - %s\n", f.Proto, f.Service))
		sb.WriteString(fmt.Sprintf("    Status : %s\n", f.State))
		if strings.TrimSpace(f.Version) != "" {
			sb.WriteString(fmt.Sprintf("    Version: %s\n", f.Version))
		}

		if len(f.Scripts) > 0 {
			sb.WriteString("\n")
			for _, s := range f.Scripts {
				sb.WriteString(fmt.Sprintf("    %s:\n", s.ID))
				for _, line := range strings.Split(s.Output, "\n") {
					line = strings.TrimRight(line, " \t")
					if line == "" {
						continue
					}
					sb.WriteString("      " + line + "\n")
				}
				sb.WriteString("\n")
			}
		} else {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
