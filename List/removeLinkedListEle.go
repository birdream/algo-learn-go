package List

func removeLinkedListEle(head *ListNode, x int) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}
	cur := dummyNode

	for cur.Next != nil {
		if cur.Next.Val == x {
			delNode := cur.Next
			cur.Next = delNode.Next
			delNode.Next = nil

		} else {
			cur = cur.Next
		}
	}

	return dummyNode.Next
}
