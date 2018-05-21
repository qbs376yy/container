// Copyright 2018 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syncmap_test

import (
	"fmt"
	"syncmap"
)

func ExampleSyncMap_Each() {
	sm := syncmap.NewSyncMap()
	for i := 0; i < 3; i++ {
		sm.Put(i, i)
	}

	sm.Each(func(k interface{}, v interface{}) {
		fmt.Printf("key: %v, value: %v\n", k, v)
	})
	// Output:
	// key: 0, value: 0
	// key: 1, value: 1
	// key: 2, value: 2
}
