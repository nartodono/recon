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

// Do implements readline.AutoCompleter
func (h *HybridCompleter) Do(line []rune, pos int) ([][]rune, int) {
	// current text until cursor
	s := string(line[:pos])

	// Parse tokens + current token (supports trailing space)
	tokens, cur := splitForCompletion(s)

	// Detect: are we completing the token right after "-f"?
	// Cases:
	//   "host -f p"      => tokens=["host","-f"], cur="p"
	//   "host -f "       => tokens=["host","-f"], cur=""
	//   "host -f ./ta"   => tokens=["host","-f"], cur="./ta"
	// If "-f" is last token in tokens -> yes, current token is file path
	isFileContext := len(tokens) >= 1 && tokens[len(tokens)-1] == "-f"
	if isFileContext {
		sug := completePath(cur)
		return sug, runeLen(cur)
	}

	// Otherwise, default command completer
	if h.Cmd != nil {
		return h.Cmd.Do(line, pos)
	}

	return nil, 0
}

// splitForCompletion splits command line into:
// - tokens: all completed tokens before the current token
// - cur: the current (possibly partial) token at cursor
//
// It treats consecutive spaces as separators and supports trailing space.
func splitForCompletion(s string) ([]string, string) {
	s = strings.ReplaceAll(s, "\t", " ")

	trimRight := strings.TrimRight(s, " ")
	if trimRight != s {
		// user ended with space => current token is empty
		fields := strings.Fields(trimRight)
		return fields, ""
	}

	// not ending with space
	// Find last space to get current token (without losing previous tokens)
	lastSpace := strings.LastIndex(s, " ")
	if lastSpace == -1 {
		return []string{}, s
	}
	head := s[:lastSpace]
	cur := s[lastSpace+1:]
	fields := strings.Fields(head)
	return fields, cur
}

func completePath(cur string) [][]rune {
	home, _ := os.UserHomeDir()

	// Expand ~ for filesystem ops, but keep ~ in suggestions if user typed it
	typedHasTilde := strings.HasPrefix(cur, "~")
	curFS := cur
	if typedHasTilde {
		curFS = filepath.Join(home, strings.TrimPrefix(cur, "~"))
	}

	// Determine directory + prefix (base)
	dirFS := "."
	dirTyped := ""
	base := cur

	// If contains '/', split by last slash
	if strings.Contains(curFS, string(os.PathSeparator)) {
		// If user typed ends with '/', base should be empty and dir is the full path
		if strings.HasSuffix(curFS, string(os.PathSeparator)) {
			dirFS = curFS
			base = ""
		} else {
			dirFS = filepath.Dir(curFS)
			base = filepath.Base(curFS)
		}

		// For typed path portion (what we return), we need the directory part too
		// Use the original 'cur' (not expanded), so suggestions keep "~" if used
		if strings.HasSuffix(cur, string(os.PathSeparator)) {
			dirTyped = cur
		} else {
			dirTyped = cur[:strings.LastIndex(cur, string(os.PathSeparator))+1]
		}
	} else {
		// No slash in input -> list current dir, match base=cur
		dirFS = "."
		dirTyped = ""
		base = cur
	}

	entries, err := os.ReadDir(dirFS)
	if err != nil {
		// if cannot read, no suggestions
		return nil
	}

	baseLower := strings.ToLower(base)

	cands := make([]string, 0, 32)
	for _, e := range entries {
		name := e.Name()
		if base == "" || strings.HasPrefix(strings.ToLower(name), baseLower) {
			s := dirTyped + name
			if e.IsDir() {
				s += string(os.PathSeparator) // add trailing /
			}
			cands = append(cands, s)
		}
	}

	sort.Strings(cands)

	out := make([][]rune, 0, len(cands))
	for _, c := range cands {
		out = append(out, []rune(c))
	}
	return out
}

func runeLen(s string) int {
	return len([]rune(s))
}
