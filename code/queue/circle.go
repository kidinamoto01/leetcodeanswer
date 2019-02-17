package queue

type MyCircularQueue struct {
	nums []int
	front, rear, size int

}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	q := MyCircularQueue{}
	q.nums = make([]int,k+1)
	q.front = 0
	q.rear = 0
	q.size = k+1
	return q
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}else{
		this.nums[this.rear] = value
		this.rear = (this.rear+1)%this.size

		return true
	}


}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty(){
		return false
	}

	this.nums[this.front] = 0
	this.front=(this.front+1+this.size)%this.size
	return true

}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty(){
		return -1
	}else{
		return this.nums[this.front]
	}
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty(){
		return -1
	}else{
		return this.nums[(this.rear-1+this.size) % this.size]

	}
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {


	return this.rear==this.front
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return (this.rear+1) % this.size ==(this.front) % this.size
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */