package algo

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
