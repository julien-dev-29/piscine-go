package piscine

import "strings"

func Join(strs []string, sep string) string {
	var res strings.Builder
	for i, str := range strs {
		res.WriteString(str)
		if i < len(strs)-1 {
			res.WriteString(sep)
		}
	}
	return res.String()
}
