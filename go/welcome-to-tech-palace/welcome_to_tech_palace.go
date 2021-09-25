package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s\n", strings.Repeat("*", numStarsPerLine))
	fmt.Fprintf(&b, "%s\n", welcomeMsg)
	fmt.Fprintf(&b, "%s", strings.Repeat("*", numStarsPerLine))
	return b.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.Trim(oldMsg, "*\n ")
}
