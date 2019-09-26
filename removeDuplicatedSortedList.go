package learn

// 83 去除有序链表重复项
func methodRDSL(head *ListNode) (h *ListNode) {
	h = head

	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return
}
