package strand

import "strings"

var nucleotides = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var rna strings.Builder
	for _, v := range dna {
		rna.WriteRune(nucleotides[v])
	}
	return rna.String()
}
