package main

import (
	"container/list"
	"fmt"
)

func main() {
	// new linked list
	queue := list.New()

	// Simply append to enqueue.
	queue.PushBack("hola")
	queue.PushBack(2)
	queue.PushBack(3)

	// Dequeue
	front := queue.Front()
	fmt.Println(queue.Front().Value)
	queue.Remove(front)
	fmt.Println(queue.Front().Value)
	queue.Remove(front)
}
