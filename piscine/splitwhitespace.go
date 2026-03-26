package piscine

import (
	"fmt"
	"strings"
)

func SplitWhiteSpaces(s string) []string {
	var res []string
	var currentWord strings.Builder
	for _, r := range s {
		if r == ' ' {
			if currentWord.Len() > 0 {
				res = append(res, currentWord.String())
				currentWord.Reset()
			}
		} else {
			currentWord.WriteRune(r)
		}
	}
	if currentWord.Len() > 0 {
		res = append(res, currentWord.String())
	}
	fmt.Println(res)
	return res
}
