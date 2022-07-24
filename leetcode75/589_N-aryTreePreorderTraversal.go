package leetcode75

type Node struct {
	Val      int
	Children []*Node
}

func dfs(node *Node, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	for _, n := range node.Children {
		dfs(n, res)
	}
}

func preorder(root *Node) []int {
	res := []int{}
	dfs(root, &res)
	return res
}
