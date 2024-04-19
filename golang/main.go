package main

import "golang/leetcode"

func main() {

	l1 := &leetcode.ListNode{
		Val: 2,
		Next: &leetcode.ListNode{
			Val: 4,
			Next: &leetcode.ListNode{
				Val: 9,
			},
		},
	}

	l2 := &leetcode.ListNode{
		Val: 5,
		Next: &leetcode.ListNode{
			Val: 6,
			Next: &leetcode.ListNode{
				Val:  4,
				Next: &leetcode.ListNode{Val: 9},
			},
		},
	}
	numbers := leetcode.AddTwoNumbers1(l1, l2)
	for numbers != nil {
		print(numbers.Val)
		numbers = numbers.Next
	}
}
