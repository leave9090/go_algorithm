package queue

import (
	"fmt"
	"testing"
)

func Test_Func(t *testing.T) {
	cq := NewCircularQueue(5)

	for i := 1; i <= 5; i++ {
		fmt.Printf("Enqueued %d\n", i)
		cq.Enqueue(i)
	}

	for !cq.IsEmpty() {
		value, _ := cq.Dequeue()
		fmt.Printf("Dequeued %d\n", value)
	}

	// 再次入队
	for i := 6; i <= 10; i++ {
		fmt.Printf("Enqueued %d\n", i)
		cq.Enqueue(i)
	}

	for !cq.IsEmpty() {
		value, _ := cq.Dequeue()
		fmt.Printf("Dequeued %d\n", value)
	}
}
