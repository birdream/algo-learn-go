package main

// InsertionSort1 native
func InsertionSort1(arr []int) {
	length := len(arr)

	for i := 1; i < length; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
