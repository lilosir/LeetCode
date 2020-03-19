package problem700

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if root.Val > val {
			return searchBST(root.Left, val)
		}
		return searchBST(root.Right, val)
	}

	return nil
}
