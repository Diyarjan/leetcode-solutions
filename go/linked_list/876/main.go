// 876. Middle of the Linked List

package main

import "fmt"

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {

	// Создание связанного списка: 1 → 2 → 3 → 4 → 5 → 6
	head := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5,
						&ListNode{6, nil}}}}}}

	mid := middleNode(head)

	for mid != nil {
		fmt.Print(mid.Val, " ")
		mid = mid.Next
	}

	fmt.Println()
}
