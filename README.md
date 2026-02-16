# Recon

Recon is a lightweight CLI tool that runs **profile-based Nmap recon** and produces **clean, structured output** (TXT + JSON).  
It supports both **interactive shell mode** and **CLI shortcut mode**, with built-in safety limits for multi-target scanning.

> ⚠️ Use responsibly. Only scan systems you own or have explicit permission to test.


## Interactive Mode

<p align="center">
  <img src="src/img/recon_sh.png" width="50%">
</p>

## CLI Help

<p align="center">
  <img src="src/img/recon_h.png" width="50%">
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
OR
```bash
GOPROXY=direct GOSUMDB=off go install github.com/nartodono/recon/cmd/recon@main
```

### using `git clone`
```bash
git clone https://github.com/nartodono/recon.git
cd recon
go build ./cmd/recon
```

