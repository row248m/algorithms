package sort

func InsertionSort(items []int) []int {
	for i:= 1; i < len(items); i++ {
		val := items[i]
		k := i - 1

		for k >= 0 && items[k] > val {
			items[k], items[k + 1] = items[k + 1], items[k]
			k--
		}
	}

	return items
}