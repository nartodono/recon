package ui

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
)

func Red(s string) string {
	return ColorRed + s + ColorReset
}
func Green(s string) string {
	return ColorGreen + s + ColorReset
}
func Yellow(s string) string {
	return ColorYellow + s + ColorReset
}
func Cyan(s string) string {
	return ColorCyan + s + ColorReset
}