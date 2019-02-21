package recursive


  //Definition for singly-linked list.
  type ListNode struct {
      Val int
     Next *ListNode
  }

func swapPairs(head *ListNode) *ListNode {
	if head == nil||head.Next == nil{
		return head
	}
	tmp := head.Next
	head.Next = swapPairs(head.Next.Next)
	tmp.Next = head
	return tmp
}
