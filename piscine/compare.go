package piscine

func Compare(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	minLen := min(len(rb), len(ra))
	for i := range minLen {
		if ra[i] > rb[i] {
			return 1
		}
		if ra[i] < rb[i] {
			return -1
		}
	}
	if len(ra) > len(rb) {
		return 1
	}
	if len(ra) < len(rb) {
		return -1
	}
	return 0
}
