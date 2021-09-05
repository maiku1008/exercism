package partyrobot

import (
	"fmt"
	"strings"
)

// Welcome greets a person by name.
// Uses string concatenation.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and stands out his age.
// Uses fmt.Sprintf.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
// Uses a strings.Builder.
func AssignTable(name string, table int, neighbour string, direction string, distance float64) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Welcome to my party, %s!\n", name)
	fmt.Fprintf(&b, "You have been assigned to table %X. ", table)
	fmt.Fprintf(&b, "Your table is %s, exactly %.1f meters from here.\n", direction, distance)
	fmt.Fprintf(&b, "You will be sitting next to %s.", neighbour)

	return b.String()
}
