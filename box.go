package ptypes

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

// The Box struct boxes a value and supports erasure
// for later conversion back into that object.
type Box struct {
	Value        unsafe.Pointer
	OriginalType reflect.Type
	HasErasure	 bool
}

// IntBox is the type alias for a box that has integer capabilities.
type IntBox Box

// The String function tries to convert a value into a string.
func (box *Box) String() (string, error) {
	if _, err := box.CanAssign(*String); !box.HasErasure && err != nil {
		return "", err
	}
	conversion := String.Convert(*box)
	return conversion.(string), nil
}

// The Interface function tries to convert a value into an interface.
func (box *Box) Interface() (interface{}, error) {
	conversion := Interface.Convert(*box)
	return conversion, nil
}

// The IntBox function turns a box into a integer capable box.
func (box Box) IntBox() IntBox {
	return IntBox(box)
}

// The Type function gets the name of the original type.
func (box *Box) Type() string {
	return box.OriginalType.String()
}

// CanAssign checks if the box can assign a builtin type
// to it's current stored value.
func (box *Box) CanAssign(newType BuiltinType) (bool, error) {
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

// makeError makes an error.
func makeError(format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	return errors.New(message)
}

// isPredefinedKind checks if the kind passed through has already
// been predefined as a reflect.Kind.
func isPredefinedKind(kind reflect.Kind) bool {
	// The trick here is that if the kind was not predefined, it is
	// unknown to go, and will contain "kind" in the string.
	// This outcome is unlikely, but it's good to check.
	return !strings.Contains(kind.String(), "kind")
}
