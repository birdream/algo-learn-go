package learn

import (
	"fmt"
	"testing"
)

func TestMethodRLL(t *testing.T) {
	t.Run("shoul work", testmethodRLL1)
}

func PrintList(l *ListNode) {
	fmt.Println("==============")
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
	fmt.Println("==============")
}

func testmethodRLL1(t *testing.T) {
	signalList := NewListNode(5)
	signalList.Next = NewListNode(6)
	signalList.Next.Next = NewListNode(7)
	signalList.Next.Next.Next = NewListNode(8)
	signalList.Next.Next.Next.Next = NewListNode(0)

	// PrintList(signalList)

	methodRLL(signalList)

	// PrintList(head)
}
