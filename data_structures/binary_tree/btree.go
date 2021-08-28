package binarytree

import "fmt"

// BTree Returns a binary tree structure which contains only a root Node
type BTree struct {
	Root *Node
}

// calculateDepth helper function for BTree's depth()
func calculateDepth(n *Node, depth int) int {
	if n == nil {
		return depth
	}
	return Max(calculateDepth(n.left, depth+1), calculateDepth(n.right, depth+1))
}

// Insert a value in the BTree
func Insert(root *Node, val int) *Node {
	if root == nil {
		return NewNode(val)
	}
	if val < root.val {
		root.left = Insert(root.left, val)
	} else {
		root.right = Insert(root.right, val)
	}
	return root
}

// Depth returns the calculated depth of a binary tree
func (t *BTree) Depth() int {
	return calculateDepth(t.Root, 0)
}

// InOrder add's children to a node in order left first then right recursively
func InOrder(n *Node) {
	if n != nil {
		InOrder(n.left)
		fmt.Print(n.val, " ")
		InOrder(n.right)
	}
}

// InOrderSuccessor Goes to the left
func InOrderSuccessor(root *Node) *Node {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

// BstDelete removes the node
func BstDelete(root *Node, val int) *Node {
	if root == nil {
		return nil
	}
	if val < root.val {
		root.left = BstDelete(root.left, val)
	} else if val > root.val {
		root.right = BstDelete(root.right, val)
	} else {
		// this is the node to delete
		// node with one child
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			n := root.right
			d := InOrderSuccessor(n)
			d.left = root.left
			return root.right
		}
	}
	return root
}

// PreOrder Preorder
func PreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.val, " ")
	PreOrder(n.left)
	PreOrder(n.right)
}

// PostOrder PostOrder
func PostOrder(n *Node) {
	if n == nil {
		return
	}
	PostOrder(n.left)
	PostOrder(n.right)
	fmt.Print(n.val, " ")
}

// LevelOrder LevelOrder
func LevelOrder(root *Node) {
	var q []*Node // queue
	var n *Node   // temporary node

	q = append(q, root)

	for len(q) != 0 {
		n, q = q[0], q[1:]
		fmt.Print(n.val, " ")
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

// Max Function that returns max of two numbers - possibly already declared.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
