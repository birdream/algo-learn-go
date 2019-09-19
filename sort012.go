package learn

// 将一个只含 0，1，2的数组快速排序
func sortColors(arr []int) []int {
	storage := make(map[int]int, 3)
	arrLen := len(arr)

	res := make([]int, 0)
	for i := 0; i < arrLen; i++ {
		storage[arr[i]]++
	}

	for i := 0; i < storage[0]; i++ {
		res = append(res, 0)
	}
	for i := 0; i < storage[1]; i++ {
		res = append(res, 1)
	}
	for i := 0; i < storage[2]; i++ {
		res = append(res, 2)
	}

	return res
}

// 3路快排
// 数组分为3个区间 0在最前面 1在中间 2在最后面
// [0..zero, 1,...., two, 2, ...]
// 如果遍历到arr[i]为1则直接跳过
// 如果是2 则和arr[two]交换 i不需要++ 因为此时的 arr[two]还为遍历所以不知数值是多少 继续循环
// 如果是0 则和arr[zero+1]交换 因为arr[zero+1]一定是1 所以无需再判断 直接i++ 继续循环
func sortColorsBy3Way(arr []int) []int {
	zero := -1      // arr[0...zero]
	two := len(arr) // arr[two... len-1]

	for i := 0; i < two; {
		if arr[i] == 1 {
			i++
		} else if arr[i] == 2 {
			two--
			arr[i], arr[two] = arr[two], arr[i]
		} else if arr[i] == 0 {
			zero++
			arr[i], arr[zero] = arr[zero], arr[i]
			i++
		} else {
			//fmt.Errorf("shouldn't have 0 1 2 else")
		}
	}

	return arr
}
