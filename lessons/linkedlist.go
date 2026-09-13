package learngo

import (
	"container/list"
	"fmt"
)

type LNode[T any] struct {
	Value any
	Next  *LNode[T]
}

type LinkedList[T any] struct {
	Head *LNode[T]
}

func (l *LinkedList[T]) AddToTail(value T) {
	newNode := &LNode[T]{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// DoubleLinkedList
type Node[T any] struct {
	Value      T
	Prev, Next *Node[T]
}

type DoubleLinkedList[T any] struct {
	Head, Tail *Node[T]
	Length     int
}

func (t *DoubleLinkedList[T]) PushBack(value T) *Node[T] {
	newNode := &Node[T]{
		Value: value,
		Prev:  nil,
		Next:  nil,
	}

	if t.Tail == nil {
		// Empty list: the node is both the first and last node.
		t.Head = newNode
		t.Tail = newNode
	} else {
		// tail is the last car in the train.
		newNode.Prev = t.Tail // The new car points back to the old last car.
		t.Tail.Next = newNode // The old last car points forward to the new car.
		t.Tail = newNode      // The new car is now the last car.
	}

	t.Length++

	return newNode
}

func (t *DoubleLinkedList[T]) PushFront(value T) *Node[T] {
	newNode := &Node[T]{
		Value: value,
		Prev:  nil,
		Next:  nil,
	}

	if t.Head == nil {
		t.Head = newNode
		t.Tail = newNode
	} else {
		newNode.Next = t.Head
		t.Head.Prev = newNode
		t.Head = newNode
	}

	t.Length++
	return newNode
}

func (t *DoubleLinkedList[T]) Remove(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		t.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		t.Tail = node.Prev
	}

	node.Prev = nil
	node.Next = nil
	t.Length--

	return node
}

func ExampleDoubleLinkedList() {
	var list DoubleLinkedList[string]

	nodeB := list.PushBack("B")
	list.PushBack("C")
	list.PushFront("A")
	removed := list.Remove(nodeB)

	for node := list.Head; node != nil; node = node.Next {
		if node != list.Head {
			fmt.Print(" <-> ")
		}
		fmt.Print(node.Value)
	}

	fmt.Printf("\nRemoved: %s\n", removed.Value)
	fmt.Printf("Head: %s, Tail: %s, Length: %d\n", list.Head.Value, list.Tail.Value, list.Length)

	// Output:
	// A <-> C
	// Removed: B
	// Head: A, Tail: C, Length: 2
}

// standart library
func ExampleLinkedList() {

	myList := list.New()
	myList.PushBack("A")
	myList.PushBack("B")
	myList.PushBack("C")

	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

}
