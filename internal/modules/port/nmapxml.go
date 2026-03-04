package port

import "encoding/xml"

type NmapRun struct {
	XMLName   xml.Name `xml:"nmaprun"`
	Hosts     []Host   `xml:"host"`
	RunStats  RunStats `xml:"runstats"`
}

type RunStats struct {
	Finished Finished `xml:"finished"`
}

type Finished struct {
	Exit string `xml:"exit,attr"`
}

type Host struct {
	Status    HostStatus `xml:"status"`
	Addresses []Address  `xml:"address"`
	Ports     Ports      `xml:"ports"`
}

type HostStatus struct {
	State  string `xml:"state,attr"`
	Reason string `xml:"reason,attr,omitempty"`
}

type Address struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
}

type Ports struct {
	Port       []Port      `xml:"port"`
	ExtraPorts []ExtraPort `xml:"extraports"`
}

type ExtraPort struct {
	State  string `xml:"state,attr"`
	Count  int    `xml:"count,attr"`
	Reason string `xml:"reason,attr,omitempty"`
}

type Port struct {
	Protocol string   `xml:"protocol,attr"`
	PortID   int      `xml:"portid,attr"`
	State    State    `xml:"state"`
	Service  Service  `xml:"service"`
	Scripts  []Script `xml:"script"`
}

type State struct {
	State  string `xml:"state,attr"`
	Reason string `xml:"reason,attr,omitempty"`
}

type Service struct {
	Name    string `xml:"name,attr,omitempty"`
	Product string `xml:"product,attr,omitempty"`
	Version string `xml:"version,attr,omitempty"`
	Extra   string `xml:"extrainfo,attr,omitempty"`
	Tunnel  string `xml:"tunnel,attr,omitempty"`
}

type Script struct {
	ID     string `xml:"id,attr"`
	Output string `xml:"output,attr"`
}
