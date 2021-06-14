package reverseString

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	bs := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(bs)
	fmt.Printf("%v\n", string(bs))
}
