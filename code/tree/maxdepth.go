package tree
func maxDepth(root *TreeNode) int {
	depth := 0

	depth = findDepth(root)

	return depth
}

func findDepth(root *TreeNode) int   {
	if root != nil {
		leftDepth := findDepth(root.Left)
		rightDepth := findDepth(root.Right)
		if leftDepth > rightDepth {
			return 1+leftDepth
		}
		return 1+rightDepth
	}
	return 0

}