package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	ret := &ListNode{}
	head := ret
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		ret.Val = sum % 10
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		ret.Next = &ListNode{}
		ret = ret.Next
	}
	return head
}

func AddTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := l1
	for {
		l1.Val += l2.Val + carry
		carry = l1.Val / 10
		l1.Val %= 10
		if l1.Next == nil && l2.Next == nil {
			if carry > 0 {
				l1.Next = &ListNode{Val: carry}
			}
			break
		}
		l1 = l1.Next
		if l1 == nil && l2.Next != nil {
			l1 = &ListNode{}
		}
		l2 = l2.Next
		if l2 == nil {
			l2 = &ListNode{}
		}
	}
	return head
}
