package port

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os/exec"
	"sort"
	"strings"
	"time"
	"github.com/nartodono/recon/internal/system"
)

type PortFinding struct {
	Port     int
	Proto    string
	State    string
	Service  string
	Version  string
	Scripts  []ScriptFinding
}

type ScriptFinding struct {
	ID     string
	Output string
}

type Result struct {
	Target   string
	Findings []PortFinding
}

func runCmd(name string, args ...string) (stdout string, stderr string, err error) {
	cmd := exec.Command(name, args...)

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err = cmd.Run()

	return outBuf.String(), errBuf.String(), err
}

func Scan(target string, extraArgs []string) (Result, error)
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
	
	if len(run.Hosts) == 0 {
		return Result{Target: target, Findings: nil}, nil
	}

	h := run.Hosts[0]

	findings := make([]PortFinding, 0, len(h.Ports.Port))
	for _, p := range h.Ports.Port {
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

	return Result{Target: target, Findings: findings}, nil
}
