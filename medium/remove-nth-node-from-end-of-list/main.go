package main

type ListNode struct {
     Val int
     Next *ListNode
}

func main(){

}

type ring struct {
	data []*ListNode
	current int
}

func (r *ring) setAndMoveToNext(node *ListNode){
	r.data[r.current] = node
	r.current = r.calculateNextPosition()
}

func (r *ring) moveToNextAndSet(node *ListNode){
	r.current = r.calculateNextPosition()
	r.data[r.current] = node
}

func (r *ring) getCurrent() *ListNode{
	return r.data[r.current]
}

func (r *ring) getAtPreviousPosition(steps int) *ListNode {
	previousPosition := r.calculatePreviousPosition(steps)
	if len(r.data) == 0 || previousPosition < 0 {
		return nil
	}
	return r.data[previousPosition]
}

func (r *ring) calculateNextPosition() int {
	r.current++
	if r.current == len(r.data) {
		r.current = 0
	}

	return r.current
}

func (r *ring) calculatePreviousPosition(steps int) int {
	previous := r.current - steps
	if previous < 0 {
		previous = (len(r.data )) + previous
	}
	return previous
}


func newRing(size int) ring {
	return ring{ make([]*ListNode, size),0 }
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	var node *ListNode = head
	var r ring = newRing(n + 1)
	r.setAndMoveToNext(head)

	for node.Next != nil {
		r.setAndMoveToNext(node.Next)
		node = node.Next
	}

	if r.getAtPreviousPosition(n) == head  {
		return head.Next
	}

	prior := r.getAtPreviousPosition(n + 1)
	post := r.getAtPreviousPosition(n - 1)
	if post == prior {
		post = nil
	}

	if prior != nil {
		prior.Next = post
	}

	return head
}

