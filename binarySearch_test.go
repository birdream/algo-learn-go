package learn

import (
	"math/rand"
	"testing"
)

func buildSortedArr(length, gap int) []int {
	arr := make([]int, 0)

	for i := 0; i <= length; i++ {
		arr = append(arr, i+gap)
	}

	return arr
}

func TestBinarySearch(t *testing.T) {
	total := 100
	sortedArray := buildSortedArr(total, 0)
	// fmt.Println(sortedArray)
	t.Run("should get correct index of binarySearch method", func(t *testing.T) {
		target := rand.Intn(total)
		expectRes := target

		res := binarySearch(sortedArray, target)
		if res != expectRes {
			t.Errorf("Result is not correct, get %d, expect %d", res, expectRes)
		} else {
			// fmt.Printf("Pass the test expect %d \n", res)
		}
	})
}
