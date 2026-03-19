package piscine

func StrRev(s string) string {
	revstr := []rune(s)
	i := 0
	j := len(s) - 1
	for i < j {
		revstr[i], revstr[j] = revstr[j], revstr[i]
		i++
		j--
	}
	return string(revstr)
}
