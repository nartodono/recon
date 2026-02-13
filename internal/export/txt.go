package export

import (
	"fmt"
	"strings"
	"time"

	"github.com/nartodono/recon/internal/modules/host"
	"github.com/nartodono/recon/internal/modules/port"
)

// ---------- HOST TXT ----------

func HostSingleTXT(r host.Result, elapsedSeconds float64, t time.Time) string {
	var sb strings.Builder
	sb.WriteString("=== recon host ===\n")
	sb.WriteString(fmt.Sprintf("Time   : %s\n", t.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("Target : %s\n\n", r.Target))

	sb.WriteString(fmt.Sprintf("[%s] %s : %s\n", hostPrefix(r.Status), r.Target, r.Status))
	sb.WriteString(fmt.Sprintf("    Signal: PING=%s | NMAP_SN=%s\n", r.Ping, r.NmapSN))
	sb.WriteString(fmt.Sprintf("    Hint  : %s\n", r.Hint))
	sb.WriteString(fmt.Sprintf("    Time  : %.2fs\n", elapsedSeconds))
	sb.WriteString("\n")
	return sb.String()
}

func HostFileTXT(results []host.Result, summaryUp, summaryDown, summaryUnknown, total int, totalElapsedSeconds float64, t time.Time) string {
	var sb strings.Builder
	sb.WriteString("=== recon host ===\n")
	sb.WriteString(fmt.Sprintf("Time   : %s\n", t.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("Mode   : file\n\n"))

	for _, r := range results {
		sb.WriteString(fmt.Sprintf("[%s] %s : %s\n", hostPrefix(r.Status), r.Target, r.Status))
		sb.WriteString(fmt.Sprintf("    Signal: PING=%s | NMAP_SN=%s\n", r.Ping, r.NmapSN))
		sb.WriteString(fmt.Sprintf("    Hint  : %s\n\n", r.Hint))
	}

	sb.WriteString(fmt.Sprintf("Summary: UP=%d, DOWN=%d, UNKNOWN=%d, TOTAL=%d\n", summaryUp, summaryDown, summaryUnknown, total))
	sb.WriteString(fmt.Sprintf("Total Time: %.2fs\n\n", totalElapsedSeconds))
	return sb.String()
}

func hostPrefix(st host.FinalStatus) string {
	switch st {
	case host.StatusUP:
		return "+"
	case host.StatusDOWN:
		return "-"
	default:
		return "?"
	}
}

// ---------- PORT TXT ----------

func PortSingleTXT(r port.Result, profile string, elapsedSeconds float64, t time.Time) string {
	var sb strings.Builder
	sb.WriteString("=== recon port ===\n")
	sb.WriteString(fmt.Sprintf("Time    : %s\n", t.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("Target  : %s\n", r.Target))
	sb.WriteString(fmt.Sprintf("Profile : %s\n\n", profile))

	sb.WriteString(renderPortFindingsTXT(r.Findings))
	sb.WriteString(fmt.Sprintf("\nTime  : %.2fs\n\n", elapsedSeconds))
	return sb.String()
}

type PortFileItem struct {
	Target         string           `json:"target"`
	Findings       []port.PortFinding `json:"findings"`
	ElapsedSeconds float64          `json:"elapsed_seconds"`
}

func PortFileTXT(items []PortFileItem, profile string, totalElapsedSeconds float64, t time.Time) string {
	var sb strings.Builder
	sb.WriteString("=== recon port ===\n")
	sb.WriteString(fmt.Sprintf("Time    : %s\n", t.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("Mode    : file\n"))
	sb.WriteString(fmt.Sprintf("Profile : %s\n\n", profile))

	for _, it := range items {
		sb.WriteString("========================================\n")
		sb.WriteString(fmt.Sprintf("Target: %s\n\n", it.Target))
		sb.WriteString(renderPortFindingsTXT(it.Findings))
		sb.WriteString(fmt.Sprintf("\nTime  : %.2fs\n\n", it.ElapsedSeconds))
	}

	sb.WriteString(fmt.Sprintf("Total Time: %.2fs\n\n", totalElapsedSeconds))
	return sb.String()
}

func renderPortFindingsTXT(findings []port.PortFinding) string {
	var sb strings.Builder
	if len(findings) == 0 {
		sb.WriteString("[!] No ports found (or host did not respond).\n")
		return sb.String()
	}

	for _, f := range findings {
		prefix := "?"
		if f.State == "OPEN" {
			prefix = "+"
		} else if f.State == "CLOSED" {
			prefix = "-"
		}

		sb.WriteString(fmt.Sprintf("[%s] Port %d ---------------------------\n", prefix, f.Port))
		sb.WriteString(fmt.Sprintf("    %s - %s\n", f.Proto, f.Service))
		sb.WriteString(fmt.Sprintf("    Status : %s\n", f.State))
		if strings.TrimSpace(f.Version) != "" {
			sb.WriteString(fmt.Sprintf("    Version: %s\n", f.Version))
		}

		if len(f.Scripts) > 0 {
			sb.WriteString("\n")
			for _, s := range f.Scripts {
				sb.WriteString(fmt.Sprintf("    %s:\n", s.ID))
				for _, line := range strings.Split(s.Output, "\n") {
					line = strings.TrimRight(line, " \t")
					if line == "" {
						continue
					}
					sb.WriteString("      " + line + "\n")
				}
				sb.WriteString("\n")
			}
		} else {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
