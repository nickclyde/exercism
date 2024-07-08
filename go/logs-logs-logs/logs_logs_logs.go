package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	apps := map[string]int{
		"recommendation": strings.IndexRune(log, '❗'),
		"search":         strings.IndexRune(log, '🔍'),
		"weather":        strings.IndexRune(log, '☀'),
	}
	firstRuneAppearance := utf8.RuneCountInString(log)
	firstRune := "default"
	for k, v := range apps {
		if v == -1 {
			continue
		}
		if v < firstRuneAppearance {
			firstRune = k
			firstRuneAppearance = v
		}
	}
	return firstRune
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
