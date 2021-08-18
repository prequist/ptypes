package ptypes

import (
	"reflect"
	"unsafe"
)

// BuiltinType a type alias for a golang builtin type.
type BuiltinType struct {
	// The nested reflect type.
	reflect.Type
	// The type's conversion method.
	Method conversionMethod
}

// InterfaceAlias is an alias type utilised for boxing
// the extremely primitive `interface{}` type.
type InterfaceAlias struct {
	// The boxed value.
	Object interface{}
}

// These exported variables are
// builtin types that can be associated with
// ConversionMethod s to be used in pointer -> value conversion.
var (
	// mock mock variable to take the address off.
	mock = 1
	String        = of("")
	Int           = of(1)
	Uint          = of(uint(1))
	Int64         = of(int64(1))
	Int32         = of(int32(1))
	Int16         = of(int16(1))
	Bool          = of(true)
	Int8          = of(int8(1))
	Uint8         = of(uint8(1))
	Uint16        = of(uint16(1))
	Uint32        = of(uint32(1))
	Uint64        = of(uint64(1))
	Uintptr       = of(uintptr(1))
	Float32       = of(float32(1))
	Float64       = of(float32(1))
	Complex64     = of(float32(1))
	Complex128    = of(float32(1))
	Array         = of(float32(1))
	Func          = of(func() {})
	Map           = of(make(map[int]int, 0))
	Slice         = of([]int{})
	Struct        = of(struct{}{})
	UnsafePointer = of(unsafe.Pointer(&mock))
	Interface     = of(InterfaceAlias{})
)

// Create a builtin type of a value, if the associated
// conversion method exists.
func of(i interface{}) *BuiltinType {
	// Get the reflect type.
	typeOf := reflect.TypeOf(i)
	// Try to get the associated conversion type.
	method := ConversionTypeMap[typeOf.Name()]
	// Method does not exist, do not continue.
	if method == nil {
		return nil
	}
	// Create the builtin type and register it.
	builtin := BuiltinType{typeOf, method}
	Builtins[typeOf.Name()] = builtin
	// Return a reference.
	return &builtin
}

// Convert a box's value into this type.
func (builtin *BuiltinType) Convert(box Box) interface{} {
	// Call the conversion method on the box.
	return builtin.Method(box)
}
