// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package singlist implements a simple list package. It defines a structer,
// with methods for managing the nodes on the list. It could be taken as a
// container in some cases to support the data storage, which are easier to
// use than involving the default golang data type like slice/map/struct.
// The single list is growing forward on a direction with the data adding in.
// Thereby different operations would be based on the list head, currently
// the extension, modification, deletion, location and the sort of the entire
// list are supportive to use. Fatal errors will occur if the restricted
// opeartion on the function is to call.Error message will also report
// if it detect the disallowed operation is being issued.


package singlelist

import (
	"errors"
	"fmt"
	"reflect"
)

// Typically, a node will contain a data with any supported type and
// a pointer which will be used to refer to the next node on the list.
type Node struct {
	Data interface{}
	Next *Node
}

// To simply, the list type will be the same as a node if the list is
// empty or within single node. This would of course a single list.
type List = Node

func NewNode() *Node {
	return &Node{
		Data: nil,
		Next: nil,
	}
}

// The mode to sort the list, whereby ASCEND means to sort the list
// with the ascend order but if using DESCEND mode, then the order
// to sort the list will be descend node by node.
const (
	ASCEND  int = iota << 1
	DESCEND int = iota << 2
)

// This single list will have the head node included. So anyhow the
// first node would be created to add into the list once it is init.
func InitList() (head *Node) {
	head = NewNode()
	return
}

// To extend a list with new node adding in. 
// Return nil if errors appear.
func (l *List) AddNode(data interface{}) (err error) {
	if l == nil {
		return errors.New("Invalid list head found")
	}

	for l.Next != nil {
		l = l.Next
	}

	l.Data = data
	l.Next = NewNode()
	return nil
}

func (l *List) IsEmpty() bool {
	if l.Data == nil && l.Next == nil {
		return true
	} else {
		return false
	}
}

// The length of the list wont walk through the last nil node.
// It only marks the numbers of the nodes within data included.
func (l *List) Length() (length int) {
	for length = 0; l.Next != nil; l = l.Next {
		length++
	}
	return
}

// Insert a node ahead of the given postion, Note the pos parameter
// will be the index of that node plugged in the list.
// Data will be assigned to the node respectively once it is located.
func (l *List) InsertBefore(pos int, data interface{}) error {
	if l.Length() <= pos+1 {
		return errors.New("Position is byond list length")
	}

	var index int = 0
	var p *Node = l
	for p.Next != nil && index < pos-1 {
		p = p.Next
		index++
	}

	if p == nil || index > pos-1 {
		return errors.New("Position is out of list bound")
	}

	tempNode := NewNode()
	tempNode.Data = data
	tempNode.Next = p.Next
	p.Next = tempNode

	return nil
}

// Insert a node after the node given the position provied.
// Likewise, pos parameter will be taken as the index of the
// node plugged inside the list and once the related node is
// located, then a new node will be created to append behind.
func (l *List) InsertAfter(pos int, data interface{}) error {
	if l.Length() <= pos {
		return errors.New("Position is byond list length")
	}

	index := 0
	p := l
	for p.Next != nil && index < pos {
		p = p.Next
		index++
	}

	if p.Next == nil || index > pos {
		return errors.New("Position is out of list bound")
	}

	tempNode := NewNode()
	tempNode.Data = data
	tempNode.Next = p.Next
	p.Next = tempNode

	return nil
}

// Delete the node with the givin position.
func (l *List) Delete(pos int) error {
	if l.Length() < pos {
		return errors.New("Deleted failed as position is byond the length")
	}

	p := l
	index := 0
	for p.Next != nil && index < pos-1 {
		p = p.Next
		index++
	}

	p.Data = nil
	p.Next = nil
	return nil

}

// Get the data of the node with the position provided.
func (l *List) Find(pos int) interface{} {
	if l.Length() < pos {
		return errors.New("Located failed as position is byond the length")
	}

	for index := 0; l.Next != nil && index < pos-1; {
		l = l.Next
		index++
	}

	return l.Data
}

// Locate the matched node if the data of the node is as expected.
// Return the index of the node inside the list.
func (l *List) FindMatchedValue(d interface{}) (pos int) {
	//Init an invalid position to return if errors occur
	pos = -1

	if l == nil || l.Data == nil {
		return
	}

	for index := 0; l.Next != nil; l = l.Next {
		if reflect.DeepEqual(d, l.Data) {
			pos = index
		} else {
			index++
		}
	}
	return
}

// Reverse the list and the next pointer in the head will refer to the nil
// once the entire list is done with the reversal. This will tolly change
// the sequence of the list.
func (l *List) Reverse() *Node {
	if l.Next == nil && l.Data == nil {
		return l
	}

	// We need to keep the first head pointer referencing onto
	// the one with member of {nil, nil} (possibly is the last
	// node of the original list) once the list has been reversed.
	// So that a temp pointer here is needed to complete it.
	head := l

	// p and q are used to walk through the orginal list nodes
	// thereby with the movement going ahead, p and q will refer
	// to the next node and the one next to the next respectively.
	// This will be terminated until the last node has {nil, nil}.
	p := l
	var q *Node

	for p != nil && p.Next != nil {
		q = p.Next
		p.Next = l
		l = p
		p = q
	}

	// Now the last node after reversing is with {nil, nil} members
	// so that we just make the orginal head pointer refer to it
	// in case this last node becomes as an orphan one.
	head.Next = p

	return l
}

// Do a comparsion bewteen v1 and v2 with any data type.
// The assumption is once v1 is larger than v2, then true
// would be returned otherwise false is returned.
func cmp(v1, v2 interface{}) bool {
	switch v1.(type) {
	case int:
		return v1.(int) >= v2.(int)
	case int64:
		return v1.(int64) >= v2.(int64)
	case float32:
		return v1.(float32) >= v2.(float32)
	case float64:
		return v1.(float64) >= v2.(float64)
	case string:
		if strings.Compare(v1.(string), v2.(string)) != 0 {
			return false
		}
	default:
		panic("Unsupport data type in the list")
	}
	return false
}

// Swap the data between the two nodes.
func swap(p *Node, q *Node) {
	tmp := p.Data
	p.Data = q.Data
	q.Data = tmp
}

// Qucik sort implementations. The tail should be always nil if
// you want to sort the entire list otherwise the special node
// needs to be provided. Parameter 'mode' is using to determine
// the sequence of the list is ascend or descend.
func (l *List) QuickSort(tail *Node, mode int) {
	// Should find a bound node to terminate the recrusive call
	// And once the walk-through of the head reaches the tail node
	// The sort process is coming to the end. Note neither the
	// single node nor the nil node would be used to sort.
	if l == nil ||
		l.Next == nil ||
		l.Data == nil ||
		l == tail {
		return
	}

	// Define two temp nodes in the each sort walk-through process
	// Specifically, p is walking behind the q which is using to
	// locate the expected node(whose value is larger or smaller than
	// the base one, here the base node 'key' will always be the first
	// node). Once the value in q is located, then we should switch
	// out to have p and q swapped with their values, this would be
	// continued until q is coming to the end of the list. In which
	// case, one partition then is completed.
	p := l
	q := l.Next
	key := l.Data

	for q != nil && q.Data != nil {
		switch key.(type) {
		case int, int64, float32, float64:
			if cmp(key, q.Data) &&
				mode == ASCEND {
				p = p.Next
				swap(p, q)
			}
		case string:
			fmt.Println("TBD..")
		}
		q = q.Next
	}

	// Now the walk-through is done as the p is to the end of the
	// node which contains the data of last expecting one(larger or
	// smaller) than the base 'key' value.
	swap(l, p)

	// Recrusively to call the function partition by partition.
	l.QuickSort(p, mode)
	p.Next.QuickSort(nil, mode)

}

// Select sort implementations. Parameter 'mode' is using to determine
// the sequence of the list is ascend or descend.
func (l *List) SelectSort(mode int) {
	if l == nil || l.Next == nil {
		return
	}

	for p := l; p.Next != nil; p = p.Next {
		base := p
		for q := p.Next; q.Next != nil; q = q.Next {

			// Here we are going to select the largest or smallest node
			// Once found, the data in base node would be exchanged with
			// one we are locating. Afterwards, the base could be used
			// to assign back onto the orignal list with the p goes forward. 
			if cmp(base.Data, q.Data) &&
				mode == ASCEND {
				fmt.Println("In the SelectSort", base.Data, q.Data)
				swap(q, base)
			}
		}

		// In case the node selected is not synced into the one in the original
		// list, here will confirm they are sharing the same value and would
		// swap if the data is not equal in between.
		if !reflect.DeepEqual(base.Data, p.Data) {
			swap(base, p)
		}
	}
}

// Walk through the list and then print it out.
func (l *List) DumpList() {
	if l == nil {
		panic("Nil node found")
	}
	for ; l.Next != nil; l = l.Next {
		fmt.Println(l.Data)
	}
}
