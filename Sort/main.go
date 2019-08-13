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

func isSorted(arr []int) {
	wrong := false
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			fmt.Println("Wrong Sort", arr[i-1], arr[i])
			wrong = true
			break
		}
	}
	if !wrong {
		fmt.Println("Correct Sort")
	}
}

func main() {
	// arr1 := createRandomArr(50)
	// fmt.Print(arr1, "\n")
	// QuickSort(arr1)
	// fmt.Print(arr1, "\n")

	// fmt.Print("----------------------\n")

	arr2 := createRandomArr(50)
	fmt.Print(arr2, "\n")
	QSortTwoWay(arr2)
	fmt.Print(arr2, "\n")
	isSorted(arr2)
	fmt.Print("----------------------\n")

	arr3 := createRandomArr(50)
	fmt.Print(arr3, "\n")
	QSortSingle(arr3)
	fmt.Print(arr3, "\n")
	isSorted(arr3)
	fmt.Print("----------------------\n")

	arr4 := createRandomArr(50)
	fmt.Print(arr4, "\n")
	QuickSortThreeWays(arr4)
	fmt.Print(arr4, "\n")
	isSorted(arr4)
	// fmt.Print(createRandomArr(12), "\n")
	// fmt.Print(createRandomArr(12), "\n")
	// fmt.Print((3 / 2), "\n")
}
