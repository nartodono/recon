package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunShell() {
	PrintBanner()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n[!] Input error: %v\n", err)
			return
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		cmd := parts[0]
		args := parts[1:]

		if RunCommand(cmd, args) {
			return
		}
	}
}