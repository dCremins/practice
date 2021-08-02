package main


type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:4,
			Next: &ListNode{
				Val:3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:6,
			Next: &ListNode{
				Val:4,
			},
		},
	}
	addTwoNumbers(l1, l2)
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	sum := 0
	node := &ListNode{}
	result := node
    for l1 != nil || l2 != nil || carry != 0 {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		} else {
			carry = 0
		}
		node.Next = &ListNode{Val:sum}
		node = node.Next
	}
	return result.Next
}