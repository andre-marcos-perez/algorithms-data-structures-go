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

type TreeNode struct {
	k      int
	parent *TreeNode
	left   *TreeNode
	right  *TreeNode
}

func NewTreeNode(k int) *TreeNode {
	return &TreeNode{
		k:     k,
		left:  nil,
		right: nil,
	}
}
