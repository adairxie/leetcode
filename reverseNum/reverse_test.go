package reverseNum

import (
	"testing"
	"math"
)

func TestReverseNum(t *testing.T) {
	res := reverse(math.MinInt32)
	t.Logf("result is:%d\n", res)
	t.Logf("max int32:%d\n", math.MaxInt32)
	t.Logf("min int32:%d\n", math.MinInt32)
}