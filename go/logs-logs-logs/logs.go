package logs

import (
	"fmt"
	"strings"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	return strings.Join(strings.Fields(line)[1:], " ")
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	return len([]rune(Message(line)))
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	line = strings.Split(line, ":")[0]
	line = strings.Trim(line, "[]")
	return strings.ToLower(line)
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return fmt.Sprintf("%s (%s)", Message(line), LogLevel(line))
}
