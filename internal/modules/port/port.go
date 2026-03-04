package port

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/nartodono/recon/internal/system"
)

type PortFinding struct {
	Port    int
	Proto   string
	State   string
	Service string
	Version string
	Scripts []ScriptFinding
}

type ScriptFinding struct {
	ID     string
	Output string
}

type Result struct {
	Target   string
	Findings []PortFinding
	Warning string

	HostUp      bool
	LatencySec  float64
	NotShown    string
	ServiceInfo string
}

func runCmd(name string, args ...string) (stdout string, stderr string, err error) {
	cmd := exec.Command(name, args...)

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err = cmd.Run()
	return outBuf.String(), errBuf.String(), err
}

func Scan(target string, extraArgs []string) (Result, error) {
	if _, err := exec.LookPath("nmap"); err != nil {
		return Result{}, fmt.Errorf("nmap not found. Install: sudo apt install nmap")
	}

	if err := system.ValidateResolvable(target); err != nil {
		return Result{}, err
	}

	baseArgs := []string{"-Pn", "-oX", "-"}
	baseArgs = append(baseArgs, extraArgs...)
	baseArgs = append(baseArgs, target)

	stdout, stderr, err := runCmd("nmap", baseArgs...)

	if err != nil && !strings.Contains(strings.ToLower(stdout), "<nmaprun") {
		combined := strings.TrimSpace(stdout + "\n" + stderr)
		return Result{}, fmt.Errorf("nmap error: %v\n%s", err, combined)
	}

	var run NmapRun
	if e := xml.Unmarshal([]byte(stdout), &run); e != nil {
		return Result{}, fmt.Errorf("failed to parse nmap XML: %w\n%s", e, stderr)
	}

	warning := ""
	exit := strings.ToLower(strings.TrimSpace(run.RunStats.Finished.Exit))
	if err != nil || (exit != "" && exit != "success") {
		if exit == "" {
			exit = "unknown"
		}
		warning = fmt.Sprintf("Scan may be incomplete (nmap exit=%s). Output may be partial.", exit)
	}

	if len(run.Hosts) == 0 {
		return Result{Target: target, Findings: nil, Warning: warning}, nil
	}

	h := run.Hosts[0]

	hostUp := strings.EqualFold(strings.TrimSpace(h.Status.State), "up")
	latencySec := 0.0
	if s := strings.TrimSpace(h.Times.SRTT); s != "" {
		if us, e := strconv.ParseFloat(s, 64); e == nil {
			latencySec = us / 1_000_000.0
		}
	}

	closedCount, filteredCount := 0, 0
	closedReason, filteredReason := "", ""
	for _, ep := range h.Ports.ExtraPorts {
		st := strings.ToLower(strings.TrimSpace(ep.State))
		reason := strings.TrimSpace(ep.Reason)
		if reason == "" && len(ep.ExtraReasons) > 0 {
			reason = strings.TrimSpace(ep.ExtraReasons[0].Reason)
		}
		if st == "closed" {
			closedCount += ep.Count
			if closedReason == "" && reason != "" {
				closedReason = reason
			}
		} else if st == "filtered" {
			filteredCount += ep.Count
			if filteredReason == "" && reason != "" {
				filteredReason = reason
			}
		}
	}

	notShown := ""
	if closedCount > 0 {
		reason := closedReason
		if reason == "" {
			reason = "reset"
		}
		notShown = fmt.Sprintf("Not shown: %d closed tcp ports (%s)", closedCount, reason)
	}

	if hostUp && filteredCount > 0 && len(h.Ports.Port) == 0 {
		reason := filteredReason
		if reason == "" {
			reason = "no-response"
		}
		warning = fmt.Sprintf(
			"All %d scanned ports on %s are in ignored states.\nNot shown: %d filtered tcp ports (%s)",
			filteredCount,
			target,
			filteredCount,
			reason,
		)
		notShown = ""
	}

	cpeSet := map[string]struct{}{}

	findings := make([]PortFinding, 0, len(h.Ports.Port))
	for _, p := range h.Ports.Port {
		for _, c := range p.Service.CPEs {
			c = strings.TrimSpace(c)
			if c != "" {
				cpeSet[c] = struct{}{}
			}
		}

		svc := p.Service.Name
		if p.Service.Tunnel != "" && p.Service.Name != "" {
			svc = p.Service.Tunnel + "/" + p.Service.Name
		}

		verParts := []string{}
		if p.Service.Product != "" {
			verParts = append(verParts, p.Service.Product)
		}
		if p.Service.Version != "" {
			verParts = append(verParts, p.Service.Version)
		}
		if p.Service.Extra != "" {
			verParts = append(verParts, "("+p.Service.Extra+")")
		}
		version := strings.TrimSpace(strings.Join(verParts, " "))

		sf := []ScriptFinding{}
		for _, s := range p.Scripts {
			sf = append(sf, ScriptFinding{ID: s.ID, Output: s.Output})
		}

		findings = append(findings, PortFinding{
			Port:    p.PortID,
			Proto:   strings.ToUpper(p.Protocol),
			State:   strings.ToUpper(p.State.State),
			Service: strings.ToUpper(svc),
			Version: version,
			Scripts: sf,
		})
	}

	sort.Slice(findings, func(i, j int) bool {
		return findings[i].Port < findings[j].Port
	})

	serviceInfo := ""
	if len(cpeSet) > 0 {
		cpes := make([]string, 0, len(cpeSet))
		for c := range cpeSet {
			cpes = append(cpes, c)
		}
		sort.Strings(cpes)

		osName := ""
		for _, c := range cpes {
			if strings.HasPrefix(c, "cpe:/o:linux") {
				osName = "Linux"
				break
			}
			if strings.HasPrefix(c, "cpe:/o:microsoft") {
				osName = "Windows"
				break
			}
		}

		if osName != "" {
			serviceInfo = fmt.Sprintf("Service Info: OS: %s; CPE: %s", osName, strings.Join(cpes, ", "))
		} else {
			serviceInfo = fmt.Sprintf("Service Info: CPE: %s", strings.Join(cpes, ", "))
		}
	}

	return Result{
		Target:      target,
		Findings:    findings,
		Warning:     warning,
		HostUp:      hostUp,
		LatencySec:  latencySec,
		NotShown:    notShown,
		ServiceInfo: serviceInfo,
	}, nil
}
