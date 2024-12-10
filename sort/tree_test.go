package sort

import (
	"testing"
)

func TestTreeSort(t *testing.T) {

	arr := []int{0, 10, 2, 8, 6, 4, 5, 7, 3, 9, 1}
	sortedArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("sort", func(t *testing.T) {
		got := make([]int, 0)
		TreeSort(&arr, &got)
		for i := range got {
			if got[i] != sortedArr[i] {
				t.Fail()
			}
		}
	})

	t.Run("empty", func(t *testing.T) {
		var expected []int = nil
		got := make([]int, 0)
		TreeSort(&expected, &got)
		if len(got) != 0 {
			t.Fail()
		}
	})
}
