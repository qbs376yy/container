// Copyright 2018 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dict_test

import (
	"dict"
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
	mDict := dict.NewDict()
	mDict = dict.FromKeys([]string{"a", "b", "c"}, "niuniu")

	if mDict["a"] != "niuniu" &&
		mDict["b"] != "niuniu" &&
		mDict["c"] != "niuniu" {
		t.Errorf("%v, %v, %v\n", mDict["a"], mDict["b"], mDict["c"])
	}

}
