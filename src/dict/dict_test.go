package dict_test

import (
	."dict"
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {

	fmt.Println("this is testing")
	dict := NewDict()

	fmt.Println(dict.Get(5, 8))

	dict.SetDefault('0', 6)
	fmt.Println(dict.Get('0', 6))

}
