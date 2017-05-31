package cubeutil_test

import (
	"testing"
	"github.com/cubex/cubeutil-go"
)

func TestRandomAlphaNum(t *testing.T) {
	if len(cubeutil.RandomAlphaNum(1)) != 1 {
		t.Fail()
	}
	if len(cubeutil.RandomAlphaNum(10)) != 10 {
		t.Fail()
	}
	if len(cubeutil.RandomAlphaNum(100)) != 100 {
		t.Fail()
	}
	if len(cubeutil.RandomAlphaNum(1000)) != 1000 {
		t.Fail()
	}
}
