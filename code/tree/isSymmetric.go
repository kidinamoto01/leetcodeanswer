package tree

func isSymmetric(root *TreeNode) bool {
	if root != nil {
		return true
	}

	return compare(root.Left,root.Right)
}

func compare(left *TreeNode,right *TreeNode) bool{
	if left == nil && right == nil{
		return true
	}
	if left==nil||right==nil{
		return false
	}
	if left.Val==right.Val{
		return compare(left.Left, right.Right) && compare(left.Right, right.Left)
	}

	return false
}


