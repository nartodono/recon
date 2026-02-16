# Recon

Recon is a lightweight CLI tool that runs **profile-based Nmap recon** and produces **clean, structured output** (TXT + JSON).  
It supports both **interactive shell mode** and **CLI shortcut mode**, with built-in safety limits for multi-target scanning.

> ⚠️ Use responsibly. Only scan systems you own or have explicit permission to test.

---

## Interactive Mode

<p align="center">
  <img src="src/img/recon_sh.png" width="55%">
</p>

---

## CLI Help

<p align="center">
  <img src="src/img/recon_h.png" width="55%">
</p>

---

## Features

- **Two modes**
  - Interactive shell: run commands inside `recon >`
  - CLI shortcut: run directly from terminal
- **Port profiles** (default/common/deep + service-specific profiles)
- **Structured output**
  - `--txt` pretty human-readable format
  - `--json` machine-readable results
- **Auto-saves results** to `~/recon_result/` with timestamped filenames
- **Progress & warnings** (e.g., deep profile in file mode)
- **Multi-target support** via `-f <file>` (one IP per line)

---

## Requirements

- Go (for install/build)
- `nmap` installed and accessible in PATH
- Standard utilities like `ping` (for host check)

> Some scan types (e.g., OS detection `-O`, traceroute) may require elevated privileges depending on your OS/environment.

---

## Installation

### Using `go install`

```bash
go install github.com/nartodono/recon/cmd/recon@latest
```
If module resolution issues occur:
```bash
GOPROXY=direct GOSUMDB=off go install github.com/nartodono/recon/cmd/recon@main
```

### Using `git clone`
```bash
git clone https://github.com/nartodono/recon.git
cd recon
go build ./cmd/recon
```

## Tools Usage

For a complete list of available port profiles, see:
[`Port Scanning Profile Lists`](src/profile_lists.txt)

INTERACTIVE MODE
----------------
Start Recon without arguments:
```bash
  recon
```
Inside the shell:
```bash
  recon > host 192.168.1.1 --txt --json
  recon > port 192.168.1.1
  recon > port web-deep 192.168.1.1 --txt
  recon > port vuln 192.168.1.1
  recon > profile
  recon > exit
```

CLI SHORTCUT MODE
-----------------
Run directly from terminal:
```bash
  recon host 192.168.1.1 --txt --json
  recon port smb 192.168.1.20 --txt
  recon port web-deep 192.168.1.20 --txt --json
  recon port vuln-deep 192.168.1.1
```
If no profile is specified:
```bash
  recon port 192.168.1.1
```
The 'default' profile will be used automatically.


FILE MODE
---------
Scan multiple targets from file (one IP per line):
```bash
  recon host -f targets.txt --txt --json
  recon port -f targets.txt --txt
  recon port deep -f targets.txt --txt --json
```

OUTPUT OPTIONS
--------------
  `--txt`   Print formatted text output
  `--json`  Print structured JSON output

Both flags can be used together.


NOTES
-----
- Results are automatically saved to:
    ~/recon_result/

- Output filenames follow:
    `recon-host-YYYYMMDD-HHMMSS.txt`
    `recon-host-YYYYMMDD-HHMMSS.json`
    `recon-port-YYYYMMDD-HHMMSS.txt
    `recon-port-YYYYMMDD-HHMMSS.json

- Multi-target limits:
    Normal profiles  → max 30 targets
    Deep profiles    → max 10 targets

- Recon checks required dependencies on startup:
    nmap
    ping

## Example

### Vulnerability Scan
<p align="center"> <img src="src/img/recon_port2.png" width="55%"> </p>

### Host Scan
<p align="center"> <img src="src/img/recon_host1.png" width="55%"> </p>

### Port Scan
<p align="center"> <img src="src/img/recon_port1.png" width="55%"> </p>

### Saved TXT Output
<p align="center"> <img src="src/img/recon_port_txt.png" width="55%"> </p>

### Saved JSON Output
<p align="center"> <img src="src/img/recon_host_json.png" width="50%"> </p>
