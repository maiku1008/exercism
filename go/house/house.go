package house

import (
	"fmt"
	"strings"
)

type Record struct {
	name   string
	action string
}

var records = []*Record{
	{"house that Jack built", "lay in"},
	{"malt", "ate"},
	{"rat", "killed"},
	{"cat", "worried"},
	{"dog", "tossed"},
	{"cow with the crumpled horn", "milked"},
	{"maiden all forlorn", "kissed"},
	{"man all tattered and torn", "married"},
	{"priest all shaven and shorn", "woke"},
	{"rooster that crowed in the morn", "kept"},
	{"farmer sowing his corn", "belonged to"},
	{"horse and the hound and the horn", ""},
}

func Verse(n int) string {
	var result strings.Builder
	result.WriteString("This is the " + records[n-1].name)
	verses := strings.Join(verse(n-1), "")
	result.WriteString(verses)
	result.WriteString(".")
	return result.String()
}

func verse(i int) []string {
	if i == 0 {
		return nil
	}
	v := fmt.Sprintf("\nthat %s the %s", records[i-1].action, records[i-1].name)
	return append([]string{v}, verse(i-1)...)
}

func Song() string {
	verses := song(len(records))
	return strings.Join(verses, "\n\n")
}

func song(i int) []string {
	if i == 0 {
		return nil
	}
	return append(song(i-1), Verse(i))
}
