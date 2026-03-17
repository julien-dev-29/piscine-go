package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	printRecursive(n)
}

func printRecursive(n int) {
	if n >= 10 {
		printRecursive(n / 10)
	}
	digit := n % 10
	z01.PrintRune(rune('0' + digit))
}
