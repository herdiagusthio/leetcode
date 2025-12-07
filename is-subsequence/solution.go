package main

import "fmt"

// isSubsequence reports whether s is a subsequence of t.
func isSubsequence(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)

	if lenS == 0 {
		return true
	}
	if lenS > lenT {
		return false
	}

	pointerS := 0

	for i := 0; i < lenT; i++ {
		if t[i] == s[pointerS] {
			pointerS++
			if pointerS == lenS {
				return true
			}
		}
	}

	return false
}

func main() {
	examples := []struct {
		s string
		t string
	}{
		{"abc", "ahbgdc"},
		{"axc", "ahbgdc"},
		{"", "anything"},
		{"ace", "abcde"},
		{"aec", "abcde"},
	}

	for _, ex := range examples {
		fmt.Printf("isSubsequence(%q, %q) = %v\n", ex.s, ex.t, isSubsequence(ex.s, ex.t))
	}
}
