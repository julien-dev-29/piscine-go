package piscine

func Capitalize(s string) string {
	rs := []rune(s)
	for i := range rs {
		if i == 0 && isLower(rs[i]) {
			rs[i] = rs[i] - 32
		}
		if i > 0 && !isAlpha(rs[i-1]) && isLower(rs[i]) {
			rs[i] = rs[i] - 32
		}
	}
	return string(rs)
}

func isLower(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	} else {
		return false
	}
}

func isAlpha(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	} else {
		return false
	}
}
