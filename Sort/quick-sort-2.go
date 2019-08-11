package main

import "fmt"

func quickSort2(sortArr []int, left, right int) {
	fmt.Println("------------------------------------in")
	fmt.Println("left: ", left)
	fmt.Println("right: ", right)

	if left < right {
		fmt.Println("----in1")
		key := sortArr[left]
		i, j := left+1, right

		for {
			fmt.Println("---111")
			for i <= j && sortArr[i] < key {
				i++
			}
			fmt.Println("---222")
			for i < j && sortArr[j] > key {
				j--
			}
			fmt.Println("---333")
			fmt.Println("i: ", i)
			fmt.Println("j: ", j)

			if i >= j {
				break
			}

			if sortArr[i] > key || sortArr[j] < key {
				break
			}

			sortArr[i], sortArr[j] = sortArr[j], sortArr[i]

			fmt.Println("---444")
		}

		fmt.Println("---5555")
		fmt.Println("left: ", left)
		fmt.Println("right: ", right)
		quickSort2(sortArr, left, i-1)
		quickSort2(sortArr, j+1, right)
	}
}

// QuickSort2 lalala
func QuickSort2(val []int) {
	quickSort2(val, 0, len(val)-1)
}
