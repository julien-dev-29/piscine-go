package piscine

import "strings"

func BasicJoin(elems []string) string {
	var res strings.Builder
	for _, str := range elems {
		res.WriteString(str)
	}
	return res.String()
}
