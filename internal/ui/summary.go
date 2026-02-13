package ui

import (
	"fmt"
	"github.com/nartodono/recon/internal/modules/host"
)

type HostCounts struct {
	Up      int
	Down    int
	Unknown int
	Total   int
}

func CountHostStatus(r host.Result, c *HostCounts) {
	c.Total++
	switch r.Status {
	case host.StatusUP:
		c.Up++
	case host.StatusDOWN:
		c.Down++
	default:
		c.Unknown++
	}
}

func PrintHostSummary(c HostCounts) {
	fmt.Printf("Summary: UP=%s, DOWN=%s, UNKNOWN=%s, TOTAL=%d\n\n",
		Green(fmt.Sprintf("%d", c.Up)),
		Red(fmt.Sprintf("%d", c.Down)),
		Yellow(fmt.Sprintf("%d", c.Unknown)),
		c.Total,
	)
}
