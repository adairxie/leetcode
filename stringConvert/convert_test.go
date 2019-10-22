package stringConvert

import (
	"testing"
)

func TestConvert(t *testing.T) {
	res := convert("LE", 1)
	t.Logf("reuslt is:%s\n", res)
}