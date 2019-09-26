package learn

import "fmt"

// ListNode single linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode xxxx
func NewListNode(arr []int) (head *ListNode) {
	l := &ListNode{}

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			l = &ListNode{
				Val:  arr[i],
				Next: nil,
			}

			head = l
			continue
		}

		l.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		l = l.Next
	}

	return
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
