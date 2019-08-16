package main

import (
	. "algo-learn-go/Sort/utils"
)

func main() {
	// arr1 := createRandomArr(50)
	// fmt.Print(arr1, "\n")
	// QuickSort(arr1)
	// fmt.Print(arr1, "\n")

	// fmt.Print("----------------------\n")

	arr2 := CreateRandomArr(50, true, 150)
	QSortTwoWay(arr2)
	IsSorted(arr2, true)

	arr3 := CreateRandomArr(50, true, 150)
	QSortSingle(arr3)
	IsSorted(arr3, true)

	arr4 := CreateRandomArr(50, true, 150)
	QuickSortThreeWays(arr4)
	IsSorted(arr4, true)
	// fmt.Print(createRandomArr(12), "\n")
	// fmt.Print(createRandomArr(12), "\n")
	// fmt.Print((3 / 2), "\n")
}
