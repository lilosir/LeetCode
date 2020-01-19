package problem83

// ListNode denotes link list
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicates(head *ListNode) *ListNode {
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
