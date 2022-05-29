package main

import (
	"errors"
	"fmt"
	"sort"
)

// https://exercism.org/tracks/go/exercises/tree-building/edit

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func (node *Node) AddNode(r Record) error {
	id := r.ID
	parent := r.Parent
	newNode := Node{ID: id, Children: make([]*Node, 0)}

	var current []*Node
	current = append(current, node)
	var next []*Node

	for len(current) > 0 {
		for _, n := range current {
			if n.ID == parent {
				n.Children = append(n.Children, &newNode)
				return nil
			} else {
				next = append(next, n.Children...)
			}
		}
		current = next
		next = make([]*Node, 0)
	}
	return errors.New("parent node not found")
}

func Build(records []Record) (*Node, error) {
	/*
		Psuedo code:
		- sort records by ID
		- add function to find node with given ID
		- step by step fill in the tree
	*/
	if len(records) == 0 {
		return nil, nil
	}

	// sort records by ID
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	// verify no duplicates and continous ID's
	for i := 0; i < len(records)-1; i++ {
		if records[i].ID == records[i+1].ID {
			return nil, errors.New("duplicate node")
		}
		if records[i].ID != i || records[i+1].ID != i+1 {
			return nil, errors.New("non-continous ID's")
		}
	}

	// step by step fill in the tree
	var rootNode Node
	for i, v := range records {
		if i == 0 {
			if v.ID != v.Parent {
				return nil, errors.New("no root node")
			}
			rootNode.ID = v.ID
		} else {
			if err := rootNode.AddNode(v); err != nil {
				return nil, err
			}
		}
	}
	return &rootNode, nil
}

func main() {
	fmt.Println()

	test1 := []Record{
		{ID: 0},
		{ID: 1, Parent: 0},
		{ID: 2, Parent: 0},
	}
	ret1, _ := Build(test1)
	fmt.Println(ret1)

	test2 := []Record{
		{ID: 2, Parent: 1},
		{ID: 1, Parent: 0},
		{ID: 3, Parent: 2},
		{ID: 0},
	}
	ret2, _ := Build(test2)
	fmt.Println(ret2)
}
