package piscine

func BasicAtoi2(s string) int {
	res := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		res = res*10 + int(c-'0')
	}
	return res
}
