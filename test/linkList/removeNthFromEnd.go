package base_demo

/*
type ListNode struct {
	Val int
	Next *ListNode
}
*/
func RemoveNthFromEnd(head *ListNode, n int)*ListNode {
	var dummy *ListNode = &ListNode{Next: head}
	slow, fast := dummy,dummy
	for i := 0; i < n ; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}