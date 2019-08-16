package main

import (
	. "algo-learn-go/Sort/utils"
)

func main() {
	arr2 := CreateRandomArr(503, true, 500)
	arr2 = TopToBottom(arr2)
	IsSorted(arr2, true)
}
