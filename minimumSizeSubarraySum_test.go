package learn

import (
	"reflect"
	"testing"
)

func TestMethod(t *testing.T) {
	t.Run("should get correct result", testMethod1)
	t.Run("should get correct result", testMethod2)
	t.Run("should get correct result", testMethod3)
	t.Run("should get correct result", testMethod4)
	t.Run("should get correct result", testMethod5)
	t.Run("should get correct result", testMethod6)
}

func testMethod1(t *testing.T) {
	array := []int{1, 2, 3, 2, 5, 4, 2, 1}
	res := method(array, 6)
	if !reflect.DeepEqual(res, []int{2, 5}) {
		t.Errorf("Expect %d, get %d \n", []int{2, 5}, res)
	}
}

func testMethod2(t *testing.T) {
	array := []int{1, 2, 3, 2, 5, 4, 2, 1}
	res := method(array, 100)
	if res != 0 {
		t.Errorf("Expect %d, get %d \n", 0, res)
	}
}

func testMethod3(t *testing.T) {
	array := []int{1, 2, 1, 2, 3, 1, 3, 2, 3, 3}
	res := method(array, 6)
	if !reflect.DeepEqual(res, []int{3, 3}) {
		t.Errorf("Expect %d, get %d \n", []int{3, 3}, res)
	}
}

func testMethod4(t *testing.T) {
	array := []int{0, 0, 0, 0, 0, 0, 0, 0, 3, 3}
	res := method(array, 6)
	if !reflect.DeepEqual(res, []int{3, 3}) {
		t.Errorf("Expect %d, get %d \n", []int{3, 3}, res)
	}
}

func testMethod5(t *testing.T) {
	array := []int{0, 0, 0, 0, 0, 0, 0, 0, 3, 3, 6}
	res := method(array, 6)
	if !reflect.DeepEqual(res, []int{6}) {
		t.Errorf("Expect %d, get %d \n", []int{6}, res)
	}
}

func testMethod6(t *testing.T) {
	array := []int{0, 0, 0, 0, 0, 0, 0, 0, 3, 3, 6, 9, 0}
	res := method(array, 6)
	if !reflect.DeepEqual(res, []int{6}) {
		t.Errorf("Expect %d, get %d \n", []int{9}, res)
	}
}
