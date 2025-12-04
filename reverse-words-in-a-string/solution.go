package main

import (
	"fmt"
	"strings"
)

// reverseWords reverses the order of words in the input string. It trims
// leading/trailing spaces and reduces multiple spaces between words to a
// single space (behavior provided by strings.Fields).
func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	examples := []string{
		"the sky is blue",
		"  hello world!  ",
		"a good   example",
		"",
		"    ",
	}

	for _, ex := range examples {
		fmt.Printf("input: %q -> output: %q\n", ex, reverseWords(ex))
	}
}
