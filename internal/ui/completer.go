package ui

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/chzyer/readline"
)

type RLCompleter struct {
	Commands *readline.PrefixCompleter
}

func (c *RLCompleter) Do(line []rune, pos int) ([][]rune, int) {
	s := string(line[:pos])

	tokens, cur := splitForCompletion(s)

	// file completion hanya jika token sebelumnya adalah "-f"
	if len(tokens) > 0 && tokens[len(tokens)-1] == "-f" {
		// readline akan append suggestion ke token, jadi kita return "suffix" saja
		return completePathSuffix(cur), 0
	}

	// default: command completion
	if c.Commands != nil {
		return c.Commands.Do(line, pos)
	}
	return nil, 0
}

func splitForCompletion(s string) ([]string, string) {
	s = strings.ReplaceAll(s, "\t", " ")

	trimRight := strings.TrimRight(s, " ")
	if trimRight != s {
		// trailing space: current token kosong
		fields := strings.Fields(trimRight)
		return fields, ""
	}

	lastSpace := strings.LastIndex(s, " ")
	if lastSpace == -1 {
		return []string{}, s
	}
	head := s[:lastSpace]
	cur := s[lastSpace+1:]
	fields := strings.Fields(head)
	return fields, cur
}

// Return suffix to append to current token, not full candidate.
func completePathSuffix(cur string) [][]rune {
	home, _ := os.UserHomeDir()
	sep := string(os.PathSeparator)

	typedHasTilde := strings.HasPrefix(cur, "~")

	// expand ~ for filesystem access
	curFS := cur
	if typedHasTilde {
		curFS = filepath.Join(home, strings.TrimPrefix(cur, "~"))
	}

	// determine dir to read + base to match
	dirFS := "."
	base := curFS

	// dirTyped preserves what user typed (keeps "~" if used)
	dirTyped := ""

	if strings.Contains(curFS, sep) {
		if strings.HasSuffix(curFS, sep) {
			dirFS = curFS
			base = ""
		} else {
			dirFS = filepath.Dir(curFS)
			base = filepath.Base(curFS)
		}

		// compute typed directory prefix from original cur
		if strings.HasSuffix(cur, sep) {
			dirTyped = cur
		} else {
			idx := strings.LastIndex(cur, sep)
			if idx >= 0 {
				dirTyped = cur[:idx+1]
			}
		}
	} else {
		dirFS = "."
		base = curFS
		dirTyped = ""
	}

	entries, err := os.ReadDir(dirFS)
	if err != nil {
		return nil
	}

	baseLower := strings.ToLower(base)

	// build full typed-style suggestions then convert to suffix
	full := make([]string, 0, 32)
	for _, e := range entries {
		name := e.Name()
		if base == "" || strings.HasPrefix(strings.ToLower(name), baseLower) {
			s := dirTyped + name
			if e.IsDir() {
				s += sep
			}
			full = append(full, s)
		}
	}

	sort.Strings(full)

	out := make([][]rune, 0, len(full))
	for _, f := range full {
		// key behavior: only append what user hasn't typed yet
		if strings.HasPrefix(f, cur) {
			out = append(out, []rune(strings.TrimPrefix(f, cur)))
		} else {
			// fallback (should be rare)
			out = append(out, []rune(f))
		}
	}

	return out
}
