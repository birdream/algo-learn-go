package main

// QuickSortThreeWays mmmmlll
func QuickSortThreeWays(arr []int) {
	quickSortThreeWays(arr, 0, len(arr)-1)
}

func quickSortThreeWays(arr []int, l, r int) {
	if l < r {
		lt, gt := sortThreeWay(arr, l, r)
		quickSortThreeWays(arr, l, lt-1)
		quickSortThreeWays(arr, gt, r)
	}
}

func sortThreeWay(arr []int, l, r int) (lt, gt int) {
	mid := arr[l]

	lt = l
	gt = r + 1
	i := l + 1

	for {
		switch {
		case i >= gt:
			arr[l], arr[lt] = arr[lt], arr[l]

			return lt, gt
		case arr[i] == mid:
			i++
		case arr[i] < mid:
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		default:
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		}
	}
}
