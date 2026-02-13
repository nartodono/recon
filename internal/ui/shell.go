package ui

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/chzyer/readline"
)

func RunShell() {
	PrintBanner()

	// command completer
	cmdCompleter := readline.NewPrefixCompleter(
		readline.PcItem("help"),
		readline.PcItem("?"),
		readline.PcItem("clear"),
		readline.PcItem("cls"),
		readline.PcItem("exit"),
		readline.PcItem("quit"),
		readline.PcItem("q"),
		readline.PcItem("host",
			readline.PcItem("-f"),
		),
		readline.PcItem("port",
			readline.PcItem("default"),
			readline.PcItem("aggr"),
			readline.PcItem("-f"),
		),
	)

	home, _ := os.UserHomeDir()
	historyPath := filepath.Join(home, ".recon_history") // kalau kamu bener2 gak mau persist, boleh hapus field ini

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "recon > ", // sengaja plain biar gak glitch
		HistoryFile:     historyPath,
		AutoComplete:    &RLCompleter{Commands: cmdCompleter},
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init readline: %v\n", err)
		return
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			// Ctrl+C cancels line, back to prompt
			continue
		} else if err == io.EOF {
			// Ctrl+D exits
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		cmd := parts[0]
		args := parts[1:]

		if RunCommand(cmd, args) {
			break
		}
	}
}
