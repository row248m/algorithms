package sort

// @todo: not event number

func merge(left, right []int) []int {
	size, iLeft, iRight := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if iLeft > len(left)-1 {
			slice[k] = right[iRight]
			iRight++
		} else if iRight > len(right)-1 {
			slice[k] = left[iLeft]
			iLeft++
		} else if left[iLeft] < right[iRight] {
			slice[k] = left[iLeft]
			iLeft++
		} else {
			slice[k] = right[iRight]
			iRight++
		}
	}
	return slice
}

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	middle := len(items) / 2

	return merge(MergeSort(items[middle:]), MergeSort(items[:middle]));
}
