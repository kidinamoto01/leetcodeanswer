package tree

func inorderTraversal(root *TreeNode) []int {
	result :=[]int{}

	find2(root,&result)

	return result
}

func find2(root *TreeNode, result *[]int){

	if root!=nil{
		find2(root.Left,result)
		*result=append(*result,root.Val)
		find2(root.Right,result)
	}
}