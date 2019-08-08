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

	fmt.Print(createRandomArr(10), "\n")
	fmt.Print(createRandomArr(12), "\n")
	fmt.Print(createRandomArr(12), "\n")
	fmt.Print((3 / 2), "\n")
}
