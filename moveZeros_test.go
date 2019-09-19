package algo

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	t.Run("move zeros by K should work", testMoveZerosByK)
	t.Run("move zeros by switch should work", testMoveZerosBySwitch)
}

func testMoveZerosByK(t *testing.T) {
	arr := []int{0, 0, 1, 2, 3, 4, 0, 0, 0, 6, 9, 0, 1, 8, 0}
	res := moveZerosByK(arr)

	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 6, 9, 1, 8, 0, 0, 0, 0, 0, 0, 0}) {
		t.Error("testMoveZerosByK Fail")
	}
}

func testMoveZerosBySwitch(t *testing.T) {
	arr := []int{0, 0, 1, 2, 3, 4, 0, 0, 0, 6, 9, 0, 1, 8, 0}
	res := moveZerosBySwitch(arr)

	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 6, 9, 1, 8, 0, 0, 0, 0, 0, 0, 0}) {
		t.Error("testMoveZerosBySwitch Fail")
	}
}
