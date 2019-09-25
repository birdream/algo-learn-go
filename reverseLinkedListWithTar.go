package learn

// 给定区间内反转链表
func methodRLLWT(head *ListNode, m, n int) *ListNode {
	var pre *ListNode
	cur := head

	for i := 1; i < m; i++ {
		pre = cur
		cur = cur.Next
	}

	for i := 0; i < n-m; i++ {
		temp := cur.Next

		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}

	return head
}
