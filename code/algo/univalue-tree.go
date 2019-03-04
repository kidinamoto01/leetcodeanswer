package algo
/**
 * Definition for a binary tree node.
*/

// Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }



func isUnivalTree(root *TreeNode) bool {

	if root ==nil {
		return true
	}else{
		val := root.Val
		return search(val, root)
	}

}

func search(i int,parent *TreeNode) bool {

	if parent == nil {
		return true
	}
	if parent.Val != i{
		return false
	}else{
		return search(i,parent.Left) && search(i,parent.Right)

	}

}
