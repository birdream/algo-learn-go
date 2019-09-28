package List

import (
	"reflect"
	"testing"
)

func TestMethodRLLWT(t *testing.T) {
	t.Run("testmethodRLLWT1 should work", testmethodRLLWT1)
	t.Run("testmethodRLLWT2 should work", testmethodRLLWT2)
	t.Run("testmethodRLLWT3 should work", testmethodRLLWT3)
}

func testmethodRLLWT1(t *testing.T) {
	signalList := NewListNode([]int{5, 6, 7, 8, 0})

	head := methodRLLWT(signalList, 2, 3)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{5, 7, 6, 8, 0}) {
		t.Errorf("Error %d", res)
	}
}

func testmethodRLLWT2(t *testing.T) {
	signalList := NewListNode([]int{5, 6, 7, 8, 0})

	head := methodRLLWT(signalList, 1, 3)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{7, 6, 5, 8, 0}) {
		t.Errorf("Error %d", res)
	}
}

func testmethodRLLWT3(t *testing.T) {
	signalList := NewListNode([]int{5, 6, 7, 8, 0})

	head := methodRLLWT(signalList, 1, 5)
	res := GetListValByArr(head)

	if !reflect.DeepEqual(res, []int{0, 8, 7, 6, 5}) {
		t.Errorf("Error %d", res)
	}
}
