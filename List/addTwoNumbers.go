package List

func Pop(a []int) (int, []int) {
	return a[len(a)-1], a[:len(a)-1]
}

// 445 给出两个非负链表表示两个非负整数 返回计算结果的链表
// 342+465 = 807
/// 3-4-2  4-6-5 返回 8-0-7
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lStack := make([]int, 0)
	rStack := make([]int, 0)

	for l1 != nil {
		lStack = append(lStack, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		rStack = append(rStack, l2.Val)
		l2 = l2.Next
	}

	total := 0
	currNode := &ListNode{
		Val:  0,
		Next: nil,
	}

	for len(lStack) != 0 || len(rStack) != 0 {
		if len(lStack) != 0 {
			v, newStack := Pop(lStack)
			lStack = newStack
			total += v
		}
		if len(rStack) != 0 {
			v, newStack := Pop(rStack)
			rStack = newStack
			total += v
		}

		currNode.Val = total % 10
		nextNode := NewNode(total / 10)
		// 因为是逆序的
		nextNode.Next = currNode
		currNode = nextNode

		total /= 10
	}

	if currNode.Val == 0 {
		return currNode.Next
	} else {
		return currNode
	}
}
