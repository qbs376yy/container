// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"fmt"
	"list"
	"testing"
)

func TestMakeList(t *testing.T) {
	mList := list.MakeList(5)

	if !mList.IsNilList() {
		t.Error("list is newly created and should be nil list", mList)
	}
}

func TestInitList(t *testing.T) {
	mList := list.MakeList(5)
	mList.InitList(1, 2, 3, 4, 5)
	if mList.IsNilList() {
		t.Error("list is failed to be initialized and should be", mList)
	}

	if len(mList) != 5 {
		t.Errorf("length of mList is expecting with:%d\n", len(mList))
	}

	for index, value := range mList {
		if index+1 != value {
			t.Errorf("Init List error, sequence corruptedi\n")
		}
	}
}

func TestBuildList(t *testing.T) {
	mSlice := []string{"hello", "world"}
	var mList list.List = list.BuildList(mSlice)

	if len(mList) != 1 {
		t.Errorf("length of sample list built via BuildList is:%d\n", len(mList))
	}

	if s := fmt.Sprintf("%v", mList[0]); s != "[hello world]" {
		t.Errorf("Build List error and the sample_list[0] is %s\n", s)
	}
}

func TestAppend(t *testing.T) {
	mList := list.MakeList(5)
	mList.InitList(1, 2, 3, 4, 5)
	mList.Append(6, 7)

	if mList[5] != 6 && mList[6] != 7 {
		t.Errorf("List after append is not as expected with mList[5]=%d, mList[6]=%d\n", mList[5], mList[6])
	}

	if len(mList) != 7 {
		t.Errorf("List length after append is:%d\n", len(mList))
	}
}

func TestExtend(t *testing.T) {
	mList := list.MakeList(5)
	mList.InitList(1, 2, 3, 4, 5)
	mList.Extend(6, []int{1, 2, 3})

	if mList[6] != 1 && mList[7] != 2 && mList[8] != 3 {
		t.Error("List after extend is not as expected", mList)
	}
}

func TestAppendIfNotExists(t *testing.T) {
	mList := list.BuildList(3)
	if err := mList.AppendIfNotExists(3); err != nil {
		t.Skip("Skip", err)
	}

	if err := mList.AppendIfNotExists(2); err != nil {
		t.Fatal(err)
	}
}

func TestCount(t *testing.T) {
	mList := list.BuildList(1, 2, 2, 2)
	if mList.Count(2) != 3 {
		t.Errorf("The count of value '2' is not as expected with times of: %d\n", mList.Count(2))
	}
}
