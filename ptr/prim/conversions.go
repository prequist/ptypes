package prim

import (
	"unsafe"
)

var (
	ConversionTypeMap = make(map[string]conversionMethod, 0)
	Builtins    = make(map[string]BuiltinType, 0)

	StringMethod = ConversionMethod("string", func(ptr unsafe.Pointer) interface{} {
		return *(*string)(ptr)
	})
	UintMethod = ConversionMethod("uint", func(ptr unsafe.Pointer) interface{} {
		return (*uint)(ptr)
	})
	IntMethod = ConversionMethod("int", func(ptr unsafe.Pointer) interface{} {
		return (*int)(ptr)
	})
)

type conversionMethod func(ptr unsafe.Pointer) interface{}

func ConversionMethod(name string, method func(ptr unsafe.Pointer) interface{}) conversionMethod {
	m := conversionMethod(method)
	ConversionTypeMap[name] = m
	return m
}

func Builtin(name string) *BuiltinType {
	bi, ok := Builtins[name]
	if !ok {
		return nil
	}
	return &bi
}