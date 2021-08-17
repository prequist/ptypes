package ptr

import (
	boxes2 "ptypes/boxes"
	"reflect"
	"unsafe"
)

func FromInt(i int) boxes2.Box {
	return boxes2.Box{Value: unsafe.Pointer(&i), OriginalType: reflect.TypeOf(i)}
}

func FromString(str string) boxes2.Box {
	return boxes2.Box{Value: unsafe.Pointer(&str), OriginalType: reflect.TypeOf(str)}
}

func FromUint(i uint) boxes2.Box {
	return boxes2.Box{Value: unsafe.Pointer(&i), OriginalType: reflect.TypeOf(i)}
}

func FromInt64(i64 int64) boxes2.Box {
	return boxes2.Box{Value: unsafe.Pointer(&i64), 	OriginalType: reflect.TypeOf(i64)}
}