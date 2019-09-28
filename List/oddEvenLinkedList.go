package List

// 328 将链表中奇数项排前面 偶数项排后面
func oddEvenLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd := head
	even := odd.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}
