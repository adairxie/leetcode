package myAtoi

import "testing"

func TestMyAtoI(t *testing.T) {
	res := myAtoi("123-")
	t.Logf("res:%d\n", res)
}