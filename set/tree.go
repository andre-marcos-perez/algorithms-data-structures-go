package set

import (
	"errors"
)

var (
	ErrTreeEmpty    = errors.New("empty")
	ErrTreeNotFound = errors.New("not found")
)

// Tree implements a binary search tree
type Tree struct {
	Root *TreeNode
}

func NewTree() *Tree {
	return &Tree{
		Root: nil,
	}
}

func (t *Tree) Iter(node *TreeNode, sequence *[]int) {
	if node.left != nil {
		t.Iter(node.left, sequence)
	}
	*sequence = append(*sequence, node.k)
	if node.right != nil {
		t.Iter(node.right, sequence)
	}
}

func (t *Tree) Insert(k int) {
	var parentNode *TreeNode = nil
	node := t.Root
	for {
		if node == nil {
			break
		}
		parentNode = node
		if k < node.k {
			node = node.left
		} else {
			node = node.right
		}
	}
	node = NewTreeNode(k)
	node.parent = parentNode
	if node.parent == nil {
		t.Root = node
		return
	}
	if k < node.parent.k {
		parentNode.left = node
	} else {
		parentNode.right = node
	}
}

func (t *Tree) Find(node *TreeNode, k int) (*TreeNode, error) {
	if node == nil {
		return nil, ErrTreeNotFound
	}
	if node.k == k {
		return node, nil
	}
	if k < node.k {
		return t.Find(node.left, k)
	} else {
		return t.Find(node.right, k)
	}
}

func (t *Tree) Min(node *TreeNode) (*TreeNode, error) {
	if node == nil {
		return nil, ErrTreeEmpty
	}
	if node.left == nil {
		return node, nil
	} else {
		return t.Min(node.left)
	}
}

func (t *Tree) Max(node *TreeNode) (*TreeNode, error) {
	if node == nil {
		return nil, ErrTreeEmpty
	}
	if node.right == nil {
		return node, nil
	} else {
		return t.Max(node.right)
	}
}

func (t *Tree) Prev(node *TreeNode, k int) (*TreeNode, error) {
	found, err := t.Find(node, k)
	if err != nil {
		return nil, err
	}
	// go down
	if found.left != nil {
		return t.Max(found.left)
	}
	// go up
	parent := found.parent
	for {
		// root reached
		if parent == nil {
			break
		}
		// inflection reached
		if parent.left != found {
			break
		}
		found = parent
		parent = found.parent
	}
	return parent, nil
}

func (t *Tree) Next(node *TreeNode, k int) (*TreeNode, error) {
	found, err := t.Find(node, k)
	if err != nil {
		return nil, err
	}
	// go down
	if found.right != nil {
		return t.Min(found.right)
	}
	// go up
	parent := found.parent
	for {
		// root reached
		if parent == nil {
			break
		}
		// inflection reached
		if parent.right != found {
			break
		}
		found = parent
		parent = found.parent
	}
	return parent, nil
}

// TODO: DELETE
