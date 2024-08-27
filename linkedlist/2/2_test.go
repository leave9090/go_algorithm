package main

import (
	"testing"
)

func TestNewListNode(t *testing.T) {

	n0 := NewListNode(0)
	n1 := NewListNode(1)
	n2 := NewListNode(2)
	n3 := NewListNode(3)
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3

	//n0.PrintNode()

	//头插入
	//ll := n0.InsertNode(9)
	//ll.PrintNode()

	////尾插入
	n0.AppendNode(4)
	n0.PrintNode()
}

func TestListNode_InsertNode(t *testing.T) {

}
