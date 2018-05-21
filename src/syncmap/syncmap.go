// Copyright 2018 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// syncmap implements a map whose data stored inside
// could be able to be manipulated with multi-thread
// safe synchonrization. It provides some of the most
// common operations towards the synchonrized map with
// the control of RWMutex.

package syncmap

import (
	"errors"
	"sync"
)

// Define a type to indicate what data this container supports
type Any = interface{}

// The base of this container consists of a rwmutex and
// the real data structure with format of 'key-value' pair.
type SyncMap struct {
	rw   *sync.RWMutex
	data map[Any]Any
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		rw:   new(sync.RWMutex),
		data: make(map[Any]Any),
	}
}

// Keys will return the list of all the key that
// is the index of each data
func (sm *SyncMap) Keys() []Any {
	var list []Any
	for k := range sm.data {
		list = append(list, k)
	}
	return list
}

// Put will store a key-value pair data into the map.
// Only one thread is allowed to update the map in
// each time.
func (sm *SyncMap) Put(k, v Any) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	sm.data[k] = v
}

// Get will acquire the value to the caller by passing
// a given key. Any possible threads are able to access
// this map and get the value they want to do.
func (sm *SyncMap) Get(k Any) Any {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	if v, ok := sm.data[k]; ok {
		return v
	} else {
		return nil
	}
}

// Delete will remove the member in the map with the
// given key provides, And this is also controlled by
// the mutex from which only one thread is allowed to.
func (sm *SyncMap) Delete(k Any) error {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	if _, ok := sm.data[k]; ok {
		delete(sm.data, k)
		return nil
	} else {
		return errors.New("Try to delete the non-existing value in sync map")
	}
}

// Each will provide any possible opeartion with the callback
// function passing into the parameter. This will range over
// the synchronized map one by one and handle the potential
// key-value to complete the desired process the callback requires.
func (sm *SyncMap) Each(cb func(Any, Any)) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	for k, v := range sm.data {
		cb(k, v)
	}
}
