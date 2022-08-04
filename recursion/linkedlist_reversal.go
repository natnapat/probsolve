package recursion

type listNode struct {
	Val  int
	Next *listNode
}

func Reverse_linkedlist(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := Reverse_linkedlist(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
