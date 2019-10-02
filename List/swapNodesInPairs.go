package List

// 24 swap nodes in pairs
// 1-2-3-4-nil --> 2-1-4-3-nil
// 1-2-3-nil --> 2-1-3-nil
func swapNodesInPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: nil,
	}
	pre := dummyNode
	cur := dummyNode.Next

	for cur != nil && cur.Next != nil {
		pre.Next = cur.Next
		cur.Next = cur.Next.Next

		pre = pre.Next
		cur = cur.Next
	}

	return dummyNode.Next
}
