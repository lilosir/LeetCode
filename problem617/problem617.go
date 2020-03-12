package problem617

import (
	"github.com/lilosir/LeetCode/kit"
)

func mergeTrees(t1 *kit.TreeNode, t2 *kit.TreeNode) *kit.TreeNode {
	result := &kit.TreeNode{}

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	result.Val = t1.Val + t2.Val
	result.Left = mergeTrees(t1.Left, t2.Left)
	result.Right = mergeTrees(t1.Right, t2.Right)

	return result
}
