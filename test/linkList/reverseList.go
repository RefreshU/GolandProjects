package base_demo

/*
type ListNode struct {
	Val int
	Next *ListNode
}
*/

func reverseList(head *ListNode) *ListNode{
	var prev *ListNode = nil
	var curr = head
	if curr != nil {
		var next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}