package tree



 //Definition for a binary tree node.
  type TreeNode struct {
      Val int
     Left *TreeNode
     Right *TreeNode
  }


func preorderTraversal(root *TreeNode) []int {

	result :=[]int{}

	find(root,&result)
	return result
}

func find(root *TreeNode, result *[]int){

	if root==nil{
		return
	}
	*result=append(*result,root.Val)

	find(root.Left,result)

	find(root.Right,result)
}