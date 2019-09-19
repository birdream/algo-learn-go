package learn

// 数组已经排序 asc
// 一定有解
// 返回任意解
// O(n)
func twoSum(arr []int, target int) []int {
	i := 0
	j := len(arr) - 1

	for i < j {
		if arr[i]+arr[j] == target {
			return []int{i, j}
		} else if arr[i]+arr[j] < target {
			i++
		} else {
			j--
		}
	}

	return []int{}
}
