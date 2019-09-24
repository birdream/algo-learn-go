package learn

// ListNode 链表反转
type ListNode struct {
	Val  int
	Next *ListNode
}

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
