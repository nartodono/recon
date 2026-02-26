package ui

import "fmt"

func PrintBanner() {
	fmt.Println(ColorCyan + `
██████╗ ███████╗ ██████╗ ██████╗ ███╗   ██╗
██╔══██╗██╔════╝██╔════╝██╔═══██╗████╗  ██║
██████╔╝█████╗  ██║     ██║   ██║██╔██╗ ██║
██╔══██╗██╔══╝  ██║     ██║   ██║██║╚██╗██║
██║  ██║███████╗╚██████╗╚██████╔╝██║ ╚████║
╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝
` + ColorReset)
	// fmt.Println()
	fmt.Printf("  %s               - Host status (ping + discovery scan)\n", Green("host <ip>"))
	fmt.Printf("  %s     - Run port scan profile\n", Green("port <profile> <ip>"))
	fmt.Printf("  %s          - Show available port profiles\n", Cyan("profile | list"))
	fmt.Printf("  %s          - Show enumeration reference\n", Cyan("info <service>"))
	fmt.Printf("  %s                - Show commands\n", Cyan("help | ?"))
	fmt.Printf("  %s             - Clear screen\n", Cyan("clear | cls"))
	fmt.Printf("  %s         - Exit Recon\n", Cyan("exit | quit | q"))
	fmt.Println()
}

func PrintBannerHelp() {
	fmt.Println(`
Recon supports two modes:

────────────────────────────────────────
1) INTERACTIVE SHELL MODE
────────────────────────────────────────
Launch Recon without arguments to enter interactive mode

Available commands:

  host <ip>
      Check host availability (ping + discovery scan)

  port <profile> <ip> [-p <port>]
      Run port scan profile against a target

  profile | list
      Show available port profiles

  info <service>
      Show enumeration & attack reference

  help | ?
      Show this help message

  clear | cls
      Clear screen

  exit | quit | q
      Exit Recon


────────────────────────────────────────
2) CLI SHORTCUT MODE
────────────────────────────────────────
Run directly from terminal:

  recon host <ip> [--txt] [--json]
  recon host -f <file> [--txt] [--json]

  recon port <profile> <ip> [--txt] [--json]
  recon port <profile> <ip> -p <port> [--txt] [--json]
  recon port <profile> -f <file> [--txt] [--json]

  recon info <service>


────────────────────────────────────────
PORT OPTIONS
────────────────────────────────────────
  -p <port>
      Override default profile port
      Example: -p 8080 or -p 80,443


────────────────────────────────────────
FILE INPUT
────────────────────────────────────────
  -f <file>
      Scan multiple targets from file
      Format: one IP per line


────────────────────────────────────────
OUTPUT OPTIONS
────────────────────────────────────────
  --txt
      Print formatted text output

  --json
      Print structured JSON output

Both --txt and --json can be used together.


────────────────────────────────────────
PORT PROFILES
────────────────────────────────────────

Standard & Deep Profiles
────────────────────────────────────────
  default        | aggr
  common         | deep
  ftp            | ftp-deep
  ssh            | ssh-deep
  smtp           | smtp-deep
  dns            | dns-deep
  web            | web-deep
  kerberos       | kerberos-deep
  snmp           | snmp-deep
  ldap           | ldap-deep
  smb            | smb-deep
  mssql          | mssql-deep
  mysql          | mysql-deep
  postgresql     | postgresql-deep
  rdp            | rdp-deep
  vnc            | vnc-deep
  winrm          | winrm-deep


Vulnerability Modes
────────────────────────────────────────
  vuln        - Safe vulnerability checks
  vuln-deep   - Aggressive checks (may include intrusive/DoS)


Note:
If no profile is specified, "default" mode will be used.


────────────────────────────────────────
EXAMPLES
────────────────────────────────────────
  recon host 192.168.1.10
  recon host -f hosts.txt --json

  recon port smb 192.168.1.20
  recon port web-deep 192.168.1.20 --txt

  recon port vuln 192.168.1.20
  recon port vuln-deep -f targets.txt --txt --json

For detailed module documentation, command references,
and internal Nmap mappings, please visit:

  https://github.com/nartodono/recon

────────────────────────────────────────
Notes
────────────────────────────────────────
Currently, recon port scanning in deep mode is limited to a
maximum of 10 IPs (deep, -deep) and 30 IPs for normal mode
(except -deep, including default and common) to ensure the scan
completes successfully and results can be properly displayed.
`)
}

func PrintHelp() {
	fmt.Println(ColorCyan + `
██████╗ ███████╗ ██████╗ ██████╗ ███╗   ██╗
██╔══██╗██╔════╝██╔════╝██╔═══██╗████╗  ██║
██████╔╝█████╗  ██║     ██║   ██║██╔██╗ ██║
██╔══██╗██╔══╝  ██║     ██║   ██║██║╚██╗██║
██║  ██║███████╗╚██████╗╚██████╔╝██║ ╚████║
╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚══════╝ ╚═╝  ╚═══╝
` + ColorReset)

fmt.Println(`
GENERAL SYNTAX
────────────────────────────────────────

Host scan:
  recon host <ip> [--txt] [--json]
  recon host -f <file> [--txt] [--json]

Port scan:
  recon port <profile> <ip> [-p <port>] [--txt] [--json]
  recon port <profile> -f <file> [--txt] [--json]

Service reference:
  recon info <service>

Profile list:
  recon profile
  recon list


INTERACTIVE MODE
────────────────────────────────────────
  recon


OPTIONS
────────────────────────────────────────
  -p <port>    Override default profile port
  -f <file>    Scan multiple targets (one IP per line)

  --txt        Print formatted text output
  --json       Print structured JSON output
`)
}

func PrintProfile() {
	fmt.Println(`
────────────────────────────────────────
PORT PROFILES
────────────────────────────────────────

Standard & Deep Profiles
────────────────────────────────────────
  default        | aggr
  common         | deep
  ftp            | ftp-deep
  ssh            | ssh-deep
  smtp           | smtp-deep
  dns            | dns-deep
  web            | web-deep
  kerberos       | kerberos-deep
  snmp           | snmp-deep
  ldap           | ldap-deep
  smb            | smb-deep
  mssql          | mssql-deep
  mysql          | mysql-deep
  postgresql     | postgresql-deep
  rdp            | rdp-deep
  vnc            | vnc-deep
  winrm          | winrm-deep


Vulnerability Modes
────────────────────────────────────────
  vuln        - Safe vulnerability checks
  vuln-deep   - Aggressive checks (may include intrusive/DoS)


Note:
If no profile is specified, "default" mode will be used.
`)
}
