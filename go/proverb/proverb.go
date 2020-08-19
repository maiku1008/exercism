package proverb

import "fmt"

// Proverb returns a proverb made out of rhymes
func Proverb(rhyme []string) []string {
	proverb := make([]string, len(rhyme))
	for i, v := range rhyme {
		if i != len(rhyme)-1 {
			proverb[i] = fmt.Sprintf("For want of a %s the %s was lost.", v, rhyme[i+1])
			continue
		}
		proverb[i] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	}
	return proverb
}
