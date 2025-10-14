package main

import "fmt"

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	bracketMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		b := s[i]
		if val, exist := bracketMap[b]; exist {
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false
			}
			// pop
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, b)
		}
	}

	return len(stack) == 0
}

func main() {
	samples := []string{"()", "()[]{}", "(]", "([])", "([)]", "((((((", "])"}
	for _, s := range samples {
		fmt.Printf("%q -> %v\n", s, isValid(s))
	}
}
