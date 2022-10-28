package gover

import "testing"

func TestVersion_UpToStr(t *testing.T) {
	v1, _ := NewVersion("0.1.22")

	if !v1.UpToStr("v0.1.50") {
		t.Errorf("got false, wanted true")
	}

	if v1.UpToStr("v0.1.21") {
		t.Errorf("got true, wanted false")
	}
}

func TestVersion_EqualOrHigherStr(t *testing.T) {
	v1, _ := NewVersion("0.1.22")

	if !v1.EqualOrHigherStr("0.1.21") {
		t.Errorf("got false, wanted true")
	}

	if v1.EqualOrHigherStr("0.1.99") {
		t.Errorf("got true, wanted false")
	}
}
