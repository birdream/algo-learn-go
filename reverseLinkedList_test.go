package learn

import (
	"reflect"
	"testing"
)

func TestMethodRLL(t *testing.T) {
	t.Run("shoul work", testmethodRLL1)
}

func testmethodRLL1(t *testing.T) {
	signalList := NewListNode([]int{5, 6, 7, 8, 0})
	head := methodRLL(signalList)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{0, 8, 7, 6, 5}) {
		t.Errorf("Error %d", res)
	}
}
