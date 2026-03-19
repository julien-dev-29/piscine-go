package piscine

import "fmt"

func BasicAtoi(s string) int {
	res := 0
	for i := range s {
		res = int(rune(s[i] - '0'))
		fmt.Println(res)
	}
	return res
}
