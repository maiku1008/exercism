package proverb

import (
	"fmt"
	"testing"
)

func TestProverb(t *testing.T) {
	for _, test := range stringTestCases {
		actual := Proverb(test.input)
		fmt.Println("expected length: ", len(test.expected))
		fmt.Println("actual length: ", len(actual))
		fmt.Printf("first item in actual slice: %q", actual[0])
		fmt.Println()
		for _, v := range actual {
			fmt.Println([]byte(v))
		}
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("FAIL %s - Proverb test [%s]\n\texpected: [%s],\n\tactual:   [%s]",
				test.description, test.input, test.expected, actual)
		}
		t.Logf("PASS %s", test.description)
	}
}

func BenchmarkProverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Proverb(test.input)
		}
	}
}
