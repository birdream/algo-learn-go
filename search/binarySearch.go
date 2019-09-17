package search

func binarySearch(arr []int, target int) int {
	length := len(arr)
	l, r := 0, length-1

	for l <= r {
		mid := (r + l) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			l = mid + 1
		}

		if arr[mid] > target {
			r = mid - 1
		}
	}

	return -1
}
