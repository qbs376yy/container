package list

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// List base structure, it is able to store any type of element.
type List []interface{}

// Return new List with specified length, which actually is a slice.
func NewList(length int) List {
	return make(List, length)
}

// Return new List with specified length and capacity.
func MakeList(length int, cap int) List {
	return make(List, length, cap)
}

// Error type for the operations towards the List
var (
	ErrRemoveFromEmptyList = errors.New("Error to remove element from an empty list")
	ErrAppendExistValueIntoList = errors.New("Error to append an existed element into the list")
	ErrExtendWithNoList = errors.New("Error to extend non-list values into a list")
	ErrIndexNotFound = errors.New("Error to locate the index of the specified value")
)


// Adds elements to the end of the specified list.
// Note since the builtin append() will return a new
// list when the length or capacity is not sufficient.
// In this case the reciver should be used with the
// pointer object rather than the 'List' object itself.
func (list *List) Append(values ...interface{}) {
	*list = append(*list, values...)
}

// Adds an element to the end of the list if it's not
// already in the list.
func (list *List) AppendIfNotExists(value interface{}) error {
	for _, ele := range *list {
		if ele == value {
			return ErrAppendExistValueIntoList
		}
	}
	*list = append(*list, value)
	return nil
}

// Returns the times of the caculated numbers in the list.
func (list List) Count(value interface{}) int {
	count := 0
	for _, listValue := range list {
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

// Extend one list with the contents of the other list.
func (list *List) Extend(values interface{}) error {
	switch values.(type) {
		case List:
			for _, value := range otherList {
				*list = append(*list, value)
			}
			return nil
		default:
			return ErrExtendWithNoList
	}
}

// Returns the index of the item in the list within the value of val.
// Note this will only seek for the index of first item in the list.
// Will returned with -1 if there is no specified item has been found.
func (list List) Index(val interface{}) (int, error) {
	for index, listValue := range list {
		if listValue == val {
			return index, nil
		}
	}
	fmt.Println("%v does not exist in the list, return index of -1 instead", val)
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
func (list List) IsEqual(otherList List) bool {
	return reflect.DeepEqual(list, otherList)
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
	errorString := fmt.Sprintf("%v is not in list", val)
	if len(*list) > 0 {
		for index, listValue := range *list {
			if listValue == val {
				(*list).Delete(index)
				return nil
			}
		}
	}
	return errors.New(errorString)
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

// Sort the list in place ordering elements from smallest to largest.
//func (list *List) Sort() {
//}

// String returns list values as string
//		l := listdict.List{"one", 2, "three"}
// 		l.String() => "one, 2, three"
func (list List) String() string {
	var out []string
	for _, val := range list {
		out = append(out, fmt.Sprintf("%v", val))
	}
	return strings.Join(out, ", ")
}
