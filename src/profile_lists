# Port Profiles

Recon uses profile-based Nmap arguments.

---

## General Profiles

| Profile | Description |
|----------|-------------|
| default | Standard scan (`-sC -sV`) |
| aggr | Aggressive scan (`-A`) |
| common | Top 1000 ports with light version detection |
| deep | OS detection + traceroute + safe/default/discovery scripts |

---

## Service Profiles

| Profile | Port(s) | Description |
|----------|---------|-------------|
| ftp | 21 | FTP enumeration |
| ftp-deep | 21 | Extended FTP safe scripts |
| ssh | 22 | SSH enumeration |
| ssh-deep | 22 | Extended SSH safe scripts |
| smtp | 25,587 | SMTP enumeration |
| smtp-deep | 25,587 | Extended SMTP safe scripts |
| dns | 53 | DNS enumeration |
| dns-deep | 53 | Extended DNS safe scripts |
| web | 80,443 | HTTP basic enumeration |
| web-deep | 80,443 | Extended HTTP safe scripts |
| kerberos | 88 | Kerberos enumeration |
| kerberos-deep | 88 | Extended Kerberos scripts |
| snmp | 161/UDP | SNMP enumeration |
| snmp-deep | 161/UDP | Extended SNMP safe scripts |
| ldap | 389 | LDAP enumeration |
| ldap-deep | 389 | Extended LDAP safe scripts |
| smb | 445 | SMB enumeration |
| smb-deep | 445 | Extended SMB safe scripts |
| mssql | 1433 | MSSQL enumeration |
| mssql-deep | 1433 | Extended MSSQL safe scripts |
| mysql | 3306 | MySQL enumeration |
| mysql-deep | 3306 | Extended MySQL safe scripts |
| postgresql | 5432 | PostgreSQL enumeration |
| postgresql-deep | 5432 | Extended PostgreSQL safe scripts |
| rdp | 3389 | RDP enumeration |
| rdp-deep | 3389 | Extended RDP safe scripts |
| vnc | 5900 | VNC enumeration |
| vnc-deep | 5900 | Extended VNC safe scripts |
| winrm | 5985,5986 | WinRM enumeration |
| winrm-deep | 5985,5986 | Extended WinRM safe scripts |

---

## Vulnerability Profiles

| Profile | Description |
|----------|-------------|
| vuln | Safe vulnerability scripts (excludes intrusive/dos/exploit) |
| vuln-deep | Includes intrusive/dos/exploit scripts (use with caution) |

---

⚠️ Deep profiles may take significantly longer per target.
⚠️ Vulnerability deep scans may cause service disruption.
