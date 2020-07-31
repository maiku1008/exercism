package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	var sentences []string
	// if len(rhyme) == 0 {
	// 	return sentences
	// }

	var lastSentence string
	if (len(rhyme) > 0 ) {
		lastSentence = "And all for the want of a " + rhyme[len(rhyme)-1] + "."
	}

	if (len(rhyme) == 1 ) {
		sentences = append(sentences, lastSentence)
	} else {
		for i , word := range rhyme {
			sentences = append(sentences, "For want of a " + word + " the " + rhyme[i+1] + " was lost.")
		}
		sentences = append(sentences, lastSentence)
	}

	return sentences
}
