package main

import "fmt"

func genParenthesis(left, right int, s []byte, result []string) []string {
	if left == 0 && right == 0 {
		result = append(result, string(s))
	}

	if left > right || left < 0 || right < 0 {
		return result
	}

	s = append(s, '(')
	result = genParenthesis(left-1, right, s, result)
	s = s[:len(s)-1]

	s = append(s, ')')
	result = genParenthesis(left, right-1, s, result)

	return result
}

func main() {
	var n int
	fmt.Print("Enter n: ")
	fmt.Scan(&n)
	
	s := make([]byte, 0, n)
	res := make([]string, 0, n)

	res = genParenthesis(n, n, s, res)
	fmt.Println(res)
}
