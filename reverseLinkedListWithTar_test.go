package learn

import (
	"reflect"
	"testing"
)

func TestMethodRLLWT(t *testing.T) {
	t.Run("shoul work", testmethodRLLWT1)
}

func testmethodRLLWT1(t *testing.T) {
	source := []int{5, 6, 7, 8, 0}
	signalList := NewListNode(11)
	signalList.addList(source)

	head := methodRLLWT(signalList, 2, 3)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{11, 6, 5, 7, 8, 0}) {
		t.Errorf("Error %d", res)
	}
}
