package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	for range n {
		for j := 0; j < n-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
