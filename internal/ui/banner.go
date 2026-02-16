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

	fmt.Println("  " + Green("host") + " <ip>            - Host status (ping + nmap -sn)")
	fmt.Println("  " + Green("port") + " <ip>            - Port checker (nmap)")
	fmt.Println("  " + Cyan("profile  / list") + "      - Show Port Profile Lists")
	fmt.Println("  " + Cyan("help / ?") + "             - Show commands")
	fmt.Println("  " + Cyan("clear / cls") + "          - Clear screen")
	fmt.Println("  " + Cyan("exi / q t") + "            - Quit")
	fmt.Println()
}

func PrintBannerHelp() {
	fmt.Println(`
Recon supports two modes:

────────────────────────────────────────
1) INTERACTIVE SHELL MODE
────────────────────────────────────────
Launch Recon without arguments to enter interactive mode

Available commands inside the shell:

  host <ip>
      Check host status (ping + nmap -sn)

  port <profile> <ip>
      Run a specific port profile against a target

  profile / list
      Show available port scanning profile

  help / ?
      Show help information

  clear / cls
      Clear the screen

  exit / quit / q
      Exit Recon


────────────────────────────────────────
2) CLI SHORTCUT MODE
────────────────────────────────────────
Run Recon directly from the terminal without entering
interactive mode.

General syntax:

  recon host <ip> [--txt] [--json]
  recon host -f <file> [--txt] [--json]

  recon port <profile> <ip> [--txt] [--json]
  recon port <profile> -f <file> [--txt] [--json]


────────────────────────────────────────
FILES
────────────────────────────────────────
  -f <file>
      Provide a file containing a list of IP addresses
      (one IP per line)


────────────────────────────────────────
OUTPUT OPTIONS
────────────────────────────────────────
  --txt
      Print formatted text output

  --json
      Print JSON output

Both --txt and --json can be used together.


────────────────────────────────────────
PORT PROFILES
────────────────────────────────────────
Standard and Deep variants are available:
  default
  common / deep
  ftp / ftp-deep
  ssh / ssh-deep
  smtp / smtp-deep
  dns / dns-deep
  web / web-deep
  kerberos / kerberos-deep
  snmp / snmp-deep
  ldap / ldap-deep
  smb / smb-deep
  mssql / mssql-deep
  mysql / mysql-deep
  postgresql / postgresql-deep
  rdp / rdp-deep
  vnc / vnc-deep
  winrm / winrm-deep

Vulnerability modes:
  vuln          Safe vulnerability checks
  vuln-deep     Aggressive checks (may include intrusive/DoS)

If no port profile is specified, the default mode will be used.


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
General syntax:

  recon host <ip> [--txt] [--json]
  recon host -f <file> [--txt] [--json]

  recon port <profile> <ip> [--txt] [--json]
  recon port <profile> -f <file> [--txt] [--json]

Use interactive shell mode:
  recon
`)
}

func PrintProfile() {
	fmt.Println(`
────────────────────────────────────────
PORT PROFILES
────────────────────────────────────────
Standard and Deep variants are available:
  default
  common / deep
  ftp / ftp-deep
  ssh / ssh-deep
  smtp / smtp-deep
  dns / dns-deep
  web / web-deep
  kerberos / kerberos-deep
  snmp / snmp-deep
  ldap / ldap-deep
  smb / smb-deep
  mssql / mssql-deep
  mysql / mysql-deep
  postgresql / postgresql-deep
  rdp / rdp-deep
  vnc / vnc-deep
  winrm / winrm-deep

Vulnerability modes:
  vuln          Safe vulnerability checks
  vuln-deep     Aggressive checks (may include intrusive/DoS)

If no port profile is specified, the default mode will be used.
`)
}
