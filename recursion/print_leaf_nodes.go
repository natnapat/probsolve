package recursion

import (
	"fmt"
	"strconv"
)

func Print_leaf_nodes(root *treeNode) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		fmt.Print(strconv.Itoa(root.Val) + ", ")
		return
	}

	if root.Left != nil {
		Print_leaf_nodes(root.Left)
	}
	if root.Right != nil {
		Print_leaf_nodes(root.Right)
	}
}
