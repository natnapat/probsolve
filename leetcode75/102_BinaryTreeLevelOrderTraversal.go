package leetcode75

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		var lvl []int
		qlen := len(q)
		for i := 0; i < qlen; i++ {
			node := q[0]
			q = q[1:]
			lvl = append(lvl, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, lvl)
	}
	return res
}
