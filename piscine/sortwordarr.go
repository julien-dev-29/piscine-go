package piscine

func SortWordArr(a []string) {
	for range a {
		for j := 1; j < len(a); j++ {
			if a[j-1] > a[j] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
