package main

import (
	"fmt"

	"github.com/isdamir/gotype"
)

// Reverse1 原地逆序list by book
func Reverse1(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var pre *gotype.LNode // pre node
	var cur *gotype.LNode // current node
	next := node.Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}

	node.Next = pre
}

// Reverse1_1 原地逆序list by norman
func Reverse1_1(head *gotype.LNode) (mewHead *gotype.LNode) {
	if head == nil || head.Next == nil {
		return
	}

	var pre *gotype.LNode
	var next *gotype.LNode

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return &gotype.LNode{Next: pre}
}

// ReverseInsert 插入法进行逆序 by norman by book
func ReverseInsert(head *gotype.LNode) {
	//     v n
	// H 1 2 3 4 5
	//       v n
	// H 2 1 3 4 5
	//         v n
	// H 3 2 1 4 5
	//           v n
	// H 4 3 2 1 5
	// H 5 4 3 2 1
	if head == nil || head.Next == nil {
		return
	}

	var next *gotype.LNode
	cur := head.Next.Next

	head.Next.Next = nil // First time need to set to nil or it will loop infinti

	for cur != nil {
		next = cur.Next
		cur.Next = head.Next
		head.Next = cur
		cur = next
	}
}

func main() {
	head := &gotype.LNode{}
	fmt.Print("Reverse on the ground\n")
	gotype.CreateNode(head, 8)
	gotype.PrintNode("Before: ", head)
	Reverse1(head)
	gotype.PrintNode("After Reverse1: ", head)
	ReverseInsert(head)
	gotype.PrintNode("After ReverseInsert: ", head)
}
