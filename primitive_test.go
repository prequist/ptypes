package ptypes

import (
	pointer "ptypes/ptr"
	"testing"
)

// TestErasure tests creation from an interface handling type erasure.
// The interface box has a value that signifies if it could have been erased,
// and handles the erasure respectfully.
func TestErasure(t *testing.T) {
	value := "hello"
	ptr := pointer.FromInterface(value)
	deref, err := ptr.String()
	if err != nil {
		t.Error(err.Error())
	}
	if deref != value {
		t.Errorf("the string '%s' was not the input string '%s'!", deref, value)
	}
}

// TestString tests the conversion of string to pointer to string.
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

// TestInterface tests the boxing of objects through a more obtust `interface`
// returning back an interface, expecting the object returned to later
// then be asserted.
func TestInterface(t *testing.T) {
	value := "hello"
	ptr := pointer.FromInterface(value)
	deref, err := ptr.Interface()
	if err != nil {
		t.Error(err.Error())
	}
	str := deref.(string)
	if value != deref {
		t.Errorf("the string '%s' was not the input string '%s'!", deref, str)
	}
}
