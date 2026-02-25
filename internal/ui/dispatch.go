package ui

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/nartodono/recon/internal/export"
	"github.com/nartodono/recon/internal/input"
	"github.com/nartodono/recon/internal/modules/host"
	"github.com/nartodono/recon/internal/modules/port"
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
		infoWebservice()
		return false

	case "smtp":
		infoSmtp()
		return false

	case "mssql":
		infoMssql()
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
		fmt.Printf("Total Time: %.2fs\n\n", time.Since(totalStart).Seconds())

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

		return false
	}

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

	return false
}

// ---------------------------
// PORT
func runPort(args []string) bool {
	args, wantJSON, wantTXT := parseExportFlags(args)

	if len(args) == 0 {
		fmt.Println(Yellow("[!] Usage: port <ip/hostname>  OR  port <profile> <ip/hostname>  OR  port -f <file.txt>  OR  port <profile> -f <file.txt>\n"))
		return false
	}

	// ---------------------------
	// FILE MODE
	//   port -f <file>
	//   port <profile> -f <file>
	profile := "default"
	filePath := ""

	if args[0] == "-f" {
		if len(args) != 2 {
			fmt.Println(Yellow("[!] Usage: port -f <file.txt>\n"))
			return false
		}
		filePath = args[1]
	} else if len(args) == 3 && args[1] == "-f" {
		profile = args[0]
		filePath = args[2]
	}

	if filePath != "" {
		targets, err := input.LoadTargetsFromFile(filePath)
		if err != nil {
			PrintError(err)
			return false
		}
		if len(targets) == 0 {
			fmt.Println(Yellow("[!] No targets found in file.\n"))
			return false
		}

		extraArgs, ok := portExtraArgs(profile)
		if !ok {
			fmt.Println(Yellow("[!] Unknown port profile. Type 'help' to see available commands.\n"))
			return false
		}

		// limits: non-deep = 30, deep = 10
		// maxTargets := 30
		isDeep := strings.Contains(profile, "deep")
		if isDeep {
			// maxTargets = 10
			fmt.Println(Yellow("[!] Deep profile selected in file mode. This may take a long time per target."))
		}

		// if len(targets) > maxTargets {
		// 	fmt.Printf(Yellow("[!] File contains %d targets. Limiting to first %d targets.\n\n"), len(targets), maxTargets)
		// 	targets = targets[:maxTargets]
		// }

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

		return false
	}

	// ---------------------------
	// SINGLE MODE
	//   port <ip>
	//   port <profile> <ip>
	target := ""
	if len(args) == 1 {
		target = args[0]
	} else if len(args) == 2 {
		profile = args[0]
		target = args[1]
	} else {
		fmt.Println(Yellow("[!] Usage: port <ip/hostname>  OR  port <profile> <ip/hostname>\n"))
		return false
	}

	extraArgs, ok := portExtraArgs(profile)
	if !ok {
		fmt.Println(Yellow("[!] Unknown port profile.\n"))
		return false
	}

	// timeout := 8 * time.Minute
	if strings.Contains(profile, "deep") {
		// timeout = 25 * time.Minute
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

	return false
}

// portExtraArgs maps a profile name to nmap arguments (excluding -Pn -oX - <target>).
func portExtraArgs(profile string) ([]string, bool) {
	switch profile {
		case "default":
			return []string{"-sC", "-sV"}, true

		case "aggr":
			return []string{
				"-A",
				"--host-timeout", "10m",
				"--script-timeout", "90s",
				"--max-retries", "2",
				"-T4",
			}, true

		case "common":
			return []string{
				"-sV",
				"--top-ports", "1000",
				"--version-light",
				"--max-retries", "2",
				"-T4",
			}, true

		case "deep":

			return []string{
				"-sC",
				"-sV",
				"--script", "(default or safe or discovery) and not (dos or intrusive or exploit or brute)",
				"--script-timeout", "90s",
				"--host-timeout", "10m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "ftp":
			return []string{
				"-p", "21", "-sV",
				"--script", "ftp-anon,ftp-syst,ftp-bounce",
				"--script-timeout", "60s",
				"--host-timeout", "5m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "ftp-deep":
			return []string{
				"-p", "21", "-sV",
				"--script", "(ftp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "8m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "ssh":
			return []string{
				"-p", "22", "-sV",
				"--script", "ssh-hostkey,ssh2-enum-algos,ssh-auth-methods,banner",
				"--script-timeout", "60s",
				"--host-timeout", "5m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "ssh-deep":
			return []string{
				"-p", "22", "-sV",
				"--script", "(ssh-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "8m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "smtp":
			return []string{
				"-p", "25,587", "-sV",
				"--script", "smtp-commands,smtp-enum-users",
				"--script-timeout", "60s",
				"--host-timeout", "6m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "smtp-deep":
			return []string{
				"-p", "25,587", "-sV",
				"--script", "(smtp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "8m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "dns":
			return []string{
				"-p", "53", "-sV",
				"--script", "dns-nsid,dns-recursion",
				"--script-timeout", "45s",
				"--host-timeout", "4m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "dns-deep":
			return []string{
				"-p", "53", "-sV",
				"--script", "(dns-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "60s",
				"--host-timeout", "6m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "web":
			return []string{
				"-p", "80,443", "-sV",
				"--script", "http-title,http-headers,http-methods,http-enum,http-server-header",
				"--script-timeout", "60s",
				"--host-timeout", "6m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "web-deep":
			return []string{
				"-p", "80,443", "-sV",
				"--script", "(http-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "9m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "kerberos":
			return []string{
				"-p", "88", "-sV",
				"--script", "krb5-enum-users",
				"--script-timeout", "60s",
				"--host-timeout", "5m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "kerberos-deep":
			return []string{
				"-p", "88", "-sV",
				"--script", "(krb5-* and (safe or default or discovery)) and not (brute or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "8m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "snmp":
			return []string{
				"-sU", "-p", "161", "-sV",
				"--script", "snmp-info,snmp-sysdescr,snmp-interfaces",
				"--script-timeout", "45s",
				"--host-timeout", "4m",
				"--max-retries", "1",
				"-T4",
			}, true
		case "snmp-deep":
			return []string{
				"-sU", "-p", "161", "-sV",
				"--script", "(snmp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "60s",
				"--host-timeout", "5m",
				"--max-retries", "1",
				"-T4",
			}, true

		case "ldap":
			return []string{
				"-p", "389", "-sV",
				"--script", "ldap-rootdse,ldap-search",
				"--script-timeout", "60s",
				"--host-timeout", "6m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "ldap-deep":
			return []string{
				"-p", "389", "-sV",
				"--script", "(ldap-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "9m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "smb":
			return []string{
				"-p", "445", "-sV",
				"--script", "smb-os-discovery,smb2-security-mode,smb2-time,smb-protocols",
				"--script-timeout", "60s",
				"--host-timeout", "6m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "smb-deep":
			return []string{
				"-p", "445", "-sV",
				"--script", "(smb-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "10m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "mssql":
			return []string{
				"-p", "1433", "-sV",
				"--script", "ms-sql-info,ms-sql-ntlm-info",
				"--script-timeout", "60s",
				"--host-timeout", "6m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "mssql-deep":
			return []string{
				"-p", "1433", "-sV",
				"--script", "(ms-sql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "9m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "mysql":
			return []string{
				"-p", "3306", "-sV",
				"--script", "mysql-info,mysql-capabilities",
				"--script-timeout", "60s",
				"--host-timeout", "6m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "mysql-deep":
			return []string{
				"-p", "3306", "-sV",
				"--script", "(mysql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "9m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "rdp":
			return []string{
				"-p", "3389", "-sV",
				"--script", "rdp-ntlm-info,rdp-enum-encryption",
				"--script-timeout", "60s",
				"--host-timeout", "6m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "rdp-deep":
			return []string{
				"-p", "3389", "-sV",
				"--script", "(rdp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "9m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "postgresql":
			return []string{
				"-p", "5432", "-sV",
				"--script", "pgsql-info",
				"--script-timeout", "60s",
				"--host-timeout", "6m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "postgresql-deep":
			return []string{
				"-p", "5432", "-sV",
				"--script", "(pgsql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "9m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "vnc":
			return []string{
				"-p", "5900", "-sV",
				"--script", "vnc-info",
				"--script-timeout", "60s",
				"--host-timeout", "6m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "vnc-deep":
			return []string{
				"-p", "5900", "-sV",
				"--script", "(vnc-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "9m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "winrm":
			return []string{
				"-p", "5985,5986", "-sV",
				"--script", "wsman-info",
				"--script-timeout", "60s",
				"--host-timeout", "6m",
				"--max-retries", "2",
				"-T4",
			}, true
		case "winrm-deep":
			return []string{
				"-p", "5985,5986", "-sV",
				"--script", "(wsman-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
				"--script-timeout", "90s",
				"--host-timeout", "9m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "vuln":
			// Basic balanced vuln
			return []string{
				"-sV",
				"--version-light",
				"--script", "vuln",
				"--host-timeout", "20m",
				"--max-retries", "2",
				"-T4",
			}, true

		case "vuln-deep":
			// Allowed to be longer, but still guardrailed.
			return []string{
				"-sV",
				"--version-light",
				"--script", "vuln or exploit",
				"--script-timeout", "3m",
				"--host-timeout", "30m",
				"--max-retries", "2",
				"-T4",
			}, true

		default:
			return nil, false
	}
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
