package tree

func postorderTraversal(root *TreeNode) []int {
	result :=[]int{}

	find3(root,&result)

	return result
}

func find3(root *TreeNode, result *[]int){

	if root!=nil{
		find3(root.Left,result)
		find3(root.Right,result)
		*result=append(*result,root.Val)
	}
}