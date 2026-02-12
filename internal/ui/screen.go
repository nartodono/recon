package ui

import "fmt"

func ClearScreen() {
	// ANSI clear screen + move cursor to home (Linux terminal)
	fmt.Print("\033[H\033[2J")
}
