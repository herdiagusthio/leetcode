package main

func generateParenthesis(n int) []string {
	var result []string
	var str []byte
	generate(0, 0, n, str, &result)

	return result
}

func generate(o, c, n int, str []byte, result *[]string) {
	if len(str) == 2*n {
		*result = append(*result, string(str))
		return
	}

	if o < n {
		generate(o+1, c, n, append(str, '('), result)
	}

	if c < o {
		generate(o, c+1, n, append(str, ')'), result)
	}
}

func main() {
	samples := []int{1, 2, 3}
	for _, n := range samples {
		res := generateParenthesis(n)
		println("n=", n)
		for _, s := range res {
			println(s)
		}
		println("---")
	}
}
