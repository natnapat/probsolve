package leetcode75

func middleNode(head *ListNode) *ListNode {
	var curr *ListNode
	curr = head
	count := 0
	for curr != nil {
		curr = curr.Next
		count++
	}
	mid := (count / 2)
	curr = head
	for i := 0; i < mid; i++ {
		curr = curr.Next
	}
	return curr
}

// fast and slow pointer approach
// func middleNode(head *ListNode) *ListNode {
//     var slow, fast *ListNode
//     slow, fast = head, head
//     for fast != nil && fast.Next != nil {
//         fast = fast.Next.Next
//         slow = slow.Next
//     }
//     return slow
// }
