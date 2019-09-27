package learn

import (
	"reflect"
	"testing"
)

func TestOddEvenLinkedList(t *testing.T) {
	t.Run("testOddEvenLinkedList should work", testOddEvenLinkedList)
}

func testOddEvenLinkedList(t *testing.T) {
	signalList := NewListNode([]int{1, 2, 3, 4, 5, 6, 7, 8})

	head := oddEvenLinkedList(signalList)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{1, 3, 5, 7, 2, 4, 6, 8}) {
		t.Errorf("Error %d", res)
	}
}
