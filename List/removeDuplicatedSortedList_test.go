package List

import (
	"reflect"
	"testing"
)

func TestMethodRDSL(t *testing.T) {
	t.Run("testMethodRDSL1 should work", testMethodRDSL1)
	t.Run("testMethodRDSL2 should work", testMethodRDSL2)
	t.Run("testMethodRDSL3 should work", testMethodRDSL3)
}

func testMethodRDSL1(t *testing.T) {
	l := NewListNode([]int{1, 1, 2, 2, 4, 4, 5, 5, 5, 5, 6, 7, 7, 8})
	res := methodRDSL(l)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{1, 2, 4, 5, 6, 7, 8}) {
		t.Errorf("Error with %d", resArr)
	}
}

func testMethodRDSL2(t *testing.T) {
	l := NewListNode([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	res := methodRDSL(l)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{1}) {
		t.Errorf("Error with %d", resArr)
	}
}

func testMethodRDSL3(t *testing.T) {
	l := NewListNode([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2})
	res := methodRDSL(l)

	resArr := GetListValByArr(res)

	if !reflect.DeepEqual(resArr, []int{1, 2}) {
		t.Errorf("Error with %d", resArr)
	}
}
