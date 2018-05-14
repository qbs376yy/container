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
