package main

import (
	"errors"
	"fmt"
)

// https://exercism.org/tracks/go/exercises/linked-list

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

type Node struct {
	prev  *Node
	next  *Node
	Value interface{}
}

type List struct {
	head *Node
	tail *Node
}

func NewList(args ...interface{}) *List {
	l := List{}
	if len(args) == 0 {
		return &l
	} else if len(args) == 1 {
		n := Node{nil, nil, args[0]}
		l.head = &n
		l.tail = &n
		return &l
	}

	last := &Node{nil, nil, 0}
	for i := 0; i < len(args); i++ {
		if i == 0 {
			n := Node{nil, nil, args[i]}
			l.head = &n
			last = l.head
		} else if i == len(args)-1 {
			n := Node{last, nil, args[i]}
			last.next = &n
			l.tail = &n
		} else {
			n := Node{last, nil, args[i]}
			last.next = &n
			last = &n
		}
	}
	return &l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	if l.head == nil {
		n := Node{nil, nil, v}
		l.head = &n
		l.tail = &n
	} else {
		n := Node{nil, l.head, v}
		l.head.prev = &n
		l.head = &n
	}
}

func (l *List) Push(v interface{}) {
	if l.head == nil {
		n := Node{nil, nil, v}
		l.head = &n
		l.tail = &n
	} else {
		n := Node{l.tail, nil, v}
		l.tail.next = &n
		l.tail = &n
	}
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return 0, errors.New("empty linked list")
	}

	v := *l.head
	if l.head.next != nil {
		l.head.next.prev = nil
		l.head = l.head.next
	} else {
		l.head = nil
		l.tail = nil
	}
	return v.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
		return 0, errors.New("empty linked list")
	}

	v := *l.tail
	if l.tail.prev != nil {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
	} else {
		l.head = nil
		l.tail = nil
	}
	return v.Value, nil
}

func (l *List) Reverse() {
	if l.head == l.tail {
		return
	}

	ptr := l.head
	for ptr != l.tail {
		temp := ptr.next
		ptr.next, ptr.prev = ptr.prev, ptr.next
		ptr = temp
	}
	ptr.next, ptr.prev = ptr.prev, ptr.next

	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func main() {
	l := NewList(1, 2, 3, 4)
	v1, _ := l.Shift()
	v2, _ := l.Shift()
	v3, _ := l.Shift()
	v4, _ := l.Shift()
	fmt.Println(v1, v2, v3)
}
