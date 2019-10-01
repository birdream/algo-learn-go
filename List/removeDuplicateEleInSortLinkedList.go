package List

// 82 delete all duplicate element from a sorted linked list
// 1-1-2-3-4-4 --> 2-3
func removedDuplicatedFromSortedLinkedList(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}

	pre := dummyNode
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			delNode1 := cur
			delNode2 := cur.Next
			pre.Next = delNode2.Next
			cur = pre.Next

			delNode1.Next = nil
			delNode2.Next = nil
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}

	return dummyNode.Next
}
