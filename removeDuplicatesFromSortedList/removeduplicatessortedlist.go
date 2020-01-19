package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func main() {
	list1 := ListNode{
		Val:  1,
		Next: nil,
	}
	list2 := ListNode{
		Val:  1,
		Next: &list1,
	}
	list3 := ListNode{
		Val:  2,
		Next: &list2,
	}

	head := ListNode{
		Val:  2,
		Next: &list3,
	}

	result := deleteDuplicates(&head)
	fmt.Printf("%v", result)
}
