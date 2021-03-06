package cubeutil_test

import (
	"testing"
	"github.com/cubex/cubeutil-go"
)

func TestCreateID(t *testing.T) {
	id := cubeutil.CreateID("kh w-ekhwlkeh  --   fklwhelkfh£$*%(^£@^$I!@&$! wjg")
	if id != "kh-w-ekhwlkeh-fklwhelkfh-i-wjg" {
		t.Fail()
	}
}

func TestValidateID(t *testing.T) {
	if cubeutil.ValidateID(cubeutil.CreateID("kh w-ekhwlkeh  --   fklwhelkfh£$*%(^£@^$I!@&$! wjg")) != nil {
		t.Fail()
	}
	if cubeutil.ValidateID(cubeutil.CreateID("abcdef")) != nil {
		t.Fail()
	}
}
