package main

import (
	"fmt"
	"math/rand"
)

func QSortTwoWay(arr []int) {

	fmt.Println("--------------begin---------------")
	fmt.Println(arr)

	arrLen := len(arr)

	if arrLen <= 1 {
		return
	}

	randNum := rand.Intn(arrLen - 1)
	fmt.Println("randNum: ", randNum, " ", arr[randNum], " ", arr[0])
	arr[randNum], arr[0] = arr[0], arr[randNum]
	mid := arr[0]

	head, tail := 0, arrLen-1

	for {
		for head <= tail && arr[head] < mid {
			head++
		}

		for tail > head && arr[tail] >= mid {
			tail--
		}

		if head > tail {
			break
		}

		arr[head], arr[tail] = arr[tail], arr[head]
		head++
		tail--
	}
	fmt.Println(arr)
	fmt.Println("head: ", head, " ", arr[head])
	fmt.Println("--------------end---------------")

	QSortTwoWay(arr[:head])
	QSortTwoWay(arr[head+1:])
}
