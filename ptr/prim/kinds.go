package prim

import (
	"reflect"
	"unsafe"
)

// BuiltinType a type alias for a golang builtin type.
type BuiltinType struct {
	reflect.Type
	Method conversionMethod
}

var (
	// mock mock variable to take the address off.
	mock = 1
	// String type of built-in string.
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
)

func of(i interface{}) BuiltinType {
	typeOf := reflect.TypeOf(i)
	method := ConversionTypeMap[typeOf.Name()]
	builtin := BuiltinType{typeOf, method}
	Builtins[typeOf.Name()] = builtin
	return builtin
}

func (builtin *BuiltinType) Convert(pointer unsafe.Pointer) interface{} {
	return builtin.Method(pointer)
}
