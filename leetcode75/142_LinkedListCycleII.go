package leetcode75

func detectCycle(head *ListNode) *ListNode {
	var fast, slow *ListNode
	fast, slow = head, head
	cycle := false
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			cycle = true
			break
		}
	}
	if cycle == true {
		slow = head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		return slow
	}
	return nil
}
