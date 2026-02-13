package port

import "encoding/xml"

type NmapRun struct {
	XMLName xml.Name `xml:"nmaprun"`
	Hosts   []Host   `xml:"host"`
}

type Host struct {
	Addresses []Address `xml:"address"`
	Ports     Ports     `xml:"ports"`
}

type Address struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"` // ipv4, ipv6, mac
}

type Ports struct {
	Port []Port `xml:"port"`
}

type Port struct {
	Protocol string `xml:"protocol,attr"` // tcp/udp
	PortID   int    `xml:"portid,attr"`
	State    State  `xml:"state"`
	Service  Service `xml:"service"`
	Scripts  []Script `xml:"script"`
}

type State struct {
	State string `xml:"state,attr"` // open/closed/filtered
	Reason string `xml:"reason,attr,omitempty"`
}

type Service struct {
	Name    string `xml:"name,attr,omitempty"`
	Product string `xml:"product,attr,omitempty"`
	Version string `xml:"version,attr,omitempty"`
	Extra   string `xml:"extrainfo,attr,omitempty"`
	Tunnel  string `xml:"tunnel,attr,omitempty"` // e.g., ssl
}

type Script struct {
	ID     string `xml:"id,attr"`
	Output string `xml:"output,attr"`
}
