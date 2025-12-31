package analyzer

import (
	"strings"
)

func CommitAn(message string) string {
	part := strings.SplitN(message, ":", 2) // Divide the string in first ":"

	if len(part) < 2 {
		return "Another's commits"
	}
	// Return the first part, clean without additional spaces
	return strings.ToLower(strings.TrimSpace(part[0]))
}
