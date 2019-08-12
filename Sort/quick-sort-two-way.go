package main

var isPrint = true

func partitionTwoWay(arr []int, l, r int) int {
	mid := arr[l]

	left := l + 1
	right := r

	for {
		switch {
		case left > right:
			arr[l], arr[right] = arr[right], arr[l]
			return right
		case arr[left] < mid:
			left++
		case arr[right] > mid:
			right--
		default:
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
}

func qSortTwoWay(arr []int, l, r int) {
	if l > r {
		return
	}

	p := partitionTwoWay(arr, l, r)

	qSortTwoWay(arr, l, p-1)
	qSortTwoWay(arr, p+1, r)
}

// QSortTwoWay bbbbb...
func QSortTwoWay(arr []int) {
	qSortTwoWay(arr, 0, len(arr)-1)
}
