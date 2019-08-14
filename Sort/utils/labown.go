package labown

import (
	"fmt"
	"math/rand"
	"time"
)

// CreateRandomArr dddd
func CreateRandomArr(len int, p bool) []int {
	res := make([]int, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		res[i] = rand.Intn(150) - 50
	}

	if p {
		fmt.Println(res)
	}
	return res
}

// IsSorted fsdsf
func IsSorted(arr []int, p bool) {
	if p {
		fmt.Println(arr)
	}

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
	if p {
		fmt.Print("----------------------\n")
	}
}
