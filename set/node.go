package set

type ListNode struct {
	k    int
	next *ListNode
	prev *ListNode
}

func NewListNode(k int) *ListNode {
	return &ListNode{
		k:    k,
		next: nil,
		prev: nil,
	}
}
