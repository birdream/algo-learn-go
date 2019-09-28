package List

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("testAddTwoNumbers1 should work", testAddTwoNumbers1)
	t.Run("testAddTwoNumbers2 should work", testAddTwoNumbers2)
}

func testAddTwoNumbers1(t *testing.T) {
	// 7888
	// 333
	l1 := NewListNode([]int{7, 8, 8, 8})
	l2 := NewListNode([]int{3, 3, 3})

	resHead := addTwoNumbers(l1, l2)
	res := GetListValByArr(resHead)

	if !reflect.DeepEqual(res, []int{8, 2, 2, 1}) {
		t.Errorf("Error %d", res)
	}
}

func testAddTwoNumbers2(t *testing.T) {
	// 7888
	// 2
	l1 := NewListNode([]int{7, 8, 8, 8})
	l2 := NewListNode([]int{2})

	resHead := addTwoNumbers(l1, l2)
	res := GetListValByArr(resHead)

	if !reflect.DeepEqual(res, []int{7, 8, 9, 0}) {
		t.Errorf("Error %d", res)
	}
}
