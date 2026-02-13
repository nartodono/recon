package export

import (
	"encoding/json"
	"fmt"
)

func WriteJSON(path string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}
	b = append(b, '\n')
	return WriteFile(path, b)
}
