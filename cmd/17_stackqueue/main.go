package main

import (
	"fmt"
)

//Stack represents a stack that holds a slice
type Stack struct {
	items []int
}
//Push
func (s *Stack) Push(i int) { 
	s.items = append(s.items, i) 
}
//Pop
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l] 
	s.items = s.items[:l]
	return toRemove
}

//Queue
type Queue struct {
	items []int
}
//Enqueue adds a value at the end
func (q *Queue) Enqueue(i int) { 
	q.items = append(q.items, i) 
}
//Dequeue removes a value at the front
func (q *Queue) Dequeue() int{
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)


	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(150)
	myQueue.Enqueue(250)
	myQueue.Enqueue(350)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}