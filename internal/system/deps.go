package system

import (
	"fmt"
	"os/exec"
)

func CheckDeps() error {
	// Check nmap
	if _, err := exec.LookPath("nmap"); err != nil {
		return fmt.Errorf("nmap not found. Install it with: sudo apt install nmap")
	}

	// Check ping
	if _, err := exec.LookPath("ping"); err != nil {
		return fmt.Errorf("ping not found. Install it with: sudo apt install iputils-ping")
	}

	return nil
}