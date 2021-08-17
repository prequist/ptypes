package ptypes

import (
	ptr2 "ptypes/ptr"
	e "ptypes/ptr/prim"
	"testing"
)

func TestErasure (t *testing.T) {
	str := e.String.Name()
	if str != "string" {
		t.Error()
	}
}

func TestString(t *testing.T) {
	str := "hello"
	ptr := ptr2.FromString(str)
	deref, err := ptr.String()
	if err != nil {
		t.Error(err.Error())
		return
	}
	if deref != str {
		t.Errorf("the string '%s' was not the input string '%s'!", deref, str)
	}
}

func TestUint(t *testing.T) {
	value := uint(10)
	ptr := ptr2.FromUint(value).IntBox()
	deref := ptr.Uint()
	if deref != value {
		t.Fail()
	}
}
