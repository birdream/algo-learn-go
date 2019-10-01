package List

func removeLinkedListEle(head *ListNode, x int) *ListNode {

	for head != nil && head.Val == x {
		delNode := head
		head = head.Next
		delNode.Next = nil
	}

	if head == nil {
		return head
	}

	cur := head

	for cur.Next != nil {
		if cur.Next.Val == x {
			delNode := cur.Next
			cur.Next = delNode.Next
			delNode.Next = nil

		} else {
			cur = cur.Next
		}
	}

	return head
}
