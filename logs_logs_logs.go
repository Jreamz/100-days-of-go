package logs

import "fmt"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for i, v := range log {
		switch v {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		default:
			if i > len(log) {
				break
			}
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := ""

	for _, v := range log {
		if v == oldRune {
			newLog += string(newRune)
		} else {
			newLog += string(v)
		}
	}
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is within the limit.
func WithinLimit(log string, limit int) bool {
	counter := 0

	for _, v := range log {
		counter += 1
		fmt.Println(v)
	}

	if counter > limit {
		return false
	}

	return true
}
