package List

// 给定区间内反转链表
// M N node初始都是在头节点 preNode 在头节点的前一个节点 dummyNode
// 先一层 m n 循环找到 m n pre node的位置
// pre.Next指向m下一个节点
// 将m 一直放到n 节点后面 直到n 移到m的位置
// 将m 的Next指向 n的Next
// 返回 head
func methodRLLWT(head *ListNode, m, n int) *ListNode {
	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}

	nodeM := head
	preNode := dummyNode // m 的前一个节点
	nodeN := head

	for i := 1; i < n; i++ {
		if i < m {
			preNode = nodeM
			nodeM = nodeM.Next
		}

		nodeN = nodeN.Next
	}

	for nodeM != nodeN {
		preNode.Next = nodeM.Next
		nodeM.Next = nodeN.Next
		nodeN.Next = nodeM

		nodeM = preNode.Next
	}

	return dummyNode.Next
}
