package cubeutil_test

import "github.com/cubex/potens-go/cubeutil"
import "testing"

func TestAcronym(t *testing.T) {
	if cubeutil.Acronym("A Really Long Acronym Goes Here", 3) != "ARL" {
		t.Error("Limited Long Acronym Failed")
	}
	if cubeutil.Acronym("A Really Long Acronym Goes Here", 0) != "ARLAGH" {
		t.Error("Long Acronym Failed")
	}
	if cubeutil.Acronym("Cubex Platform System", 3) != "CPS" {
		t.Error("Length Matching Size Acronym Failed")
	}
	if cubeutil.Acronym("Cubex Platform", 3) != "CPL" {
		t.Error("Extended End Acronym Failed")
	}
	if cubeutil.Acronym("Cubex Platform", 0) != "CP" {
		t.Error("Short Acronym Failure")
	}
	if cubeutil.Acronym("Cubex", 3) != "CUB" {
		t.Error("Single Word Acronym with Length Failed")
	}
	if cubeutil.Acronym("Cubex", 0) != "C" {
		t.Error("Single Word Acronym Failed")
	}
}
