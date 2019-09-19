package learn

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("test two sum should work", testTwoSum)
}

func testTwoSum(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	res1 := twoSum(arr, 6)
	res2 := twoSum(arr, 5)
	res3 := twoSum(arr, 9)
	res4 := twoSum(arr, 21)
	if !reflect.DeepEqual(res1, []int{0, 6}) {
		t.Errorf("Result not correct 1 %d", res1)
	}

	if !reflect.DeepEqual(res2, []int{0, 5}) {
		t.Errorf("Result not correct 2 %d", res2)
	}

	if !reflect.DeepEqual(res3, []int{0, 9}) {
		t.Errorf("Result not correct 3 %d", res3)
	}

	if !reflect.DeepEqual(res4, []int{8, 13}) {
		t.Errorf("Result not correct 4 %d", res4)
	}
}
