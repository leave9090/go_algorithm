package linkedlist

import "testing"

func Test_ss(t *testing.T) {
	list := newList()
	//list.add(1)
	//list.add(2)
	//list.add(3)
	//list.add(4)
	//list.add(5)
	//list.add(6)
	list.Insert(1, 1)
	list.Insert(1, 2)
	list.Insert(1, 3)
	list.rangeLink()
}
