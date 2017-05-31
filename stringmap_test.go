package cubeutil_test

import (
	"testing"
	"github.com/cubex/cubeutil-go"
)

func TestGetMapValue(t *testing.T) {
	useMap := make(map[string]string)
	useMap["found"] = "Found"
	if cubeutil.GetStringMapValue(useMap, "found", "Not Found") != "Found" {
		t.Fail()
	}
	if cubeutil.GetStringMapValue(useMap, "not_found", "Not Found") != "Not Found" {
		t.Fail()
	}
}
