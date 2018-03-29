// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package singlelist_test

import (
	_ "fmt"
	"singlelist"
	"testing"
)

func CreateList(n int) (head *singlelist.List) {
	head = singlelist.InitList()
	cur := head
	for i := 0; i < n; i++ {
		cur.AddNode(i)
	}

	return
}

func TestCreateList(t *testing.T) {
	head := CreateList(5)

	length := head.Length()
	if length != 5 {
		t.Errorf("Length expected: 5, got: %d", length)
	}

	is_empty := head.IsEmpty()
	if is_empty {
		t.Error("List is empty")
	}
}

func TestListInsert(t *testing.T) {
	head := CreateList(6)

	if err := head.InsertAfter(3, 20); err != nil {
		t.Fatalf("Error occured during the insertion in the list: %s", err)
	} else {
		if length := head.Length(); length != 7 {
			t.Errorf("Length expected: 7, got: %d", length)
		}

		if index := head.FindMatchedValue(20); index != 4 {
			t.Errorf("index of value 20 expected: 4, got: %d", index)
		}
	}

	if err := head.InsertBefore(3, 20); err != nil {
		t.Fatalf("Error occured during the insertion in the list: %s", err)
	} else {
		if length := head.Length(); length != 8 {
			t.Errorf("Length expected: 8, got: %d", length)
		}

		if index := head.FindMatchedValue(20); index != 4 {
			t.Errorf("index of value 20 expected: 4, got: %d", index)
		}
	}
}

func TestListDelete(t *testing.T) {
	head := CreateList(2)

	if err := head.Delete(1); err != nil {
		t.Fatalf("Error occured during the deletion: %s", err)
	} else {
		if length := head.Length(); length != 0 {
			t.Errorf("Length expected: 0, got: %d", length)
		}
	}
}

func TestListFind(t *testing.T) {
	head := CreateList(5)

	if index := head.Find(4); index != 3 {
		t.Errorf("Index expected: 3, got: %d", index)
	}
}

func TestListReverse(t *testing.T) {
	head := CreateList(3)

	tail := head.Reverse()
	if length := tail.Length(); length != 3 {
		t.Errorf("Length reversed expected: 3, got :%d", length)
	}

	if data := tail.Data; data != 2 {
		t.Errorf("Data of reversal head node expected: 2, got: %d", data)
	}

	if data := head.Next.Data; data != nil {
		t.Errorf("Data of original head.next node after reverse expected: <nil>, got: %d", data)
	}

	if data := tail.Next.Data; data != 1 {
		t.Errorf("Data of reversal head.next node expected: 1, got: %d", data)
	}
}

func TestListQuickSort(t *testing.T) {
	head := CreateList(5)

	head.QuickSort(nil, singlelist.ASCEND)
	if data := head.Data; data != 0 {
		t.Errorf("first smallest data after sort expected: 0, got: %d", data)
	}

	if data := head.Next.Data; data != 1 {
		t.Errorf("Second smallest data after sort expected: 1, got: %d", data)
	}

	head.QuickSort(nil, singlelist.DESCEND)
	if data := head.Data; data != 4 {
		t.Errorf("First largest data after descend sort expected: 4, got: %d", data)
	}

	if data := head.Next.Data; data != 3 {
		t.Errorf("Second largest data after descend sort expected: 3, got: %d", data)
	}

}

func TestListSelectSort(t *testing.T) {
	head := CreateList(5)

	head.SelectSort(singlelist.ASCEND)
	if data := head.Data; data != 0 {
		t.Errorf("first smallest data after sort expected: 0, got: %d", data)
	}

	if data := head.Next.Data; data != 1 {
		t.Errorf("Second smallest data after sort expected: 1, got: %d", data)
	}

	head.SelectSort(singlelist.DESCEND)
	if data := head.Data; data != 4 {
		t.Errorf("First largest data after descend sort expected: 4, got: %d", data)
	}

	if data := head.Next.Data; data != 3 {
		t.Errorf("Second largest data after descend sort expected: 3, got: %d", data)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	head := CreateList(10000)

	for i := 0; i < b.N; i++ {
		head.QuickSort(nil, singlelist.ASCEND)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	head := CreateList(10000)

	for i := 0; i < b.N; i++ {
		head.SelectSort(singlelist.ASCEND)
	}

}
