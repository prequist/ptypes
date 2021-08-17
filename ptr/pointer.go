package ptr

import (
	"ptypes/boxes"
	"reflect"
	"unsafe"
)

func FromInt(i int) boxes.Box {
	return boxes.Box{Value: unsafe.Pointer(&i), OriginalType: reflect.TypeOf(i)}
}

func FromString(str string) boxes.Box {
	return boxes.Box{Value: unsafe.Pointer(&str), OriginalType: reflect.TypeOf(str)}
}

func FromUint(i uint) boxes.Box {
	return boxes.Box{Value: unsafe.Pointer(&i), OriginalType: reflect.TypeOf(i)}
}

func FromInt64(i64 int64) boxes.Box {
	return boxes.Box{Value: unsafe.Pointer(&i64), 	OriginalType: reflect.TypeOf(i64)}
}