package List

import (
	"reflect"
	"testing"
)

func TestPartitionList(t *testing.T) {
	t.Run("testPartitionList1 should work", testPartitionList1)
}

func testPartitionList1(t *testing.T) {
	signalList := NewListNode([]int{1, 5, 2, 5, 6, 3, 7, 8, 0})

	head := partitionList(signalList, 4)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{1, 2, 3, 0, 5, 5, 6, 7, 8}) {
		t.Errorf("Error %d", res)
	}
}
