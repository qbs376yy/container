// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package singlelist_test

import (
	"fmt"
	"singlelist"
	"testing"
)

var (
	Head *singlelist.Node
	revList *singlelist.Node
)

func CreateList() (*singlelist.List) {
	Head = singlelist.InitList()
	Cur := Head
	for i := 0; i < 5; i++ {
		Cur.AddNode(i)
	}

	return Head

}

func TestCreateList(t *testing.T) {
	Head = CreateList()

	length := Head.Length()
	if length != 5 {
		t.Errorf("Length expected: 5, got: %d", length)
	}

	is_empty := Head.IsEmpty()
	if is_empty {
		t.Error("List is empty")
	}
}

func TestListInsert() {

	if err := Head.InsertAfter(3, 20); err != nil {
		panic(err)
	} else {
		fmt.Println(Head.Length())
		Head.DumpList()
	}
	fmt.Println("Matched value index is", Head.FindMatchedValue(4))

	// Output:
	// 6
	// 0
	// 1
	// 2
	// 3
	// 20
	// 4
	// Matched value index is 5
}

func ExampleList_Delete() {

	if err := Head.Delete(6); err != nil {
		panic(err)
	} else {
		Head.DumpList()
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 20
}

func ExampleList_Find() {
	fmt.Println("data located is:", Head.Find(4))

	// Output:
	// data located is: 3

}

func ExampleList_Reverse() {
	revList = Head.Reverse()
	fmt.Println(revList.Length())
	fmt.Println("Data of original head node is:", Head.Data)
	fmt.Println("Data of reversal head node is:", revList.Data)
	fmt.Println("Data of original head.next node is:", Head.Next.Data)
	fmt.Println("Data of reversal head.next node is:", revList.Next.Data)
	revList.DumpList()

	// Output:
	// 5
	// Data of original head node is: 0
	// Data of reversal head node is: 20
	// Data of original head.next node is: <nil>
	// Data of reversal head.next node is: 3
	// 20
	// 3
	// 2
	// 1
	// 0
}

func ExampleList_QuickSort() {

	revList.QuickSort(nil, singlelist.ASCEND)
	revList.DumpList()

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 20
}

func ExampleList_SelectSort() {

	revList.SelectSort(singlelist.DESCEND)
	revList.DumpList()

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 20

}
