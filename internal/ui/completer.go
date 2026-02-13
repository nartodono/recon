package ui

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/chzyer/readline"
)

type HybridCompleter struct {
	Cmd *readline.PrefixCompleter
}

func (h *HybridCompleter) Do(line []rune, pos int) ([][]rune, int) {
	s := string(line[:pos])

	tokens, cur := splitForCompletion(s)

	// only do file completion when the token right before cursor is "-f"
	isFileContext := len(tokens) >= 1 && tokens[len(tokens)-1] == "-f"
	if isFileContext {
		// IMPORTANT: return suffixes to append, not full paths
		return completePathSuffix(cur), 0
	}

	if h.Cmd != nil {
		return h.Cmd.Do(line, pos)
	}
	return nil, 0
}

func splitForCompletion(s string) ([]string, string) {
	s = strings.ReplaceAll(s, "\t", " ")

	trimRight := strings.TrimRight(s, " ")
	if trimRight != s {
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

// completePathSuffix returns the *suffix* to append to current token (cur)
// so tab completion behaves like a normal shell.
func completePathSuffix(cur string) [][]rune {
	home, _ := os.UserHomeDir()

	// Expand "~" only for filesystem operations
	curFS := cur
	if strings.HasPrefix(cur, "~") {
		curFS = filepath.Join(home, strings.TrimPrefix(cur, "~"))
	}

	// Determine filesystem directory to list + base prefix to match
	dirFS := "."
	base := curFS

	// Determine what user typed as directory prefix (keeps "~" if present)
	dirTyped := ""

	sep := string(os.PathSeparator)

	if strings.Contains(curFS, sep) {
		if strings.HasSuffix(curFS, sep) {
			// user typed a dir and ended with '/', match everything inside it
			dirFS = curFS
			base = ""
		} else {
			dirFS = filepath.Dir(curFS)
			base = filepath.Base(curFS)
		}

		// typed dir prefix should be based on original cur (not expanded),
		// so suggestions keep "~" if user typed "~"
		if strings.HasSuffix(cur, sep) {
			dirTyped = cur
		} else {
			idx := strings.LastIndex(cur, sep)
			if idx >= 0 {
				dirTyped = cur[:idx+1]
			}
		}
	} else {
		// no slash, complete in current working directory
		dirFS = "."
		base = curFS
		dirTyped = ""
	}

	entries, err := os.ReadDir(dirFS)
	if err != nil {
		return nil
	}

	baseLower := strings.ToLower(base)

	// Build full suggestion strings (typed-style), then convert to suffixes
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

	// Convert full suggestions to suffixes to append
	out := make([][]rune, 0, len(full))
	for _, f := range full {
		// append only the part beyond what user already typed (cur)
		if strings.HasPrefix(f, cur) {
			out = append(out, []rune(strings.TrimPrefix(f, cur)))
		} else {
			// fallback: if somehow doesn't match, just append the whole thing
			out = append(out, []rune(f))
		}
	}

	return out
}
