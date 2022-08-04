package recursion

func Merge_sorted_list(a *listNode, b *listNode) *listNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	if a.Val < b.Val {
		a.Next = Merge_sorted_list(a.Next, b)
		return a
	} else {
		b.Next = Merge_sorted_list(a, b.Next)
		return b
	}
}
