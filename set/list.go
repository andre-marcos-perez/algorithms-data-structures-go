package set

import "errors"

var (
	ErrListEmpty    = errors.New("empty")
	ErrListNotFound = errors.New("not found")
)

// List implements an unsorted and double linked list
type List struct {
	head *ListNode
	size int
}

func NewList() *List {
	return &List{
		head: nil,
		size: 0,
	}
}

// Insert to head
func (l *List) Insert(k int) {
	node := NewListNode(k)
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.size += 1
}

func (l *List) Peek() (*ListNode, error) {
	if l.head == nil {
		return nil, ErrListEmpty
	}
	return l.head, nil
}

func (l *List) Find(k int) (*ListNode, error) {
	node, err := l.Peek()
	if err != nil {
		return nil, err
	}
	for {
		if node == nil {
			return nil, ErrListNotFound
		}
		if node.k == k {
			return node, nil
		}
		node = node.next
	}
}

func (l *List) Delete(k int) (*ListNode, error) {
	node, err := l.Find(k)
	if err != nil {
		return nil, err
	}
	// delete head
	if node.prev == nil {
		l.head = node.next
	} else {
		node.prev.next = node.next
	}
	// delete tail
	if node.next != nil {
		node.next.prev = node.prev
	}
	return node, nil
}
