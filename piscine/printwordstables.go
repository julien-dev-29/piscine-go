package piscine

import "fmt"

func PrintWordsTables(a []string) {
	for i := range a {
		fmt.Println(a[i])
	}
}
