# Port Profiles

Recon uses base Nmap arguments:

  -Pn -oX -

Each profile appends additional flags as defined below.

---

## General Profiles

| Profile | Nmap Flags | Description |
|----------|------------|-------------|
| default | `-sC -sV` | Runs default NSE scripts (`-sC`) and performs service version detection (`-sV`). Balanced baseline scan. |
| aggr | `-A --host-timeout 10m --script-timeout 90s --max-retries 2 -T4` | Aggressive scan with OS detection, version detection, default scripts, and traceroute. Execution time is capped to prevent excessive delays. |
| common | `-sV --top-ports 1000 --version-light --max-retries 2 -T4` | Scans top 1000 most common ports with lightweight version detection. Faster and optimized for broader coverage. |
| deep | `-sC -sV --script "(default or safe or discovery) and not (dos or intrusive or exploit or brute)" --script-timeout 90s --host-timeout 10m --max-retries 2 -T4` | Extended discovery scan focused on safe/default/discovery scripts. Intrusive, brute-force, exploit, and DoS categories are excluded. Execution time is limited for stability and predictability. |
| vuln | `-sV --version-light --script vuln` | Basic vulnerability scan using NSE vuln category with lightweight version detection. Designed for balanced runtime and useful findings. |
| vuln-deep | `-sV --version-light --script "vuln or exploit" --script-timeout 3m --host-timeout 30m --max-retries 2 -T4` | Extended vulnerability scan including exploit-category scripts. Allows longer execution time while still enforcing runtime limits. |

---

## Service Profiles

Each service profile focuses on specific ports and script sets.

| Profile | Port(s) | Nmap Flags | Description |
|----------|---------|------------|-------------|
| ftp | 21 | `-p 21 -sV --script ftp-anon,ftp-syst,ftp-bounce --script-timeout 60s --host-timeout 5m --max-retries 2 -T4` | Checks anonymous login, FTP system info, and bounce capability. |
| ftp-deep | 21 | `-p 21 -sV --script "(ftp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 8m --max-retries 2 -T4` | Runs extended safe FTP-related scripts with execution guardrails. |
| ssh | 22 | `-p 22 -sV --script ssh-hostkey,ssh2-enum-algos,ssh-auth-methods,banner --script-timeout 60s --host-timeout 5m --max-retries 2 -T4` | Enumerates SSH host keys, supported algorithms, auth methods, and banner. |
| ssh-deep | 22 | `-p 22 -sV --script "(ssh-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 8m --max-retries 2 -T4` | Extended safe SSH scripts with timeout limits. |
| smtp | 25,587 | `-p 25,587 -sV --script smtp-commands,smtp-enum-users --script-timeout 60s --host-timeout 6m --max-retries 2 -T4` | Enumerates SMTP commands and possible users. |
| smtp-deep | 25,587 | `-p 25,587 -sV --script "(smtp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 8m --max-retries 2 -T4` | Extended safe SMTP scripts with guardrails. |
| dns | 53 | `-p 53 -sV --script dns-nsid,dns-recursion --script-timeout 45s --host-timeout 4m --max-retries 2 -T4` | Checks DNS server ID and recursion capability. |
| dns-deep | 53 | `-p 53 -sV --script "(dns-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 60s --host-timeout 6m --max-retries 2 -T4` | Extended safe DNS scripts with timeout limits. |
| web | 80,443 | `-p 80,443 -sV --script http-title,http-headers,http-methods,http-enum,http-server-header --script-timeout 60s --host-timeout 6m --max-retries 2 -T4` | Basic HTTP enumeration (title, headers, methods, content discovery). |
| web-deep | 80,443 | `-p 80,443 -sV --script "(http-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 9m --max-retries 2 -T4` | Extended safe HTTP scripts with controlled execution time. |
| kerberos | 88 | `-p 88 -sV --script krb5-enum-users --script-timeout 60s --host-timeout 5m --max-retries 2 -T4` | Attempts Kerberos user enumeration. |
| kerberos-deep | 88 | `-p 88 -sV --script "(krb5-* and (safe or default or discovery)) and not (brute or dos or exploit)" --script-timeout 90s --host-timeout 8m --max-retries 2 -T4` | Extended Kerberos scripts excluding intrusive/exploit. |
| snmp | 161/UDP | `-sU -p 161 -sV --script snmp-info,snmp-sysdescr,snmp-interfaces --script-timeout 45s --host-timeout 4m --max-retries 1 -T4` | SNMP info, system description, and interface enumeration. |
| snmp-deep | 161/UDP | `-sU -p 161 -sV --script "(snmp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 60s --host-timeout 5m --max-retries 1 -T4` | Extended safe SNMP scripts with tighter retry control (UDP optimized). |
| ldap | 389 | `-p 389 -sV --script ldap-rootdse,ldap-search --script-timeout 60s --host-timeout 6m --max-retries 2 -T4` | LDAP rootDSE query and basic search. |
| ldap-deep | 389 | `-p 389 -sV --script "(ldap-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 9m --max-retries 2 -T4` | Extended safe LDAP scripts with execution limits. |
| smb | 445 | `-p 445 -sV --script smb-os-discovery,smb2-security-mode,smb2-time,smb-protocols --script-timeout 60s --host-timeout 6m --max-retries 2 -T4` | SMB OS discovery, protocol and security mode enumeration. |
| smb-deep | 445 | `-p 445 -sV --script "(smb-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 10m --max-retries 2 -T4` | Extended safe SMB scripts with capped execution time. |
| mssql | 1433 | `-p 1433 -sV --script ms-sql-info,ms-sql-ntlm-info --script-timeout 60s --host-timeout 6m --max-retries 2 -T4` | MSSQL instance and NTLM info enumeration. |
| mssql-deep | 1433 | `-p 1433 -sV --script "(ms-sql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 9m --max-retries 2 -T4` | Extended safe MSSQL scripts with timeout control. |
| mysql | 3306 | `-p 3306 -sV --script mysql-info,mysql-capabilities --script-timeout 60s --host-timeout 6m --max-retries 2 -T4` | MySQL server info and capability enumeration. |
| mysql-deep | 3306 | `-p 3306 -sV --script "(mysql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 9m --max-retries 2 -T4` | Extended safe MySQL scripts with timeout limits. |
| postgresql | 5432 | `-p 5432 -sV --script pgsql-info --script-timeout 60s --host-timeout 6m --max-retries 2 -T4` | PostgreSQL info enumeration. |
| postgresql-deep | 5432 | `-p 5432 -sV --script "(pgsql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 9m --max-retries 2 -T4` | Extended safe PostgreSQL scripts with timeout control. |
| rdp | 3389 | `-p 3389 -sV --script rdp-ntlm-info,rdp-enum-encryption --script-timeout 60s --host-timeout 6m --max-retries 2 -T4` | RDP NTLM info and encryption enumeration. |
| rdp-deep | 3389 | `-p 3389 -sV --script "(rdp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 9m --max-retries 2 -T4` | Extended safe RDP scripts with execution guardrails. |
| vnc | 5900 | `-p 5900 -sV --script vnc-info --script-timeout 60s --host-timeout 6m --max-retries 2 -T4` | VNC server information gathering. |
| vnc-deep | 5900 | `-p 5900 -sV --script "(vnc-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 9m --max-retries 2 -T4` | Extended safe VNC scripts with timeout limits. |
| winrm | 5985,5986 | `-p 5985,5986 -sV --script wsman-info --script-timeout 60s --host-timeout 6m --max-retries 2 -T4` | WinRM service information gathering. |
| winrm-deep | 5985,5986 | `-p 5985,5986 -sV --script "(wsman-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)" --script-timeout 90s --host-timeout 9m --max-retries 2 -T4` | Extended safe WinRM scripts with execution guardrails. |

---

## Vulnerability Profiles

| Profile | Nmap Flags | Description |
|----------|------------|-------------|
| vuln | `-sV --script "vuln and not (dos or intrusive or exploit)"` | Runs vulnerability scripts excluding intrusive, exploit, and DoS categories. |
| vuln-deep | `-sV --script "(vuln or dos or intrusive or exploit)"` | Includes intrusive, exploit, and DoS scripts. Use with caution. |

---

⚠️ Deep profiles may take significantly longer per target.  
⚠️ `vuln-deep` may cause service disruption. Use only with proper authorization.
