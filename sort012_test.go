package learn

import (
	"reflect"
	"testing"
)

func TestSort012(t *testing.T) {
	t.Run("sort 0 1 2 should work", testSort012)
	t.Run("sort by 3 way should work", testSortColorsBy3Way)
}

func testSort012(t *testing.T) {
	arr := []int{0, 0, 0, 1, 2, 1, 0, 2, 0, 2, 1, 0, 2, 2, 1, 1}

	res := sortColors(arr)
	if len(arr) != len(res) {
		t.Errorf("Len should be same %d %d", len(arr), len(res))
	}

	if !reflect.DeepEqual(res, []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}) {
		t.Errorf("Result not correct")
	}
}

func testSortColorsBy3Way(t *testing.T) {
	arr := []int{0, 0, 0, 1, 2, 1, 0, 2, 0, 2, 1, 0, 2, 2, 1, 1}

	res := sortColors(arr)
	if len(arr) != len(res) {
		t.Errorf("Len should be same %d %d", len(arr), len(res))
	}

	if !reflect.DeepEqual(res, []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}) {
		t.Errorf("Result not correct")
	}
}
