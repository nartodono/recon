package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

func RunShell() {
	PrintBanner()

	executor := func(in string) {
		in = strings.TrimSpace(in)
		if in == "" {
			return
		}
		parts := strings.Fields(in)
		cmd := parts[0]
		args := parts[1:]

		// NOTE: kalau RunCommand return true (exit/quit), kita exit proses
		if RunCommand(cmd, args) {
			fmt.Println("Bye 👋")
			os.Exit(0)
		}
	}

	completer := func(d prompt.Document) []prompt.Suggest {
		text := d.TextBeforeCursor()
		text = strings.TrimLeft(text, " \t")

		// Kalau lagi completion path setelah "-f"
		if isFileArgContext(text) {
			cur := currentWordForFileArg(text) // token setelah -f (bisa kosong)
			return completePathSuggest(cur)
		}

		// Default: completion command & subcommand sederhana
		return completeCommandSuggest(d.GetWordBeforeCursor(), text)
	}

	// WordSeparator penting: jangan pisahin kata di "/" supaya path tetap 1 token
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(Cyan("recon")+" > "),
	)
	p.Run()
}

// ---------------------------
// Completion helpers
// ---------------------------

func completeCommandSuggest(word string, fullLine string) []prompt.Suggest {
	// suggestions command utama
	base := []prompt.Suggest{
		{Text: "help", Description: "Show help"},
		{Text: "?", Description: "Show help"},
		{Text: "clear", Description: "Clear screen"},
		{Text: "cls", Description: "Clear screen"},
		{Text: "exit", Description: "Exit"},
		{Text: "quit", Description: "Exit"},
		{Text: "q", Description: "Exit"},
		{Text: "host", Description: "Host check (ping + nmap -sn)"},
		{Text: "port", Description: "Port scan (nmap)"},
	}

	// sub-suggest kalau sudah ketik command tertentu
	fields := strings.Fields(fullLine)
	if len(fields) >= 1 {
		switch fields[0] {
		case "host":
			// hanya kasih hint flag -f (file)
			base = []prompt.Suggest{
				{Text: "-f", Description: "Load targets from file"},
			}
		case "port":
			base = []prompt.Suggest{
				{Text: "default", Description: "nmap -sC -sV"},
				{Text: "aggr", Description: "nmap -A"},
				{Text: "-f", Description: "Load targets from file"},
			}
		}
	}

	return prompt.FilterHasPrefix(base, word, true)
}

// Detect context: completion untuk arg setelah "-f"
func isFileArgContext(text string) bool {
	// Kita pakai fields untuk deteksi posisi "-f"
	// Kasus valid:
	//   "host -f "           => fields=["host","-f"] => true
	//   "host -f p"          => fields=["host","-f","p"] => true
	//   "port -f ./ta"       => true
	//   "port aggr -f ./ta"  => true (kalau kamu nanti mau, ini sudah support)
	fields := strings.Fields(text)
	if len(fields) == 0 {
		return false
	}

	// Cari "-f" terakhir (biar aman kalau ada yang ngetik "-f" lebih dari sekali)
	lastF := -1
	for i := range fields {
		if fields[i] == "-f" {
			lastF = i
		}
	}
	if lastF == -1 {
		return false
	}

	// Kalau cursor sudah lewat "-f", berarti kita sedang mengisi arg setelah -f
	// True kalau:
	// 1) "-f" adalah token terakhir (artinya user baru ketik -f dan spasi atau belum isi)
	// 2) "-f" adalah token ke-(n-2) (artinya token terakhir adalah path fragment)
	if lastF == len(fields)-1 {
		return true
	}
	if lastF == len(fields)-2 {
		return true
	}

	// Kalau ada token-token lain setelah itu (misal user sudah kasih path lalu ngetik arg lain),
	// kita tidak autocomplete path lagi.
	return false
}

// Mengambil current token setelah -f (path fragment) dari line raw
func currentWordForFileArg(text string) string {
	// Kalau text berakhir dengan spasi, berarti current word kosong
	if strings.TrimRight(text, " ") != text {
		// trailing spaces
		fields := strings.Fields(strings.TrimSpace(text))
		if len(fields) > 0 && fields[len(fields)-1] == "-f" {
			return ""
		}
	}

	fields := strings.Fields(text)
	if len(fields) == 0 {
		return ""
	}

	// cari "-f" terakhir
	lastF := -1
	for i := range fields {
		if fields[i] == "-f" {
			lastF = i
		}
	}
	if lastF == -1 {
		return ""
	}

	// kalau "-f" token terakhir => empty fragment
	if lastF == len(fields)-1 {
		return ""
	}

	// kalau token setelah "-f" ada => itu fragment
	if lastF == len(fields)-2 {
		return fields[len(fields)-1]
	}

	return ""
}

func completePathSuggest(cur string) []prompt.Suggest {
	home, _ := os.UserHomeDir()

	// Expand ~ untuk listing FS, tapi suggestion tetap pakai "~" kalau user mulai dengan "~"
	typedHasTilde := strings.HasPrefix(cur, "~")
	curFS := cur
	if typedHasTilde {
		curFS = filepath.Join(home, strings.TrimPrefix(cur, "~"))
	}

	sep := string(os.PathSeparator)

	// Tentukan dir untuk list + base prefix
	dirFS := "."
	dirTyped := ""
	base := curFS

	if strings.Contains(curFS, sep) {
		if strings.HasSuffix(curFS, sep) {
			// user sudah menunjuk folder (ends with /) => list isi folder tsb
			dirFS = curFS
			base = ""
		} else {
			dirFS = filepath.Dir(curFS)
			base = filepath.Base(curFS)
		}

		// prefix yang akan dipakai di suggestion (pakai yang user ketik, bukan hasil expand)
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
		dirTyped = ""
		base = curFS
	}

	entries, err := os.ReadDir(dirFS)
	if err != nil {
		return nil
	}

	baseLower := strings.ToLower(base)

	type cand struct {
		text string
		desc string
	}
	cands := make([]cand, 0, 32)

	for _, e := range entries {
		name := e.Name()
		// match prefix (case-insensitive)
		if base == "" || strings.HasPrefix(strings.ToLower(name), baseLower) {
			s := dirTyped + name
			desc := "file"
			if e.IsDir() {
				s += sep
				desc = "dir"
			}
			cands = append(cands, cand{text: s, desc: desc})
		}
	}

	sort.Slice(cands, func(i, j int) bool { return cands[i].text < cands[j].text })

	out := make([]prompt.Suggest, 0, len(cands))
	for _, c := range cands {
		out = append(out, prompt.Suggest{Text: c.text, Description: c.desc})
	}
	return out
}
