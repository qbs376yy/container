// Copyright 2018 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dict_test

import (
	"dict"
	"fmt"
)

func ExampleDict_NewDict() {
	mDict := dict.NewDict()
	mDict.SetDefault(rune('0'), 6)

	if !mDict.HasKey('0') {
		fmt.Println("Dict has not the key")
	} else {
		fmt.Println(mDict)
	}

	// Output:
	// map[48:6]
}
