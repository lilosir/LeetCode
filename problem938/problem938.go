package problem938

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root != nil {
		if root.Val < L {
			return rangeSumBST(root.Right, L, R)
		}
		if root.Val > R {
			return rangeSumBST(root.Left, L, R)
		}
		return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
	}
	return 0
}