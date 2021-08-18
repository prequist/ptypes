package boxes

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type Box struct {
	Value        unsafe.Pointer
	OriginalType reflect.Type
	HasErasure	 bool
}

type IntBox Box

func (box *Box) String() (string, error) {
	if _, err := box.check(*String); err != nil {
		return "", err
	}
	conversion := String.Convert(*box)
	return conversion.(string), nil
}

func (box *Box) Interface() (interface{}, error) {
	conversion := Interface.Convert(*box)
	return conversion, nil
}

func (box Box) IntBox() IntBox {
	return IntBox(box)
}


func (box *Box) Type() string {
	return box.OriginalType.String()
}

func (box *Box) check(newType BuiltinType) (bool, error) {
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
