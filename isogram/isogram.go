package isogram

import (
	"strings"
)

// IsIsogram determines whether given string is or is not a isogram
func IsIsogram(s string) bool {
	r := strings.NewReplacer("-", "", " ", "")
	s = strings.ToLower(r.Replace(s))
	m := make(map[rune]int)

	if len(s) > 0 {
		for _, r := range s {
			m[r]++
			if m[r] > 1 {
						return false
				}
		}
	}

	return true
}