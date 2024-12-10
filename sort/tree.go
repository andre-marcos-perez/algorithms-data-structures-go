package sort

import (
	"github.com/andre-marcos-perez/algorithms-data-structures-go/set"
)

func TreeSort(sequence *[]int, sorted *[]int) {
	if *sequence == nil {
		return
	}
	tree := set.NewTree()
	for _, v := range *sequence {
		tree.Insert(v)
	}
	tree.Iter(tree.Root, sorted)
}
