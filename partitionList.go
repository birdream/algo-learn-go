package learn

// 86 给定链表和数x 大于等于x的在右边 小于x在左边 返回新链表
func partitionList(head *ListNode, x int) (h *ListNode) {
	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}
	preNode := dummyNode
	leftNode := dummyNode
	cur := head

	for cur != nil {
		if preNode == leftNode {
			if cur.Val < x {
				leftNode = leftNode.Next
			}

			preNode = cur
			cur = cur.Next
		} else {
			if cur.Val >= x {
				preNode = cur
				cur = cur.Next
			} else {
				preNode.Next = cur.Next
				cur.Next = leftNode.Next
				leftNode.Next = cur

				leftNode = leftNode.Next
				cur = preNode.Next
			}
		}

	}

	return dummyNode.Next
}
