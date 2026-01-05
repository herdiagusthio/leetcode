package main

import (
	"fmt"
	"strings"
)

// decodeString decodes an encoded string using the pattern k[encoded_string].
// The encoded_string inside brackets is repeated exactly k times.
// Supports nested encodings like "3[a2[c]]" which becomes "accaccacc".
func decodeString(s string) string {
	var currentNumber int
	var currentString strings.Builder

	type state struct {
		prevString string
		repeat     int
	}

	var saveState []state
	for _, c := range s {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// READ_NUMBER state
			digit := int(c - '0')
			currentNumber = currentNumber*10 + digit
		case '[':
			// ENTER_CONTEXT: save current state and reset
			saveState = append(saveState, state{
				prevString: currentString.String(),
				repeat:     currentNumber,
			})

			// RESET STATE
			currentString.Reset()
			currentNumber = 0

		case ']':
			// EXIT_CONTEXT: pop state and build repeated string
			prevState := saveState[len(saveState)-1]
			saveState = saveState[:len(saveState)-1]

			// Repeat currentString
			innerString := currentString.String()
			repeatedString := strings.Repeat(innerString, prevState.repeat)

			// Rebuild currentString with previous content + repeated string
			currentString.Reset()
			currentString.WriteString(prevState.prevString)
			currentString.WriteString(repeatedString)
		default:
			// READ_STRING: append character
			currentString.WriteRune(c)
		}
	}

	return currentString.String()
}

func main() {
	// Example 1: Simple repetition
	fmt.Println(decodeString("3[a]2[bc]")) // "aaabcbc"

	// Example 2: Nested encoding
	fmt.Println(decodeString("3[a2[c]]")) // "accaccacc"

	// Example 3: Mixed with regular characters
	fmt.Println(decodeString("2[abc]3[cd]ef")) // "abcabccdcdcdef"

	// Example 4: Complex nesting
	fmt.Println(decodeString("10[a]")) // "aaaaaaaaaa"
}
