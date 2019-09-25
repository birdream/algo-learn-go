package learn

// 反转链表
func methodRLL(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next

		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
