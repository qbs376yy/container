// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"list"
	"testing"
)

func TestMakeList(t *testing.T) {
	mList := list.MakeList(5)

	if mList.IsNilList {
		t.Errorln("list is newly created and expected with true but return:", length)
	}
}
