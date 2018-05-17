// Copyright 2018 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dict_test

import (
	"dict"
	"reflect"
	"testing"
)

func TestHasKey(t *testing.T) {
	mDict := dict.NewDict()
	mDict.SetDefault('0', 6)

	if !mDict.HasKey('0') {
		t.Error("Dict has not the key")
	}
}

func TestClear(t *testing.T) {
	mDict := dict.NewDict()
	mDict.SetDefault("hello", "world")

	mDict.Clear()

	if len(mDict) > 0 {
		t.Error("dict now should be cleaned")
	}
}

func TestIsEqual(t *testing.T) {
	mDict := dict.NewDict()
	nDict := dict.NewDict()

	mDict["hi"] = "ok"
	nDict.SetDefault("hi", "ok")

	if !mDict.IsEqual(nDict) {
		t.Error("mDict should be equal with nDict")
	}
}

func TestFromKeys(t *testing.T) {
	mDict, err := dict.FromKeys([]string{"a", "b", "c"}, "niuniu")

	if mDict["a"] != "niuniu" ||
		mDict["b"] != "niuniu" ||
		mDict["c"] != "niuniu" {
		t.Errorf("%v, %v, %v\n", mDict["a"], mDict["b"], mDict["c"])
	}

	if err != nil {
		t.Error(err)
	}
}

func TestKeys(t *testing.T) {
	mDict, _ := dict.FromKeys([]int{0, 1, 2}, 5)
	mList := mDict.Keys()
	for index, key := range mList {
		if mList[index] != key {
			t.Errorf("key not expected from dict.Keys():%v, %v\n", index, key)
		}
	}
}

func TestValues(t *testing.T) {
	mDict, _ := dict.FromKeys([]int{0, 1, 2}, 5)
	mList := mDict.Values()
	for index, value := range mList {
		if value != 5 {
			t.Errorf("value not expected from dict.Values():%v, %v\n", index, value)
		}
	}
}

func TestItems(t *testing.T) {
	mDict := dict.NewDict()
	for i := 0; i < 5; i++ {
		mDict[i] = i
	}

	mItems := mDict.Items()
	for index, item := range mItems {
		if reflect.TypeOf(item).Kind() != reflect.Slice {
			t.Errorf("items serialized is not expected:%v, %v\n", index, item)
		}
	}
}

func TestPop(t *testing.T) {
	mDict := dict.NewDict()
	value, err := mDict.Pop(1, 2)
	if err != nil {
		t.Log(value, err)
	}

	mDict[1] = 2
	value, err = mDict.Pop(1, 2)
	if err != nil {
		t.Log(value, err)
	} else {
		if len(mDict) > 0 {
			t.Errorf("Data not expected after pop: %v\n", mDict)
		}
	}
}

func TestPopItem(t *testing.T) {
	mDict := dict.NewDict()
	mList, _ := mDict.PopItem()
	if len(mList) > 0 {
		t.Errorf("length of the list from PopItem() is not as expected:%v\n", mList)
	}

	for i := 0; i < 5; i++ {
		mDict[i] = i
	}

	mList, _ = mDict.PopItem()
	if len(mDict) == 5 {
		t.Log(mList, mDict)
	}
}

func TestGet(t *testing.T) {
	mDict := dict.NewDict()
	mDict['a'] = 'A'
	value := mDict.Get('a', 'B')
	if value != 'A' {
		t.Errorf("Error to get the element:%v\n", value)
	}
}

func TestSetDefault(t *testing.T) {
	mDict := dict.NewDict()
	if value, err := mDict.SetDefault('a', 'b'); err != nil {
		t.Errorf("Value being set as default is:%v, %v\n", value, err)
	} else {
		t.Logf("Default value set as expected: %v\n", value)
	}
}

func TestUpdate(t *testing.T) {
	mDict := dict.NewDict()
	nDict := dict.NewDict()
	for i := 0; i < 5; i++ {
		mDict[i] = i
	}

	for j := 0; j < 5; j++ {
		nDict[j] = 5
	}

	mDict.Update(nDict)
	for k, v := range mDict {
		if v != 5 {
			t.Errorf("Dict update failed for: %v, %v\n", k, v)
		}
	}
}
