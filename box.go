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
	// The pointer of the value.
	Value        unsafe.Pointer
	// The original type of the value.
	OriginalType reflect.Type
	// If the type has erasure.
	HasErasure	 bool
}

// IntBox is the type alias for a box that has integer capabilities.
type IntBox Box

// The String function tries to convert a value into a string.
func (box *Box) String() (string, error) {
	// Check if the box has no erasure, and the String type can
	// be assigned.
	if _, err := box.CanAssign(*String); !box.HasErasure && err != nil {
		// If it can't be assigned, return an empty string and the error.
		return "", err
	}
	// Convert the string and return the result.
	conversion := String.Convert(*box)
	return conversion.(string), nil
}

// The Interface function tries to convert a value into an interface.
func (box *Box) Interface() interface{} {
	// Call the conversion function.
	conversion := Interface.Convert(*box)
	// The function does not have an error, so this should probably return,
	// with the possibility of `conversion` being nil.
	return conversion
}

// The IntBox function turns a box into a integer capable box.
func (box Box) IntBox() IntBox {
	// Convert the box into an IntBox.
	return IntBox(box)
}

// The Type function gets the name of the original type.
func (box *Box) Type() string {
	// Get the name of the original type.
	return box.OriginalType.String()
}

// CanAssign checks if the box can assign a builtin type
// to it's current stored value.
func (box *Box) CanAssign(newType BuiltinType) (bool, error) {
	// Get the new type's kind and the original type.
	kind := newType.Kind()
	og := box.OriginalType
	// If the kind is predefined, we can compare by string.
	if isPredefinedKind(kind) {
		// If they are not the same, they are not assignable.
		// We can throw an error and exit the flow early.
		if kind.String() != og.Name() {
			return false, makeError(
				"the type %s is not assignable to %s!",
				og.Name(),
				kind.String(),
			)
		}
	}
	// If the types are convertible to each other,
	// there is no error and we can return true.
	if og.ConvertibleTo(newType.Type) {
		return true, nil
	}
	// If the types are not convertible to each other,
	// we create an error and send back the names of the types,
	// also stopping the flow.
	return false, makeError(
		"the type %s is not convertible to %s!",
		og.Name(),
		newType.Name(),
	)
}

// makeError makes an error.
func makeError(format string, args ...interface{}) error {
	// Use c-like function `sprintf` to format the string.
	message := fmt.Sprintf(format, args...)
	// Create the new error.
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
