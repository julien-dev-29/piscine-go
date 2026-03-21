package piscine

func Atoi(s string) int {
	res := 0
	i := 0
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	}
	if s[i] == '+' {
		i++
	}
	if i == StrLen(s) {
		return 0
	}
	for ; i < StrLen(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0
		}
		res = res*10 + int(c-'0')
	}
	return res * sign
}
