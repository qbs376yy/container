// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package single_test


func ExampleList() {

	Head := InitList()
	Cur := Head
	for i := 0; i < 5; i++ {
		Cur.AddNode(i)
	}

	Head.DumpList()

	fmt.Println(Head.Length())
	fmt.Println(Head.IsEmpty())
	if err := Head.InsertAfter(3, 20); err != nil {
		panic(err)
	} else {
		fmt.Println(Head.Length())
		Head.DumpList()
	}
	fmt.Println("matched value index is", Head.FindMatchedValue(4))

	if err := Head.Delete(6); err != nil {
		panic(err)
	} else {
		Head.DumpList()
	}

	fmt.Println("matched value index is:", Head.FindMatchedValue(2))

	fmt.Println("data located is:", Head.Find(4))
	revList := Head.Reverse()
	fmt.Println(revList.Length())
	revList.DumpList()

	fmt.Println(Head.Next)

	revList.QuickSort(0, revList.Length())
	revList.DumpList()
}
