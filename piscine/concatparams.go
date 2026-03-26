package piscine

import "strings"

func ConcatParams(args []string) string {
	var res strings.Builder
	for i, arg := range args {
		res.WriteString(arg)
		if i < len(args)-1 {
			res.WriteRune('\n')
		}
	}
	return res.String()
}
