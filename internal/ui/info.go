package ui

import "fmt"

func PrintSaved(path string) {
	fmt.Println(Cyan("[i] Saved: ") + path)
}
