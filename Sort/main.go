package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createRandomArr(len int) []int {
	res := make([]int, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		res[i] = rand.Intn(150) - 50
	}

	return res
}

func main() {
	arr1 := createRandomArr(50)
	fmt.Print(arr1, "\n")
	QuickSort(arr1)
	fmt.Print(arr1, "\n")

	fmt.Print("----------------------\n")

	arr2 := createRandomArr(50)
	fmt.Print(arr2, "\n")
	QuickSort2(arr2)
	fmt.Print(arr2, "\n")
	// fmt.Print(createRandomArr(12), "\n")
	// fmt.Print(createRandomArr(12), "\n")
	// fmt.Print((3 / 2), "\n")
}
