package ui

func parseExportFlags(args []string) ([]string, bool, bool) {
	wantJSON := false
	wantTXT := false

	clean := make([]string, 0, len(args))
	for _, a := range args {
		switch a {
		case "--json":
			wantJSON = true
		case "--txt":
			wantTXT = true
		default:
			clean = append(clean, a)
		}
	}
	return clean, wantJSON, wantTXT
}
