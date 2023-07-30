package linkedlist

import "testing"

func Test_ss(t *testing.T) {
	linkedList := &LinkedList{}
	linkedList.TailAdd(1)
	linkedList.TailAdd(2)
	linkedList.TailAdd(3)
	linkedList.Print()
}
