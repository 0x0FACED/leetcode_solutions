package queue

import "fmt"

// This is our item of queue
type Node struct {
	Val  interface{}
	next *Node
}

// Private func to create a node
func newNode(x interface{}) *Node {
	return &Node{
		Val: x,
	}
}

// This is our queue struct
type Queue struct {
	head *Node
	tail *Node
}

// Just return head of queue to check if it empty
func (q *Queue) Head() *Node {
	return q.head
}

// Create empty queue and return link to it
func New() *Queue {
	return &Queue{}
}

// Method to push value to the end of queue
func (q *Queue) PushLast(value interface{}) {
	var node = newNode(value)

	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

// Delete first element fom queue and return it
func (q *Queue) PopFirst() interface{} {
	if q.head == nil {
		return nil
	}

	var value = q.head.Val
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return value
}

func (q *Queue) PrintQueue() {
	if q.head == nil {
		return
	}

	var ptr = q.head

	for ptr != nil {
		fmt.Printf("%d ", ptr.Val)
		ptr = ptr.next
		fmt.Println()
	}

}
