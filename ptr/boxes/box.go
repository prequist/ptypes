package boxes

import (
	"errors"
	"fmt"
	"ptypes/ptr/prim"
	"reflect"
	"strings"
	"unsafe"
)

type Box struct {
	Value        unsafe.Pointer
	OriginalType reflect.Type
}

type IntBox Box
type StringBox Box

func (box *Box) String() (string, error) {
	if _, err := box.check(prim.String); err != nil {
		return "", err
	}
	conversion := prim.String.Convert(box.Value)
	return conversion.(string), nil
}

func (box Box) IntBox() IntBox {
	return IntBox(box)
}


func (box *Box) Type() string {
	return box.OriginalType.String()
}

func (box *Box) check(newType prim.BuiltinType) (bool, error) {
	kind := newType.Kind()
	og := box.OriginalType
	if isPredefinedKind(kind) {
		if kind.String() != og.Name() {
			return false, makeError(
				"the type %s is not assignable to %s!",
				og.Name(),
				kind.String(),
			)
		}
	}
	if og.ConvertibleTo(newType.Type) {
		return true, nil
	}
	return false, makeError(
		"the type %s is not assignable to %s!",
		og.Name(),
		newType.Name(),
	)
}

func makeError(format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	return errors.New(message)
}

func isPredefinedKind(kind reflect.Kind) bool {
	return !strings.Contains(kind.String(), "kind")
}
