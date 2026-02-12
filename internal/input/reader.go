package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"net"
)

func LoadTargetsFromFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	var targets []string
	seen := make(map[string]bool)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// support IP only
		line := strings.TrimSpace(scanner.Text())
		ip := net.ParseIP(line)
		if ip == nil || ip.To4() == nil {
			continue
		}

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		if idx := strings.Index(line, "#"); idx != -1 {
			line = strings.TrimSpace(line[:idx])
		}

		if line == "" {
			continue
		}

		if !seen[line] {
			seen[line] = true
			targets = append(targets, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return targets, nil
}
