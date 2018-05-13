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

func TestDelete(t *testing.T) {
	mList := list.BuildList()
	if len(mList) > 0 {
		t.Error("Unexpected list length")
	}

	if err := mList.Delete(0); err != nil {
		t.Log("Log", err)
	}

	mList.Append(2)
	if err := mList.Delete(0); err != nil {
		t.Error(err)
	}

	if len(mList) > 0 {
		t.Error("Unexpected list length after delete")
	}
}

func TestIndex(t *testing.T) {
	mList := list.BuildList(1, 2, 3)
	if index, err := mList.Index(1); err != nil || index != 0 {
		t.Errorf("Index:%d is not as expected, error is: %v\n", index, err)
	}

	if index, err := mList.Index(5); err != nil || index != -1 {
		t.Log(err, index)
	}
}

func TestInsert(t *testing.T) {
	mList := list.BuildList()
	mList.Insert(0, 1, 2, 3)
	if len(mList) != 3 && mList[0] != 1 {
		t.Errorf("Unexpected result of mList after insertion:%v with len:%d\n", mList, len(mList))
	}

	mList.Insert(0, 4)
	if len(mList) != 4 && mList[0] != 4 {
		t.Errorf("Unexpected result of mList after insertion:%v with len:%d\n", mList, len(mList))
	}
}

func TestIsEqual(t *testing.T) {
	mList := list.BuildList(1, 2, 3)
	nList := list.BuildList(1, 2, 4)
	sList := list.BuildList(1, 2, 3)

	if mList.IsEqual(nList) {
		t.Errorf("List:%v should not be the same as: %v.\n", mList, nList)
	}

	if !mList.IsEqual(sList) {
		t.Errorf("List:%v should be the same as: %v\n", mList, sList)
	}
}

func TestPop(t *testing.T) {
	mList := list.BuildList(1, 2, 3)
	if val, err := mList.Pop(); val != 3 || err != nil {
		t.Errorf("Value:%v populated from list: %v is not as expected, err is:%v\n", val, mList, err)
	}
}

func TestPopItem(t *testing.T) {
	mList := list.BuildList(1, 2, 3)
	if val, err := mList.PopItem(0); val != 1 || err != nil {
		t.Errorf("Value:%v populated from list is not mList[0]: %v, err is:%v\n", val, mList[0], err)
	}
}

func TestRemove(t *testing.T) {
	mList := list.BuildList(1, 2, 3)
	if err := mList.Remove(1); err != nil {
		t.Error(err, "value 1 is in the list\n")
	}

}

func TestReverse(t *testing.T) {
	mList := list.BuildList(1, 2, 3)
	mList.Reverse()
	var v1, v2, v3 int
	if v1 = mList[0].(int); int(v1) != 3 {
		t.Errorf("list after reverse is :%v\n", mList[0])
	}
	if v2 = mList[1].(int); int(v2) != 2 {
		t.Errorf("list after reverse is :%v\n", mList[1])
	}
	if v3 = mList[2].(int); int(v3) != 1 {
		t.Errorf("list after reverse is :%v\n", mList[2])
	}

	res := false
	if v1 == 3 && v2 == 2 && v3 == 1 {
		res = true
	}

	if !res {
		t.Errorf("List after reverse is: %v\n", mList)
	}
}

func TestString(t *testing.T) {
	mList := list.BuildList(1, 2, 3)
	if str := mList.String(""); str != "123" {
		t.Errorf("List join failed: %s\n", str)
	}
}

func TestSort(t *testing.T) {
	mList := list.BuildList(3, 2, 1)
	if res, err := mList.Sort(); err != nil {
		t.Error(err, res)
	} else {
		t.Log(res)
	}

	mList = list.BuildList("hello", "world", "abc", "c")
	if res, err := mList.Sort(); err != nil {
		t.Error(err, res)
	} else {
		t.Log(res)
	}

	mList = list.BuildList(3.0, 2.1, 1.0)
	if res, err := mList.Sort(); err != nil {
		t.Error(err, res)
	} else {
		t.Logf("%0.1f", res)
	}

	mList = list.BuildList(3, 2, 1.0)
	if res, err := mList.Sort(); err != nil {
		t.Error(err, res)
	} else {
		t.Logf("%v", res)
	}
}
