package learn

// ListNode 链表反转
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode xxxx
func NewListNode(v int) *ListNode {
	return &ListNode{
		Val:  v,
		Next: nil,
	}
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
