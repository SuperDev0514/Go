package linkedlist

import "fmt"

// Doubly structure with just the Head Node
// We call it `Doubly` to make it easier to
// understand when calling this in peoples
// own local code to understand and experiment
// with this easily.
// For example, we can use gotools `go get` command to get
// this repository cloned inside the
// $GOPATH/src/github.com/TheAlgorithms/Go (you can do this
// manually as well) and use the import statement as follows:
//
// `import "github.com/TheAlgorithms/Go/structure/linkedlist"`
//
// and call linkedlist.Doubly to create a new doubly linked list.
type Doubly[T any] struct {
	Head *Node[T]
}

func NewDoubly[T any]() *Doubly[T] {
	return &Doubly[T]{}
}

// AddAtBeg Add a node to the beginning of the linkedlist
func (ll *Doubly[T]) AddAtBeg(val T) {
	n := NewNode(val)
	n.Next = ll.Head

	if ll.Head != nil {
		ll.Head.Prev = n
	}

	ll.Head = n

}

// AddAtEnd Add a node at the end of the linkedlist
func (ll *Doubly[T]) AddAtEnd(val T) {
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = n
	n.Prev = cur
}

// DelAtBeg Delete the node at the beginning of the linkedlist
func (ll *Doubly[T]) DelAtBeg() (T, bool) {
	if ll.Head == nil {
		var r T
		return r, false
	}

	cur := ll.Head
	ll.Head = cur.Next

	if ll.Head != nil {
		ll.Head.Prev = nil
	}
	return cur.Val, true
}

// DetAtEnd Delete a node at the end of the linkedlist
func (ll *Doubly[T]) DelAtEnd() (T, bool) {
	// no item
	if ll.Head == nil {
		var r T
		return r, false
	}

	// only one item
	if ll.Head.Next == nil {
		return ll.DelAtBeg()
	}

	// more than one, go to second last
	cur := ll.Head
	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	retval := cur.Next.Val
	cur.Next = nil
	return retval, true
}

// Count Number of nodes in the linkedlist
func (ll *Doubly[T]) Count() int {
	var ctr int = 0

	for cur := ll.Head; cur != nil; cur = cur.Next {
		ctr += 1
	}

	return ctr
}

// Reverse Reverse the order of the linkedlist
func (ll *Doubly[T]) Reverse() {
	var Prev, Next *Node[T]
	cur := ll.Head

	for cur != nil {
		Next = cur.Next
		cur.Next = Prev
		cur.Prev = Next
		Prev = cur
		cur = Next
	}

	ll.Head = Prev
}

// Display display the linked list
func (ll *Doubly[T]) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}

// DisplayReverse Display the linkedlist in reverse order
func (ll *Doubly[T]) DisplayReverse() {
	if ll.Head == nil {
		return
	}
	var cur *Node[T]
	for cur = ll.Head; cur.Next != nil; cur = cur.Next {
	}

	for ; cur != nil; cur = cur.Prev {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}
