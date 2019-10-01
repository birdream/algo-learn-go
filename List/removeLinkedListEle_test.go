package List

import (
	"reflect"
	"testing"
)

func TestRemoveLinkedListEle(t *testing.T) {
	t.Run("testRemoveLinkedListEle1 should work", testRemoveLinkedListEle1)
	t.Run("testRemoveLinkedListEle2 should work", testRemoveLinkedListEle2)
	t.Run("testRemoveLinkedListEle3 should work", testRemoveLinkedListEle3)
	t.Run("testRemoveLinkedListEle4 should work", testRemoveLinkedListEle4)
}

func testRemoveLinkedListEle1(t *testing.T) {
	l := NewListNode([]int{1, 1, 2, 2, 4, 4, 5, 5, 5, 5, 6, 7, 7, 8})
	res := removeLinkedListEle(l, 1)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{2, 2, 4, 4, 5, 5, 5, 5, 6, 7, 7, 8}) {
		t.Errorf("Error with %d", resArr)
	}
}

func testRemoveLinkedListEle2(t *testing.T) {
	l := NewListNode([]int{1, 1, 2, 2, 4, 4, 5, 5, 5, 5, 6, 7, 7, 8})
	res := removeLinkedListEle(l, 5)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{1, 1, 2, 2, 4, 4, 6, 7, 7, 8}) {
		t.Errorf("Error with %d", resArr)
	}
}

func testRemoveLinkedListEle3(t *testing.T) {
	l := NewListNode([]int{1})
	res := removeLinkedListEle(l, 1)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{}) {
		t.Errorf("Error with %d", resArr)
	}
}

func testRemoveLinkedListEle4(t *testing.T) {
	l := NewListNode([]int{1, 1, 2})
	res := removeLinkedListEle(l, 1)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{2}) {
		t.Errorf("Error with %d", resArr)
	}
}
