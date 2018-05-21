// Copyright 2018 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// syncmap implements a map whose data stored inside
// could be able to be manipulated with multi-thread
// safe synchonrization. It provides some of the most
// common operations towards the synchonrized map with
// the control of RWMutex.

package syncmap_test

import (
	"sync"
	"syncmap"
	"testing"
)

var sm *syncmap.SyncMap
var wg sync.WaitGroup

func TestCreateSyncMap(t *testing.T) {
	sm = syncmap.NewSyncMap()

	if sm == nil {
		t.Errorf("syncmap created is: %v\n", sm)
	}
}

func TestPutAndGet(t *testing.T) {

	// Emulate a multiple threads scenario and pass each value onto the map
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) {
			sm.Put(x, x)
			defer wg.Done()
		}(i)
	}
	// Wait all of the threads in the group to exit
	wg.Wait()

	for i := 0; i < 5; i++ {
		if sm.Get(i) != i {
			t.Errorf("sync map with value is: %v\n", sm)
		}
	}

}

func TestKeys(t *testing.T) {
	keys := sm.Keys()
	if len(keys) != 5 {
		t.Errorf("length of the keys in sync map is: %v\n", len(keys))
	}
}

func TestEach(t *testing.T) {

	sm.Each(func(k interface{}, v interface{}) {
		t.Logf("value of index %v in the map is %v\n", k, v)
	})
}

func TestDelete(t *testing.T) {
	sm.Delete(4)

	if sm.Get(4) == 4 {
		t.Errorf("data after delete is still existing: %v\n", sm)
	}
}
