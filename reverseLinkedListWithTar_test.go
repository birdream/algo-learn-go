package learn

import (
	"testing"
)

func TestMethodRLLWT(t *testing.T) {
	t.Run("shoul work", testmethodRLLWT1)
}

func testmethodRLLWT1(t *testing.T) {
	signalList := NewListNode(5)
	signalList.Next = NewListNode(6)
	signalList.Next.Next = NewListNode(7)
	signalList.Next.Next.Next = NewListNode(8)
	signalList.Next.Next.Next.Next = NewListNode(0)

	PrintList(signalList)

	head := methodRLLWT(signalList, 2, 3)

	PrintList(head)
}
