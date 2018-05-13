// Copyright 2018 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// List impletments a bunch of operations which are pretty
// alike what other languages already made. It is based on
// the Go slice data type to support different API that will
// be feasible to be exported for the data management or
// data organization. And the data stores the List would be
// handled with any type.

package list

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Error type for the operations towards the List
var (
	ErrRemoveFromEmptyList      = errors.New("Error to remove element from an empty list")
	ErrAppendExistValueIntoList = errors.New("Error to append an existed element into the list")
	ErrExtendWithNoList         = errors.New("Error to extend non-list values into a list")
	ErrIndexNotFound            = errors.New("Error to locate the index of the specified value")
	ErrListNotNew               = errors.New("Error to init a list that is not newly created")
	ErrListNotSupportSort       = errors.New("Error to sort a list that does not support to")
)

// List is based on the low layer slice,
// it is able to store any type of element.
type List []interface{}

// Return new List with specified length, which actually is a slice.
func MakeList(length int) List {
	return make(List, length)
}

// Return new List with specified length and capacity.
func MakeListWithCap(length int, capacity int) List {
	return make(List, length, capacity)
}

// Form the data into a list as required, all of the data
// pass from the caller would be stored into the list.
func BuildList(values ...interface{}) List {
	var mList List
	for _, value := range values {
		mList = append(mList, value)
	}
	return mList
}

// Determine a given list is with all <nil> value stored.
func (list List) IsNilList() bool {
	var res bool = true

	for _, value := range list {
		if value != nil {
			res = false
		}
	}

	return res
}

// Initialize an existing list which is created by MakeList()
// with the values that are needed to be restored into it.
func (list *List) InitList(values ...interface{}) error {
	if !list.IsNilList() {
		return ErrListNotNew
	}

	if len(*list) >= len(values) {
		copy((*list)[0:], values)
	} else {
		copy((*list)[0:], values[0:len(*list)])
		for _, remaining_value := range values[len(*list):] {
			*list = append(*list, remaining_value)
		}
	}
	return nil
}

// Adds elements to the end of the specified list.
// Note since the builtin append() will return a new
// list when the length or capacity is not sufficient.
// In this case the receiver should be used with the
// pointer object rather than the 'List' object itself.
// This will ensure the receiver pointer refer to the new
// buffer with cap and length expanded. And as a return
// the original slice shares the same address with its
// copy(a.k.a: receiver pointer *List in this case).
func (list *List) Append(values ...interface{}) error {
	if len(values) == 0 {
		return nil
	}

	if list.IsNilList() {
		return list.InitList(values...)
	} else {
		*list = append(*list, values...)
	}
	return nil
}

// Extend one list with the contents of the other list.
func (list *List) Extend(values ...interface{}) error {
	if len(values) == 0 {
		return nil
	}

	if list.IsNilList() {
		return list.InitList(values)
	}

	for _, element := range values {
		rtype := reflect.TypeOf(element)
		rvalue := reflect.ValueOf(element)
		switch rtype.Kind() {
		case reflect.Slice:
			// Thanks to the discussion from here:
			// https://stackoverflow.com/questions/14025833/range-over-interface-which-stores-a-slice
			// And https://github.com/golang/go/wiki/InterfaceSlice
			// That we cannot range over the type of interface{} .i.e: reflect.Value
			// In which case, we need leverage the help of reflec package to extend each value.
			for i := 0; i < rvalue.Len(); i++ {
				*list = append(*list, rvalue.Index(i).Interface())
			}
		default:
			*list = append(*list, element)
		}
	}
	return nil
}

// Adds an element to the end of the list if it's not
// already in the list. Likewise, should use pointer as
// the receiver.
func (list *List) AppendIfNotExists(value interface{}) error {
	for _, element := range *list {
		if element == value {
			return ErrAppendExistValueIntoList
		}
	}
	*list = append(*list, value)
	return nil
}

// Returns the times of the caculated numbers in the list.
func (list *List) Count(value interface{}) int {
	count := 0
	for _, listValue := range *list {
		if listValue == value {
			count++
		}
	}
	return count
}

// Removes element from the list with given index.
func (list *List) Delete(index int) error {
	if len(*list) <= 0 {
		return ErrRemoveFromEmptyList
	}

	length := len(*list)

	copy((*list)[index:], (*list)[index+1:])
	(*list)[length-1] = nil
	*list = (*list)[:length-1]
	return nil
}

// Returns the index of the item in the list within the value of val.
// Note this will only seek for the index of first item in the list.
// Will returned with -1 if there is no specified item has been found.
func (list *List) Index(val interface{}) (int, error) {
	for index, listValue := range *list {
		if listValue == val {
			return index, nil
		}
	}
	return -1, ErrIndexNotFound
}

// Insert an element at a given position. If the position exceeds to
// the end of the list, then append the element into the end.
func (list *List) Insert(index int, values ...interface{}) {
	if len(*list) > index {
		for i := 0; i < len(values); i++ {
			*list = append(*list, 0)
		}
		copy((*list)[index+len(values):], (*list)[index:])
		copy((*list)[index:], values)
	} else {
		*list = append(*list, values...)
	}
}

// IsEqual returns true if lists are equal.
func (list *List) IsEqual(otherList List) bool {
	return reflect.DeepEqual(*list, otherList)
}

// Remove and returns the last element in the list.
func (list *List) Pop() (interface{}, error) {
	if len(*list) <= 0 {
		return nil, ErrRemoveFromEmptyList
	}

	listLen := len(*list)
	val := (*list)[listLen-1]
	(*list).Delete(listLen - 1)

	return val, nil
}

// Remove and returns the element at the given position in the list.
func (list *List) PopItem(index int) (interface{}, error) {
	if len(*list) <= 0 {
		return nil, ErrRemoveFromEmptyList
	}

	val := (*list)[index]
	(*list).Delete(index)

	return val, nil
}

// Remove the first element from the list whose value matches the given value.
// Error if no match is found.
func (list *List) Remove(val interface{}) error {
	if len(*list) > 0 {
		for index, listValue := range *list {
			if listValue == val {
				(*list).Delete(index)
				return nil
			}
		}
	}
	return ErrRemoveFromEmptyList
}

// Reverse the elements of the list in place.
func (list *List) Reverse() {
	if len(*list) > 0 {
		maxIndex := len(*list) - 1
		for index := 0; index < (maxIndex/2)+1; index++ {
			(*list)[index], (*list)[maxIndex-index] =
				(*list)[maxIndex-index], (*list)[index]
		}
	}
}

// To rebuild the list with values from a slice
func formListFromSlice(slice interface{}) List {
	var mList List
	rt := reflect.TypeOf(slice)
	rv := reflect.ValueOf(slice)
	if rt.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			mList = append(mList, rv.Index(i).Interface())
		}
	}
	return mList
}

// Sort the list as needed, currently only the data type with
// int, float64, string is support to sort in a given slice.
func (list *List) Sort() (List, error) {
	if len(*list) == 0 {
		return *list, ErrListNotSupportSort
	}

	var mIntSlice sort.IntSlice
	var mFloat64Slice sort.Float64Slice
	var mStringSlice sort.StringSlice

	for _, value := range *list {
		switch (value).(type) {
		case int:
			mIntSlice = append(mIntSlice, value.(int))
		case float64:
			mFloat64Slice = append(mFloat64Slice, value.(float64))
		case string:
			mStringSlice = append(mStringSlice, value.(string))
		default:
			fmt.Printf("Invalid data type detected to sort: %v\n", value)
			return *list, ErrListNotSupportSort
		}
	}

	if len(mIntSlice) > 0 && len(*list) == len(mIntSlice) {
		mIntSlice.Sort()
		mList := formListFromSlice(mIntSlice)
		return mList, nil
	}

	if len(mFloat64Slice) > 0 && len(*list) == len(mFloat64Slice) {
		mFloat64Slice.Sort()
		mList := formListFromSlice(mFloat64Slice)
		return mList, nil
	}

	if len(mStringSlice) > 0 && len(*list) == len(mStringSlice) {
		mStringSlice.Sort()
		mList := formListFromSlice(mStringSlice)
		return mList, nil
	}

	return *list, ErrListNotSupportSort
}

// String returns list values as string
func (list *List) String(sep string) string {
	var out []string
	for _, val := range *list {
		out = append(out, fmt.Sprintf("%v", val))
	}
	return strings.Join(out, sep)
}
