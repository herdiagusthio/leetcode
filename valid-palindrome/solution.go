package main

import (
	"fmt"
	"regexp"
	"strings"
)

var reg = regexp.MustCompile("[^a-zA-Z0-9]+")

func isPalindrome(s string) bool {

	cleaned := reg.ReplaceAllString(s, "")
	cleaned = strings.ToLower(cleaned)

	for f, b := 0, len(cleaned)-1; f < b; f, b = f+1, b-1 {
		if cleaned[f] != cleaned[b] {
			return false
		}
	}

	return true
}

func main() {
	samples := []string{
		"",
		"A man, a plan, a canal: Panama",
		"race a car",
		"0P",
		".,",
		"ab_a",
	}

	for _, s := range samples {
		fmt.Printf("%q -> %v\n", s, isPalindrome(s))
	}
}
