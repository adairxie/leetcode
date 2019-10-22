package mySqrt

import "testing"

func TestMySqrt(t *testing.T) {
	res := mySqrt(16)
	t.Logf("res:%d\n", res)
}