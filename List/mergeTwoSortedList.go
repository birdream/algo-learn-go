package List

// 21 merge two sorted linked list
func mergeTwoSortedList(l1, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: nil,
	}
	cur := dummyNode

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = &ListNode{
				Val:  l2.Val,
				Next: nil,
			}

			l2 = l2.Next
		} else {
			cur.Next = &ListNode{
				Val:  l1.Val,
				Next: nil,
			}

			l1 = l1.Next
		}

		cur = cur.Next
	}

	if l1 == nil {
		cur.Next = l2
	}

	if l2 == nil {
		cur.Next = l1
	}

	return dummyNode.Next
}
