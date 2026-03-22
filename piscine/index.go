package piscine

func Index(s string, toFind string) int {
	rs := []rune(s)
	rt := []rune(toFind)

	if len(rt) == 0 {
		return 0
	}
	for i := 0; i < len(rs)-len(rt); i++ {
		match := true
		for j := range rt {
			if rs[i+j] != rt[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
