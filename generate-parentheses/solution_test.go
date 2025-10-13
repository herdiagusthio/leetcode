package main

import (
	"reflect"
	"sort"
	"testing"
)

func normalize(ss []string) []string {
	out := make([]string, len(ss))
	copy(out, ss)
	sort.Strings(out)
	return out
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{0, []string{""}},
		{1, []string{"()"}},
		{2, []string{"()()", "(())"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, tc := range tests {
		got := generateParenthesis(tc.n)
		if !reflect.DeepEqual(normalize(got), normalize(tc.want)) {
			t.Fatalf("n=%d\n got: %v\n want: %v", tc.n, got, tc.want)
		}
	}
}
