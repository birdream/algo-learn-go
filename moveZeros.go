package moveZeros

func moveZerosByK(arr []int) []int {
	k := 0
	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		if arr[i] != 0 {
			arr[k] = arr[i]
			k++
		}
	}

	for i := k; i < arrLen; i++ {
		arr[i] = 0
	}

	return arr
}

func moveZerosBySwitch(arr []int) []int {
	k := 0
	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		if arr[i] != 0 && k != 0 {
			arr[i], arr[k] = arr[k], arr[i]
			k = i
		}
	}

	return arr
}
