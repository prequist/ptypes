package ptypes

import (
	"reflect"
	"unsafe"
)

// Why do we have a bunch of functions with the same functionality?
// This is the closest we're getting to overloading here, utilising
// functions that do NOT support type erasure.
//
// Ideally, we'd want to use `interface{}`, however, there seems to be a
// sort of type erasure that ends up actually segfaulting the program.
// This is something that we're going to have to cope with, and figure out
// later.
//
// This is bypassed by using the `FromInterface` method which converts the value
// back into interface, then handles the type assertion.
//
// The statically typed functions `From` + Type, do not support
// type erasure, whereas `FromInterface` does.

// FromInt create a box from an integer.
func FromInt(i int) Box {
	return Box{
		Value: unsafe.Pointer(&i),
		OriginalType: reflect.TypeOf(i),
		HasErasure: false,
	}
}

// FromInt64 create a box from an int64
func FromInt64(i64 int64) Box {
	return Box{
		Value: unsafe.Pointer(&i64),
		OriginalType: reflect.TypeOf(i64),
		HasErasure: false,
	}
}

// FromString create a box fom a string.
func FromString(str string) Box {
	return Box{
		Value: unsafe.Pointer(&str),
		OriginalType: reflect.TypeOf(str),
		HasErasure: false,
	}
}

// FromUint create a box from a unsigned integer.
func FromUint(i uint) Box {
	return Box{
		Value: unsafe.Pointer(&i),
		OriginalType: reflect.TypeOf(i),
		HasErasure: false,
	}
}

// FromInterface create a box from an interface, supporting
// type erasure in the conversion.
func FromInterface(i interface{}) Box {
	wrapped := InterfaceAlias{Object: i}
	return Box{
		Value: unsafe.Pointer(&wrapped),
		OriginalType: reflect.TypeOf(wrapped),
		HasErasure: true,
	}
}