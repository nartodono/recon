package export

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func DefaultDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home dir: %w", err)
	}
	return filepath.Join(home, "recon"), nil
}

func EnsureDir(dir string) error {
	return os.MkdirAll(dir, 0o755)
}

func Filename(module string, ext string, t time.Time) string {
	// recon-host-YYYYMMDD-HHMMSS.json
	ts := t.Format("20060102-150405")
	return fmt.Sprintf("recon-%s-%s.%s", module, ts, ext)
}

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0o644)
}
