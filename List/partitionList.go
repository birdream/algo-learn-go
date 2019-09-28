package List

// 86 给定链表和数x 大于等于x的在右边 小于x在左边 返回新链表

// left为分割点 将比x小的移到left的左边 比x大的不需要动 将边界left往右移
// 当遇到比x小的时候
//    pre.next指向 cur next
//    cur next 指向 left next
//    left next 指向 cur
//    left cur 右移一位继续循环
// 当遇到比x大时候
//    pre 和 left 右移一位
// 注意 因为初始时候 pre 和 left 都指向dummy 按上面的移动是会死循环的
// 所以 pre==left 时候（循环的第一次进入时）要单独判断
//    如果 cur 比 x 小: left 先右移一位
//    pre 指向 cur， cur 右移一位
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
