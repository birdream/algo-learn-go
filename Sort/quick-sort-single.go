package main

func partitionSingle(arr []int, l, r int) int {
	mid := arr[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < mid {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[j], arr[l] = arr[l], arr[j]
	return j
}

func qSortSingle(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partitionSingle(arr, l, r)
	qSortSingle(arr, l, p-1)
	qSortSingle(arr, p+1, r)
}

// QSortSingle bbbbbb
func QSortSingle(arr []int) {
	qSortSingle(arr, 0, len(arr)-1)
}
