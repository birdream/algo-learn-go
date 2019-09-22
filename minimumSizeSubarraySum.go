package learn

// Minimum Size Subarray Sum
// 给定一个整数数组和数字target 找到数组中最短的一个连续子数组，使得连续子数组和>=target
// 返回这个最短的连续子数组的返回值
// 没有则返回0
func method(arr []int, target int) interface{} {
	l := 0
	r := -1 // 一开始不包含任何元素
	sum := 0
	res := make([]int, len(arr)+1)

	for l < len(arr)-1 {
		if r+1 <= len(arr)-1 && sum < target {
			r++
			sum += arr[r]
		} else {
			sum -= arr[l]
			l++
		}

		if sum >= target && len(res) > r-l+1 {
			res = arr[l : r+1]
		}
	}

	if len(res) == len(arr)+1 {
		return 0
	}

	return res
}
