package main

import (
	. "algo-learn-go/Sort/utils"
)

func main() {
	arr2 := CreateRandomArr(20, true, 150)
	arr2 = LSD(arr2)
	IsSorted(arr2, true)
}
