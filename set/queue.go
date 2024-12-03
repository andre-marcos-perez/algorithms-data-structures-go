package set

import "errors"

var (
	ErrQueueUnderflow = errors.New("queue underflow ")
	ErrQueueOverflow  = errors.New("queue overflow")
)

type Queue struct {
	arr    [10]int
	head   int
	tail   int
	isFull bool
}

func NewQueue() *Queue {
	return &Queue{
		arr:    [10]int{},
		head:   0,
		tail:   0,
		isFull: false,
	}
}

func (q *Queue) Enqueue(k int) error {
	if q.isFull {
		return ErrQueueOverflow
	}
	nextTail := (q.tail + 1) % len(q.arr)
	q.isFull = nextTail == q.head
	q.arr[q.tail] = k
	q.tail = nextTail
	return nil
}

func (q *Queue) Peek() (int, error) {
	if q.head == q.tail {
		return 0, ErrQueueUnderflow
	}
	k := q.arr[q.head]
	return k, nil
}

func (q *Queue) Dequeue() (int, error) {
	k, err := q.Peek()
	if err != nil {
		return 0, err
	}
	q.head = (q.head + 1) % len(q.arr)
	return k, nil
}
