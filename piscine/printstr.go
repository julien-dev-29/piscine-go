package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(rune(c))
	}
	z01.PrintRune('\n')
}
