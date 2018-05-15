// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Dict impletments a bunch of operations which are pretty
// simliar to what python does. It is majorly based on
// the Go map type to support different API that will
// be supportive with data management with key-value pait.
// And data stored in dict would be applicable to any type.

package dict

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
)

// Any type for the dict keys and values.
type Any = interface{}

// Go list aligned with python style.
// Each element inside from the List will be treated
// as a value of the Dict despite any type it belongs.
type List = []Any

// Go dict aligned with python style.
// Initialization about a dict could be easy by using
// the 'Dict{}' through which it actually leverages
// map for the definition of the key-value elements.
// Considering the hash towards the key inside the map
// only the basic type(string, int) are commonly supported.
type Dict map[Any]Any

// Error types for different operations for Dict
var (
	ErrRemoveFromEmptyDict   = errors.New("Trying to remove element from empty dict")
	ErrUnsupportKeyTypeFound = errors.New("Unsupportive key type found")
	ErrValueNotExist         = errors.New("Value not exist")
)

// IsValidKeys will determine the any type come from interface{}
// is fine to be hashed or not. Currently the type like func()
// chan or even the other Go objects are seem to be wired if they
// are set as the key inside from the map[] so here just filter
// invalid syntax expression for the key.
func IsValidKeys(key Any) (err error) {
	err = ErrUnsupportKeyTypeFound
	switch key.(type) {
	case string:
		err = nil
	case byte:
		err = nil
	case float32, float64:
		err = nil
	case int, int8, int32, int64:
		err = nil
	case uint, uint16, uint32, uint64:
		err = nil
	}
	return
}

// Return new Dict object and using this object as the
// receiver by calling some python style method is just
// the nearly same as what python functions.
func NewDict() Dict {
	return make(Dict)
}

// DictFromKeys creates a new dictionary with keys from
// list and values set to defaultVal. Returns a new dict
// if the loading from the list is succeeded.
func FromKeys(keys List, defaultVal Any) (Dict, error) {
	newDict := NewDict()
	for _, key := range keys {
		if err := IsValidKeys(key); err != nil {
			return newDict, err
		} else {
			newDict[key] = defaultVal
		}
	}
	return newDict, nil
}

// Clear up all elements from the dictionary.
func (dict Dict) Clear() {
	for key := range dict {
		delete(dict, key)
	}
}

// HasKey returns true if key is in the dictionary, false otherwise.
func (dict Dict) HasKey(key Any) bool {
	if _, ok := dict[key]; ok {
		return true
	}
	return false
}

// IsEqual returns true if dicts are equal.
func (dict Dict) IsEqual(otherDict Dict) bool {
	return reflect.DeepEqual(dict, otherDict)
}

// Keys returns a list of the dictionary's keys, unordered.
func (dict Dict) Keys() List {
	list := make(List, len(dict))
	i := 0
	for key, _ := range dict {
		list[i] = key
		i++
	}
	return list
}

// Values returns a list of with the values belongs to the dict.
// These values will also be unordered.
func (dict Dict) Values() List {
	list := make(List, len(dict))
	i := 0
	for _, value := range dict {
		list[i] = value
		i++
	}
	return list
}

// Items returns an unordered list with the element of
// each key-value pairs. Saying, the result in the list
// will be [[key1, value1], [key2,value2],[key3,value3]..]
func (dict Dict) Items() []List {
	mList := []List{}
	for key, value := range dict {
		mList = append(mList, List{key, value})
	}
	return mList
}

// Pop returns value and remove the given key from the dictionary.
// If the given key is NOT in the dictionary return defaultVal.
// defaultVal should be same type as you expect to get.
func (dict Dict) Pop(key Any, defaultVal Any) (Any, error) {
	if len(dict) <= 0 {
		return defaultVal, ErrRemoveFromEmptyDict
	}

	if dict.HasKey(key) {
		val := dict[key]
		delete(dict, key)
		return val, nil
	}

	return defaultVal, nil
}

// PopItem returns and removes a random key-value pair from the dict.
// And the random elment returned will be set into a specified list.
func (dict Dict) PopItem() (List, error) {
	if len(dict) <= 0 {
		return List{}, ErrRemoveFromEmptyDict
	}

	// Get dict keys
	dictKeys := dict.Keys()

	// Return random key as string
	randKey := fmt.Sprintf("%v", dictKeys[rand.Intn(len(dictKeys))])

	list := make(List, 2)
	list = List{randKey, dict[randKey]}

	delete(dict, randKey)

	return list, nil

}

// Get returns value for the given key or defaultVal if key is NOT in
// the dictionary. defaultVal should be same type as you expect to get.
func (dict Dict) Get(key Any, defaultVal Any) Any {
	if dict.HasKey(key) {
		return dict[key]
	}
	return defaultVal
}

// Set a default value into a dict with the specfied key.
// Note if a value along with a key in the dict has already
// existed,then the corresponding pair will not be changed.
// Whereas if the value is not presented in that dict, then
// a new key-pair will go into the dict as well. Either way
// the default value of the second parameter will be returned.
func (dict Dict) SetDefault(key Any, defaultVal Any) (Any, error) {
	if dict.HasKey(key) {
		return dict[key], nil
	}
	if err := IsValidKeys(key); err != nil {
		return defaultVal, err
	} else {
		dict[key] = defaultVal
	}
	return defaultVal, nil
}

// Update updates the dictionary with the key-value pairs in the mDict
// dictionary replacing current values and adding new if found.
func (dict Dict) Update(mDict Dict) {
	for key, value := range mDict {
		dict[key] = value
	}
}
