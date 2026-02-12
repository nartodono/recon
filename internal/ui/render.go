package ui

import (
	"fmt"
	"recon/internal/modules/host"
)

func RenderHostResult(r host.Result) {
	var prefix string
	var statusText string

	switch r.Status {
	case host.StatusUP:
		prefix = ColorGreen + "[+]" + ColorReset
		statusText = ColorGreen + string(r.Status) + ColorReset
	case host.StatusDOWN:
		prefix = ColorRed + "[-]" + ColorReset
		statusText = ColorRed + string(r.Status) + ColorReset
	default:
		prefix = ColorYellow + "[?]" + ColorReset
		statusText = ColorYellow + string(r.Status) + ColorReset
	}

	fmt.Printf("%s %s : %s\n", prefix, r.Target, statusText)
	fmt.Printf("    Signal: PING=%s | NMAP_SN=%s\n", r.Ping, r.NmapSN)
	fmt.Printf("    Hint  : %s\n\n", r.Hint)
}
