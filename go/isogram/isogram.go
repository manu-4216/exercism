package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks if a word is an isogram
func IsIsogram(s string) bool {
	var runePresence = map[rune]bool{}

	for _, r := range strings.ToUpper(s) {
		if !unicode.IsLetter(r) {
			continue
		}

		if runePresence[r] {
			return false
		}

		runePresence[r] = true
	}

	return true
}
