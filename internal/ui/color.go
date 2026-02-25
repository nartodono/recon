package ui

const (
	ColorReset = "\033[0m"

	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"

	ColorBrightRed     = "\033[91m"
	ColorBrightGreen   = "\033[92m"
	ColorBrightYellow  = "\033[93m"
	ColorBrightBlue    = "\033[94m"
	ColorBrightMagenta = "\033[95m"
	ColorBrightCyan    = "\033[96m"
	ColorBrightWhite   = "\033[97m"
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
func Blue(s string) string {
	return ColorBlue + s + ColorReset
}

func Magenta(s string) string {
	return ColorMagenta + s + ColorReset
}

func White(s string) string {
	return ColorWhite + s + ColorReset
}

func BrightRed(s string) string {
	return ColorBrightRed + s + ColorReset
}

func BrightGreen(s string) string {
	return ColorBrightGreen + s + ColorReset
}

func BrightCyan(s string) string {
	return ColorBrightCyan + s + ColorReset
}
