package linkedlist

type ListNodeDouble struct {
	Val  int
	Next *ListNodeDouble
	Prev *ListNodeDouble
}

func NewListNodeDouble(val int) *ListNodeDouble {
	return &ListNodeDouble{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}
