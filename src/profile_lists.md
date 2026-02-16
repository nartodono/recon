# Port Profiles

Recon uses base Nmap arguments:

  -Pn -oX -

Each profile appends additional flags as defined below.

---

## General Profiles

| Profile | Nmap Flags | Description |
|----------|------------|-------------|
| default | `-sC -sV` | Runs default NSE scripts (`-sC`) and performs service version detection (`-sV`). |
| aggr | `-A` | Aggressive scan: enables OS detection, version detection, default scripts, and traceroute. |
| common | `-sV --top-ports 1000 --version-light` | Scans top 1000 most common ports with lightweight version detection. Faster than default. |
| deep | `-sC -sV -O --traceroute --script "(default or safe or discovery) and not (dos or intrusive or exploit)"` | Extended discovery: includes OS detection, traceroute, and runs safe/default/discovery NSE scripts while excluding intrusive, exploit, brute, and DoS categories. |

---

## Service Profiles

Each service profile focuses on specific ports and script sets.

| Profile | Port(s) | Nmap Flags | Description |
|----------|---------|------------|-------------|
| ftp | 21 | `-p 21 -sV --script=ftp-anon,ftp-syst,ftp-bounce` | Checks anonymous login, FTP system info, and bounce capability. |
| ftp-deep | 21 | `-p 21 -sV --script "(ftp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Runs extended safe FTP-related scripts. |
| ssh | 22 | `-p 22 -sV --script=ssh-hostkey,ssh2-enum-algos,ssh-auth-methods,banner` | Enumerates SSH host keys, supported algorithms, auth methods, and banner. |
| ssh-deep | 22 | `-p 22 -sV --script "(ssh-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Runs extended safe SSH scripts. |
| smtp | 25,587 | `-p 25,587 -sV --script=smtp-commands,smtp-enum-users` | Enumerates SMTP commands and possible users. |
| smtp-deep | 25,587 | `-p 25,587 -sV --script "(smtp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe SMTP scripts. |
| dns | 53 | `-p 53 -sV --script=dns-nsid,dns-recursion` | Checks DNS server ID and recursion capability. |
| dns-deep | 53 | `-p 53 -sV --script "(dns-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe DNS scripts. |
| web | 80,443 | `-p 80,443 -sV --script=http-title,http-headers,http-methods,http-enum,http-server-header` | Basic HTTP enumeration (title, headers, methods, content discovery). |
| web-deep | 80,443 | `-p 80,443 -sV --script "(http-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe HTTP scripts. |
| kerberos | 88 | `-p 88 -sV --script=krb5-enum-users` | Attempts Kerberos user enumeration. |
| kerberos-deep | 88 | `-p 88 -sV --script "(krb5-* and not (brute or dos or exploit))"` | Extended Kerberos scripts excluding intrusive/exploit. |
| snmp | 161/UDP | `-sU -p 161 -sV --script=snmp-info,snmp-sysdescr,snmp-interfaces` | SNMP info, system description, and interface enumeration. |
| snmp-deep | 161/UDP | `-sU -p 161 -sV --script "(snmp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe SNMP scripts. |
| ldap | 389 | `-p 389 -sV --script=ldap-rootdse,ldap-search` | LDAP rootDSE query and basic search. |
| ldap-deep | 389 | `-p 389 -sV --script "(ldap-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe LDAP scripts. |
| smb | 445 | `-p 445 -sV --script=smb-os-discovery,smb2-security-mode,smb2-time,smb-protocols` | SMB OS discovery, protocol and security mode enumeration. |
| smb-deep | 445 | `-p 445 -sV --script "(smb-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe SMB scripts. |
| mssql | 1433 | `-p 1433 -sV --script=ms-sql-info,ms-sql-ntlm-info` | MSSQL instance and NTLM info enumeration. |
| mssql-deep | 1433 | `-p 1433 -sV --script "(ms-sql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe MSSQL scripts. |
| mysql | 3306 | `-p 3306 -sV --script=mysql-info,mysql-capabilities` | MySQL server info and capability enumeration. |
| mysql-deep | 3306 | `-p 3306 -sV --script "(mysql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe MySQL scripts. |
| postgresql | 5432 | `-p 5432 -sV --script=pgsql-info` | PostgreSQL info enumeration. |
| postgresql-deep | 5432 | `-p 5432 -sV --script "(pgsql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe PostgreSQL scripts. |
| rdp | 3389 | `-p 3389 -sV --script=rdp-ntlm-info,rdp-enum-encryption` | RDP NTLM info and encryption enumeration. |
| rdp-deep | 3389 | `-p 3389 -sV --script "(rdp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe RDP scripts. |
| vnc | 5900 | `-p 5900 -sV --script=vnc-info` | VNC server information gathering. |
| vnc-deep | 5900 | `-p 5900 -sV --script "(vnc-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe VNC scripts. |
| winrm | 5985,5986 | `-p 5985,5986 -sV --script=wsman-info` | WinRM service information. |
| winrm-deep | 5985,5986 | `-p 5985,5986 -sV --script "(wsman-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)"` | Extended safe WinRM scripts. |

---

## Vulnerability Profiles

| Profile | Nmap Flags | Description |
|----------|------------|-------------|
| vuln | `-sV --script "vuln and not (dos or intrusive or exploit)"` | Runs vulnerability scripts excluding intrusive, exploit, and DoS categories. |
| vuln-deep | `-sV --script "(vuln or dos or intrusive or exploit)"` | Includes intrusive, exploit, and DoS scripts. Use with caution. |

---

⚠️ Deep profiles may take significantly longer per target.  
⚠️ `vuln-deep` may cause service disruption. Use only with proper authorization.
