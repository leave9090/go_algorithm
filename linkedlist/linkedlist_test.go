package linkedlist

import (
	"fmt"
	"testing"
)

func Test_Linked(t *testing.T) {
	linkedList := &LinkedList{}
	linkedList.TailAdd(1)
	linkedList.TailAdd(2)
	linkedList.TailAdd(3)
	linkedList.Print()
	fmt.Println("=====>", linkedList.Length())

	linkedList.Del(2)
	linkedList.Print()

}
