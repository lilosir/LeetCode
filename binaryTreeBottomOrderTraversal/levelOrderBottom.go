package problem107

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	return nil
}

func visitCurrentLevel(roots []*TreeNode) []int {
	result := []int{}
	for _, node := range roots {
		result = append(result, node.Val)
	}
	return result
}
