package recursion

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func Insert_node_binary_tree(head *treeNode, data int) *treeNode {
	if head == nil {
		head.Val = data
		return head
	}
	if data > head.Val {
		head.Right = Insert_node_binary_tree(head.Right, data)
	} else {
		head.Left = Insert_node_binary_tree(head.Left, data)
	}
	return head
}
