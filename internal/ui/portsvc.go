package ui

var portServices = map[string]struct{}{
	"default":      {},
	"aggr":         {},
	"common":       {},
	"deep":         {},
	"ftp":          {},
	"ftp-deep":     {},
	"ssh":          {},
	"ssh-deep":     {},
	"smtp":         {},
	"smtp-deep":    {},
	"dns":          {},
	"dns-deep":     {},
	"web":          {},
	"web-deep":     {},
	"kerberos":     {},
	"kerberos-deep":{},
	"snmp":         {},
	"snmp-deep":    {},
	"ldap":         {},
	"ldap-deep":    {},
	"smb":          {},
	"smb-deep":     {},
	"mssql":        {},
	"mssql-deep":   {},
	"mysql":        {},
	"mysql-deep":   {},
	"rdp":          {},
	"rdp-deep":     {},
	"postgres":     {},
	"postgres-deep":{},
	"vnc":          {},
	"vnc-deep":     {},
	"winrm":        {},
	"winrm-deep":   {},
	"vuln":         {},
	"vuln-deep":    {},
}

// portProfile maps a profile name to nmap arguments (excluding -Pn -oX - <target>).
var portProfile map[string]PortProfile {
	"default":      {Args: []string{
								"-sC",
								"-sV"
							}
					},
	"aggr":         {Args: []string{
								"-A",
								"--host-timeout", "10m",
								"--script-timeout", "90s",
								"--max-retries", "2",
								"-T4"
							}
					},
	"common":       {Args: []string{
								"-sV",
								"--top-ports", "1000",
								"--version-light",
								"--max-retries", "2",
								"-T4"
							}
					},
	"deep":         {Args: []string{
								"-sC",
								"-sV",
								"--script", "(default or safe or discovery) and not (dos or intrusive or exploit or brute)",
								"--script-timeout", "90s",
								"--host-timeout", "10m",
								"--max-retries", "2",
								"-T4"
							}
					},
	"ftp":          {Args: []string{
								"-sV",
								"--script", "ftp-anon,ftp-syst,ftp-bounce",
								"--script-timeout", "60s",
								"--host-timeout", "5m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "21"
					},
	"ftp-deep":     {Args: []string{
								"-sV",
								"--script", "(ftp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "8m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "21"
					},
	"ssh":          {Args: []string{
								"-sV",
								"--script", "ssh-hostkey,ssh2-enum-algos,ssh-auth-methods,banner",
								"--script-timeout", "60s",
								"--host-timeout", "5m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "22"
					},
	"ssh-deep":     {Args: []string{
								"-sV",
								"--script", "(ssh-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "8m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "22"
					},
	"smtp":         {Args: []string{
								"-sV",
								"--script", "smtp-commands,smtp-enum-users",
								"--script-timeout", "60s",
								"--host-timeout", "6m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "25,587"
					},
	"smtp-deep":    {Args: []string{
								"-sV",
								"--script", "(smtp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "8m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "25,587"
					},
	"dns":          {Args: []string{
								"-sV",
								"--script", "dns-nsid,dns-recursion",
								"--script-timeout", "45s",
								"--host-timeout", "4m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "53"
					},
	"dns-deep":     {Args: []string{
								"-sV",
								"--script", "(dns-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "60s",
								"--host-timeout", "6m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "53"
					},
	"web":          {Args: []string{
								"-sV",
								"--script", "http-title,http-headers,http-methods,http-enum,http-server-header",
								"--script-timeout", "60s",
								"--host-timeout", "6m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "80,443"
					},
	"web-deep":     {Args: []string{
								"-sV",
								"--script", "(http-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "9m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "80,443"
					},
	"kerberos":     {Args: []string{
								"-sV",
								"--script", "krb5-enum-users",
								"--script-timeout", "60s",
								"--host-timeout", "5m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "88"
						},
	"kerberos-deep":{Args: []string{
								"-sV",
								"--script", "(krb5-* and (safe or default or discovery)) and not (brute or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "8m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "88"
					},
	"snmp":         {Args: []string{
								"-sU", "-sV",
								"--script", "snmp-info,snmp-sysdescr,snmp-interfaces",
								"--script-timeout", "45s",
								"--host-timeout", "4m",
								"--max-retries", "1",
								"-T4"
							},	DefaultPorts: "161"
					},
	"snmp-deep":    {Args: []string{
								"-sU","-sV",
								"--script", "(snmp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "60s",
								"--host-timeout", "5m",
								"--max-retries", "1",
								"-T4"
							},	DefaultPorts: "161"
					},
	"ldap":         {Args: []string{
								"-sV",
								"--script", "ldap-rootdse,ldap-search",
								"--script-timeout", "60s",
								"--host-timeout", "6m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "389"
					},
	"ldap-deep":    {Args: []string{
								"-sV",
								"--script", "(ldap-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "9m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "389"
					},
	"smb":          {Args: []string{
								"-sV",
								"--script", "smb-os-discovery,smb2-security-mode,smb2-time,smb-protocols",
								"--script-timeout", "60s",
								"--host-timeout", "6m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "445"
					},
	"smb-deep":     {Args: []string{
								"-sV",
								"--script", "(smb-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "10m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "445"
					},
	"mssql":        {Args: []string{
								"-sV",
								"--script", "ms-sql-info,ms-sql-ntlm-info",
								"--script-timeout", "60s",
								"--host-timeout", "6m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "1433"
					},
	"mssql-deep":   {Args: []string{
								"sV",
								"--script", "(ms-sql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "9m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "1433"
					},
	"mysql":        {Args: []string{
								"-sV",
								"--script", "mysql-info,mysql-capabilities",
								"--script-timeout", "60s",
								"--host-timeout", "6m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "3306"
					},
	"mysql-deep":   {Args: []string{
								"-sV",
								"--script", "(mysql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "9m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "3306"
					},
	"rdp":          {Args: []string{
								"-sV",
								"--script", "rdp-ntlm-info,rdp-enum-encryption",
								"--script-timeout", "60s",
								"--host-timeout", "6m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "3389"
					},
	"rdp-deep":     {Args: []string{
								"-sV",
								"--script", "(rdp-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "9m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "3389"
					},
	"postgresql":	{Args: []string{
								"-sV",
								"--script", "pgsql-info",
								"--script-timeout", "60s",
								"--host-timeout", "6m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "5432"
					},
	"postgresql-deep":{Args: []string{
								"-sV",
								"--script", "(pgsql-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "9m",
								"--max-retries", "2",
								"-T4",
							},	DefaultPorts: "5432"
					},
	"vnc":          {Args: []string{
								"-sV",
								"--script", "vnc-info",
								"--script-timeout", "60s",
								"--host-timeout", "6m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "5900"
					},
	"vnc-deep":     {Args: []string{
								"-sV",
								"--script", "(vnc-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "9m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "5900"
					},
	"winrm":        {Args: []string{
								"-sV",
								"--script", "wsman-info",
								"--script-timeout", "60s",
								"--host-timeout", "6m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "5985,5986"
					},
	"winrm-deep":   {Args: []string{
								"-sV",
								"--script", "(wsman-* and (safe or default or discovery)) and not (brute or intrusive or dos or exploit)",
								"--script-timeout", "90s",
								"--host-timeout", "9m",
								"--max-retries", "2",
								"-T4"
							},	DefaultPorts: "5985,5986"
					},
	"vuln":         {Args: []string{
								"-sV",
								"--version-light",
								"--script", "vuln",
								"--host-timeout", "20m",
								"--max-retries", "2",
								"-T4"
							}
					},
	"vuln-deep":    {Args: []string{
								"-sV",
								"--version-light",
								"--script", "vuln or exploit",
								"--script-timeout", "3m",
								"--host-timeout", "30m",
								"--max-retries", "2",
								"-T4"
							}
					},
}

func isPortService(s string) bool {
	_, ok := portServices[s]
	return ok
}
