package ui

import (
	"fmt"
	"io"
	"os"
	"strings"
	"github.com/chzyer/readline"
)

func RunShell() {
	PrintBanner()

	// Autocomplete command utama
	completer := readline.NewPrefixCompleter(
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

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          Cyan("recon") + " > ",
		HistoryFile:     os.TempDir() + "/recon_history.tmp",
		AutoComplete:    completer,
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
			if len(line) == 0 {
				continue
			}
			continue
		} else if err == io.EOF {
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
