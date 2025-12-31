package analyzer

import (
	"strings"
)

func CommitAn(message string) string {
	part := strings.SplitN(message, ":", 2)

	if len(part) < 2 {
		return "Another's commits"
	}

	return strings.ToLower(strings.TrimSpace(part[0]))
}
