// Copyright 2018 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"fmt"
	"list"
)

var mList list.List

func ExampleList_MakeList() {
	mList = list.MakeList(5)
	fmt.Println(mList)

	//Output:
	//[<nil> <nil> <nil> <nil> <nil>]
}

func ExampleList_BuildList() {
	mList = list.BuildList()
	fmt.Println(mList)

	//Output:
	//[1,2,3,4]
}

func ExampleList_InitList() {
	sample_list := list.MakeList(6)
	err := sample_list.InitList(7,2,999, 0, 4)
	fmt.Println("err is ", err, ", sample list is: ", sample_list)
	//Output:
	//error
}

func ExampleList_Append() {
	mSlice := []string{"hi", "world"}
	mList.Append(mSlice)
	fmt.Println(mList)

	//Output:
	//[1 2 3]
}

func ExampleList_Extend() {
	if err := mList.Extend(7, []int{1,2,3}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mList)
	}

	//Output:
	//[1 2 3 1 2 3]

}
