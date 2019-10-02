package List

import (
	"reflect"
	"testing"
)

func TestSwapNodesInPairs(t *testing.T) {
	t.Run("testSwapNodesInPairs1 should work", testSwapNodesInPairs1)
	t.Run("testSwapNodesInPairs2 should work", testSwapNodesInPairs2)
}

func testSwapNodesInPairs1(t *testing.T) {
	signalList := NewListNode([]int{5, 6, 7, 8, 0})

	head := swapNodesInPairs(signalList)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{6, 5, 8, 7, 0}) {
		t.Errorf("Error %d", res)
	}
}

func testSwapNodesInPairs2(t *testing.T) {
	signalList := NewListNode([]int{5, 6, 7, 8})

	head := swapNodesInPairs(signalList)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{6, 5, 8, 7}) {
		t.Errorf("Error %d", res)
	}
}
