package queue

// CircularQueue 循环队列
type CircularQueue struct {
	data     []int
	head     int
	tail     int
	size     int
	capacity int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		data:     make([]int, capacity),
		head:     0,
		tail:     0,
		size:     0,
		capacity: capacity,
	}
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.size == 0
}

func (cq *CircularQueue) IsFull() bool {
	return cq.size == cq.capacity
}

func (cq *CircularQueue) Enqueue(value int) bool {
	if cq.IsFull() {
		return false
	}
	cq.data[cq.tail] = value
	cq.tail = (cq.tail + 1) % cq.capacity
	cq.size++
	return true
}

func (cq *CircularQueue) Dequeue() (int, bool) {
	if cq.IsEmpty() {
		return -1, false
	}
	value := cq.data[cq.head]
	cq.head = (cq.head + 1) % cq.capacity
	cq.size--
	return value, true
}

func (cq *CircularQueue) Front() (int, bool) {
	if cq.IsEmpty() {
		return -1, false
	}
	return cq.data[cq.head], true
}
