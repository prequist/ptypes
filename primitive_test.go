package ptypes

import (
	pointer "ptypes/ptr"
	"ptypes/ptr/prim"
	"testing"
)

func TestErasure(t *testing.T) {
	str := prim.String.Name()
	if str != "string" {
		t.Error()
	}
}

func TestString(t *testing.T) {
	str := "hello"
	ptr := pointer.FromString(str)
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
	ptr := pointer.FromUint(value).IntBox()
	deref := ptr.Uint()
	if deref == nil {
		t.Fail()
	}
	if *deref != value {
		t.Fail()
	}
}

func TestInt(t *testing.T) {
	value := 10
	ptr := pointer.FromInt(value).IntBox()
	deref := ptr.Int()
	if deref == nil {
		t.Fail()
	}
	if *deref != value {
		t.Fail()
	}
}
