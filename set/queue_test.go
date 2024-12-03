package set

import (
	"errors"
	"testing"
)

func TestQueue(t *testing.T) {

	t.Run("create queue", func(t *testing.T) {
		queue := NewQueue()
		if queue == nil {
			t.Fail()
		}
	})

	t.Run("enqueue", func(t *testing.T) {
		expected := 10
		queue := NewQueue()
		err := queue.Enqueue(expected)
		if err != nil {
			t.Fail()
		}
		got, err := queue.Peek()
		if err != nil {
			t.Fail()
		}
		if got != expected {
			t.Fail()
		}
	})

	t.Run("dequeue", func(t *testing.T) {
		queue := NewQueue()
		expectedDequeued := 10
		err := queue.Enqueue(expectedDequeued)
		if err != nil {
			t.Fail()
		}
		expectedPeeked := 20
		err = queue.Enqueue(expectedPeeked)
		if err != nil {
			t.Fail()
		}
		got, err := queue.Dequeue()
		if err != nil {
			t.Fail()
		}
		if got != expectedDequeued {
			t.Fail()
		}
		got, err = queue.Peek()
		if err != nil {
			t.Fail()
		}
		if got != expectedPeeked {
			t.Fail()
		}
	})

	t.Run("queue underflow", func(t *testing.T) {
		queue := NewQueue()
		_, err := queue.Dequeue()
		if err == nil {
			t.Fail()
		} else {
			if !errors.Is(err, ErrQueueUnderflow) {
				t.Fail()
			}
		}
	})

	t.Run("queue overflow", func(t *testing.T) {
		queue := NewQueue()
		var err error
		for i := range len(queue.arr) + 1 {
			err = queue.Enqueue(i)
			if err != nil {
				break
			}
		}
		if err == nil {
			t.Fail()
		} else {
			if !errors.Is(err, ErrQueueOverflow) {
				t.Fail()
			}
		}
	})

	t.Run("rotate array 3 steps", func(t *testing.T) {
		var v int
		var err error
		steps := 3
		queue := NewQueue()
		arrA := []int{1, 2, 3, 4, 5}
		arrB := []int{0, 0, 0, 0, 0}
		for _, v = range arrA {
			err = queue.Enqueue(v)
			if err != nil {
				t.Fail()
			}
		}
		for i := 0; i < steps; i++ {
			v, err = queue.Dequeue()
			if err != nil {
				t.Fail()
			}
			err = queue.Enqueue(v)
			if err != nil {
				t.Fail()
			}
		}
		for i := range arrA {
			v, err = queue.Dequeue()
			if err != nil {
				t.Fail()
			}
			arrB[i] = v
		}
		for i := range arrA {
			if arrA[i] != arrB[(len(arrB)-steps+i)%len(arrB)] {
				t.Fail()
			}
		}
	})
}
