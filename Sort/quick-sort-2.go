package main

func quickSort2(sortArr []int, left, right int) {
	if left < right {
		key := sortArr[(left+right)/2]
		i, j := left, right

		for {
			for sortArr[i] < key {
				i++
			}

			for sortArr[j] > key {
				j--
			}

			if i >= j {
				break
			}

			sortArr[i], sortArr[j] = sortArr[j], sortArr[i]
		}

		quickSort2(sortArr, left, i-1)
		quickSort2(sortArr, j+1, right)
	}
}

// QuickSort2 lalala
func QuickSort2(val []int) {
	quickSort2(val, 0, len(val)-1)
}
