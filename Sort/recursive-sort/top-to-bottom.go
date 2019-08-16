package main

// TopToBottom 递归排序: 自上向下
func TopToBottom(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	mid := length / 2

	return merger(TopToBottom(arr[:mid]), TopToBottom(arr[mid:]))
}

func merger(arr1, arr2 []int) (resp []int) {
	l1 := len(arr1)
	l2 := len(arr2)

	i, j := 0, 0

	for i < l1 || j < l2 {
		switch {
		case j == l2:
			resp = append(resp, arr1[i])
			i++
		case i == l1:
			resp = append(resp, arr2[j])
			j++

			// i < l1, j < l2
		case arr1[i] <= arr2[j]:
			resp = append(resp, arr1[i])
			i++

		default:
			// arr2[i] > arr2[j]
			resp = append(resp, arr2[j])
			j++
		}
	}

	return resp
}
