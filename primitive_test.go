package ptypes

import (
	"testing"
)

// TestStringErasure tests creation from an interface handling type erasure
// utilising a string.
// The interface box has a value that signifies if it could have been erased,
// and handles the erasure respectfully.
func TestStringErasure(t *testing.T) {
	value := "hello"
	ptr := FromInterface(value)
	deref, err := ptr.String()
	if err != nil {
		t.Error(err.Error())
	}
	if deref != value {
		t.Errorf("the string '%s' was not the input string '%s'!", deref, value)
	}
}

// TestIntErasure tests erasure and conversion from an interface-typed box
// back into an integer.
func TestIntErasure(t *testing.T) {
	value := 1
	ptr := FromInterface(value)
	deref := ptr.IntBox().Int()
	if deref == nil {
		t.Error("the interface value was not an integer!")
	}
	if *deref != value {
		t.Errorf("the number '%d' was not the input number '%d'!", *deref, value)
	}
}

// TestString tests the conversion of string to pointer to string.
func TestString(t *testing.T) {
	str := "hello"
	ptr := FromString(str)
	deref, err := ptr.String()
	if err != nil {
		t.Error(err.Error())
		return
	}
	if deref != str {
		t.Errorf("the string '%s' was not the input string '%s'!", deref, str)
	}
}

// TestUint tests that a uint pointer can convert back into the uint
// specified at the box creation.
func TestUint(t *testing.T) {
	value := uint(10)
	ptr := FromUint(value).IntBox()
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
	ptr := FromInt(value).IntBox()
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
	ptr := FromInterface(value)
	deref := ptr.Interface()
	str := deref.(string)
	if value != deref {
		t.Errorf("the string '%s' was not the input string '%s'!", deref, str)
	}
}
