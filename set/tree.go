package set

import (
	"errors"
	"fmt"
)

var (
	ErrTreeEmpty    = errors.New("empty")
	ErrTreeNotFound = errors.New("not found")
)

// Tree implements a binary search tree
type Tree struct {
	root *TreeNode
}

func NewTree() *Tree {
	return &Tree{
		root: nil,
	}
}

func (t *Tree) Iter(node *TreeNode) {
	if node.left != nil {
		t.Iter(node.left)
	}
	fmt.Println(node.k)
	if node.right != nil {
		t.Iter(node.right)
	}
}

func (t *Tree) Insert(k int) {
	var parentNode *TreeNode = nil
	node := t.root
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
		t.root = node
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
