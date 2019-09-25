package learn

import (
	"reflect"
	"testing"
)

func TestMethodRLL(t *testing.T) {
	t.Run("shoul work", testmethodRLL1)
}

func testmethodRLL1(t *testing.T) {
	source := []int{5, 6, 7, 8, 0}
	signalList := NewListNode(11)
	signalList.addList(source)

	head := methodRLL(signalList)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{0, 8, 7, 6, 5, 11}) {
		t.Errorf("Error %d", res)
	}
}
