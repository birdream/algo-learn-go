package List

import (
	"reflect"
	"testing"
)

func TestMergeTwoSortedList(t *testing.T) {
	t.Run("testMergeTwoSortedList1 should work", testMergeTwoSortedList1)
	t.Run("testMergeTwoSortedList2 should work", testMergeTwoSortedList2)
}

func testMergeTwoSortedList1(t *testing.T) {
	l1 := NewListNode([]int{1, 2, 3})
	l2 := NewListNode([]int{4, 5, 6, 7, 8})

	head := mergeTwoSortedList(l1, l2)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("Error %d", res)
	}
}

func testMergeTwoSortedList2(t *testing.T) {
	l1 := NewListNode([]int{1, 2, 6, 9, 10})
	l2 := NewListNode([]int{0, 4, 5, 6, 7, 8})

	head := mergeTwoSortedList(l1, l2)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{0, 1, 2, 4, 5, 6, 6, 7, 8, 9, 10}) {
		t.Errorf("Error %d", res)
	}
}
