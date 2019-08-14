package main

import (
	. "algo-learn-go/Sort/utils"
)

func main() {
	arr2 := CreateRandomArr(20, true)
	InsertionSort1(arr2)
	IsSorted(arr2, true)
}
