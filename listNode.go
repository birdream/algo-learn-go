package learn

import "fmt"

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

func (l *ListNode) addList(arr []int) {
	for {
		if l.Next == nil {
			for i := 0; i < len(arr); i++ {
				l.Next = &ListNode{
					Val:  arr[i],
					Next: nil,
				}

				l = l.Next
			}

			break
		}

		l = l.Next
	}
}

func PrintList(l *ListNode) {
	fmt.Println("==============")
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
	fmt.Println("==============")
}

func GetListValByArr(l *ListNode) []int {
	res := make([]int, 0)

	for l != nil {
		res = append(res, l.Val)

		l = l.Next
	}

	return res
}
