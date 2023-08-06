package linkedlist

import (
	"fmt"
	"testing"
)

func Test_Linked(t *testing.T) {
	linkedList := &LinkedList{}
	linkedList.TailAdd(1)
	linkedList.TailAdd(1)
	linkedList.TailAdd(1)

	linkedList.TailAdd(2)
	linkedList.TailAdd(3)
	linkedList.TailAdd(1)

	linkedList.TailAdd(4)
	linkedList.TailAdd(5)

	linkedList.Print()
	fmt.Println("=====>", linkedList.Length())

	//linkedList.delete(1)
	linkedList.Print()

}

func Test_removeElements(t *testing.T) {
	linkedList := &LinkedList{}
	linkedList.TailAdd(1)
	linkedList.TailAdd(1)
	linkedList.TailAdd(1)

	linkedList.TailAdd(2)
	linkedList.TailAdd(3)
	linkedList.TailAdd(1)

	linkedList.TailAdd(4)
	linkedList.TailAdd(5)

	linkedList.Print()
	fmt.Println("=====>", linkedList.Length())
	linkedList.removeElements(1)
	//linkedList.Head = removeElements(linkedList.Head, 1)
	linkedList.Print()
	fmt.Println("=====>", linkedList.Length())

}
