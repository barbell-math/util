package strops

import (
	"strings"
)

// Breaks a body of text into separate lines that are no longer than maxLen.
// Line splits are made on word (a.k.a. white space delimited) boundaries.
func WordWrapLines(text string, maxLen int) []string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return []string{text}
	}

	line := words[0]
	wrapped := []string{}
	spaceLeft := maxLen - len(line)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped = append(wrapped, line)
			line = word
			spaceLeft = maxLen - len(word)
		} else {
			line += " " + word
			spaceLeft -= len(word) + 1
		}
	}
	if len(line) > 0 {
		wrapped = append(wrapped, line)
	}
	return wrapped
}
