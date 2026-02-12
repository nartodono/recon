package host

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
	"recon/internal/system"
)

type PingSignal string
type NmapSignal string
type FinalStatus string

const (
	PingOK          PingSignal = "OK"
	PingRTO         PingSignal = "RTO"
	PingUnreachable PingSignal = "HOST_UNREACHABLE"
	PingUnknown     PingSignal = "UNKNOWN"
)

const (
	NmapUp        NmapSignal = "HOST_UP"
	NmapDown      NmapSignal = "HOST_DOWN"
	NmapNoConfirm NmapSignal = "NO_CONFIRM"
	NmapError     NmapSignal = "ERROR"
)

const (
	StatusUP      FinalStatus = "UP"
	StatusDOWN    FinalStatus = "DOWN"
	StatusUNKNOWN FinalStatus = "UNKNOWN OR FILTERED"
)

type Result struct {
	Target string
	Ping   PingSignal
	NmapSN NmapSignal
	Status FinalStatus
	Hint   string
}

func runCmd(timeout time.Duration, name string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	err := cmd.Run()
	out := buf.String()

	if ctx.Err() == context.DeadlineExceeded {
		return out, fmt.Errorf("%s timed out", name)
	}
	return out, err
}

func pingCheck(target string) PingSignal {
	out, _ := runCmd(4*time.Second, "ping", "-c", "1", "-W", "1", target)
	l := strings.ToLower(out)

	if strings.Contains(l, "1 received") || strings.Contains(l, "bytes from") {
		return PingOK
	}
	if strings.Contains(l, "destination host unreachable") ||
		strings.Contains(l, "host unreachable") ||
		strings.Contains(l, "network is unreachable") {
		return PingUnreachable
	}
	if strings.Contains(l, "100% packet loss") ||
		strings.Contains(l, "request timeout") ||
		strings.Contains(l, "time out") {
		return PingRTO
	}
	return PingUnknown
}

func nmapSnCheck(target string) NmapSignal {
	out, err := runCmd(20*time.Second, "nmap", "-sn", target)
	l := strings.ToLower(out)

	if strings.Contains(l, "host is up") {
		return NmapUp
	}
	if strings.Contains(l, "host seems down") {
		return NmapDown
	}
	if err != nil {
		return NmapError
	}
	return NmapNoConfirm
}

func DecideStatus(p PingSignal, n NmapSignal) (FinalStatus, string) {

	if p == PingOK || n == NmapUp {
		if p != PingOK && n == NmapUp {
			return StatusUP, "Host up, but ping blocked → possible ICMP filtered (firewall/ACL)."
		}
		return StatusUP, "Host reachable."
	}

	if p == PingUnreachable && n != NmapUp {
		return StatusDOWN, "Unreachable from this host (routing/gateway/ACL)."
	}

	return StatusUNKNOWN, "No confirmation. Host may be down or probes filtered."
}

func Check(target string) (Result, error) {
	// Dependecy Check
	if _, err := exec.LookPath("ping"); err != nil {
		return Result{}, fmt.Errorf("ping not found. Install: sudo apt install iputils-ping")
	}
	if _, err := exec.LookPath("nmap"); err != nil {
		return Result{}, fmt.Errorf("nmap not found. Install: sudo apt install nmap")
	}

	if err := system.ValidateResolvable(target); err != nil {
		return Result{}, err
	}

	p := pingCheck(target)
	n := nmapSnCheck(target)
	st, hint := DecideStatus(p, n)

	return Result{
		Target: target,
		Ping:   p,
		NmapSN: n,
		Status: st,
		Hint:   hint,
	}, nil
}
