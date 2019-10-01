package List

import (
	"reflect"
	"testing"
)

func TestRemovedDuplicatedFromSortedLinkedList(t *testing.T) {
	t.Run("testRemovedDuplicatedFromSortedLinkedList1 should work", testRemovedDuplicatedFromSortedLinkedList1)
	t.Run("testRemovedDuplicatedFromSortedLinkedList2 should work", testRemovedDuplicatedFromSortedLinkedList2)
	t.Run("testRemovedDuplicatedFromSortedLinkedList3 should work", testRemovedDuplicatedFromSortedLinkedList3)
	t.Run("testRemovedDuplicatedFromSortedLinkedList4 should work", testRemovedDuplicatedFromSortedLinkedList4)
	t.Run("testRemovedDuplicatedFromSortedLinkedList5 should work", testRemovedDuplicatedFromSortedLinkedList5)
}

func testRemovedDuplicatedFromSortedLinkedList1(t *testing.T) {
	l := NewListNode([]int{1, 1, 2, 2, 4, 4, 5, 5, 5, 5, 6, 7, 7, 8})
	res := removedDuplicatedFromSortedLinkedList(l)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{6, 8}) {
		t.Errorf("Error with %d", resArr)
	}
}

func testRemovedDuplicatedFromSortedLinkedList2(t *testing.T) {
	l := NewListNode([]int{1, 2, 3, 4, 5, 6, 6})
	res := removedDuplicatedFromSortedLinkedList(l)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Error with %d", resArr)
	}
}

func testRemovedDuplicatedFromSortedLinkedList3(t *testing.T) {
	l := NewListNode([]int{1})
	res := removedDuplicatedFromSortedLinkedList(l)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{1}) {
		t.Errorf("Error with %d", resArr)
	}
}

func testRemovedDuplicatedFromSortedLinkedList4(t *testing.T) {
	l := NewListNode([]int{1, 2, 3})
	res := removedDuplicatedFromSortedLinkedList(l)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{1, 2, 3}) {
		t.Errorf("Error with %d", resArr)
	}
}

func testRemovedDuplicatedFromSortedLinkedList5(t *testing.T) {
	l := NewListNode([]int{1, 1, 2, 2, 3, 3})
	res := removedDuplicatedFromSortedLinkedList(l)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{}) {
		t.Errorf("Error with %d", resArr)
	}
}
