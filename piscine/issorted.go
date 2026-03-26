package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}
	var direction int = 0
	for i := 1; i < len(a); i++ {
		result := f(a[i-1], a[i])
		if result != 0 {
			if direction == 0 {
				direction = result
			} else if (direction > 0 && result < 0) || (direction < 0 && result > 0) {
				return false
			}
		}
	}
	return true
}
