package ptypes

const (
	// The UINT (unsigned int) type name.
	UINT = "uint"
	// The INT (integer) type name.
	INT              = "int"
	// NoWidthSpecifier is the specifier used to
	// specify the int or uint has no width specifier (such as 64)
	NoWidthSpecifier = 1
)

// The Int function returns an int pointer to originally stored value,
// if it exists.
func (box IntBox) Int() *int {
	// The conversion, or the reason that the conversion failed.
	conversion, err := box.IntOfWidth(NoWidthSpecifier)
	// If the conversion failed, return back no value.
	if err != nil {
		return nil
	}
	// The conversion succeeded, return back the value with
	// the int* assertion.
	return conversion.(*int)
}

// The IntOfWidth function handles the conversion of pointer to the original
// integer of the original width.
func (box *IntBox) IntOfWidth(width int) (interface{}, error) {
	// Get the parent type.
	parent := Box(*box)
	// Get the builtin name based off the width specifier.
	builtinName := INT + determine(
		width == NoWidthSpecifier,
		"",
		string(rune(width)),
	).(string)
	// Get the builtin.
	p := Builtin(builtinName)
	// If there was no builtin, return describing the error.
	if p == nil {
		return nil, makeError("there is no builtin for the given name " + builtinName)
	}
	// Check if the parent can assign to the builtin specified,
	// and that the box has no erasure.
	if _, err := parent.CanAssign(*p); !box.HasErasure && err != nil {
		return nil, err
	}
	// Do the conversion.
	conversion := p.Convert(parent)
	return conversion, nil
}

// The Uint function handles conversion the value in the pointer back
// into the original `uint` value.
func (box *IntBox) Uint() *uint {
	// The conversion, or the reason that the conversion failed.
	conversion, err := box.UintOfWidth(NoWidthSpecifier)
	// If the conversion failed, return back no value.
	if err != nil {
		return nil
	}
	// The conversion succeeded, return back the value with
	// the uint* assertion.
	return conversion.(*uint)
}

// The UintOfWidth function converts the data stored in the pointer
// into a uint of it's original width.
func (box *IntBox) UintOfWidth(width int) (interface{}, error) {
	// Get the parent-typed box.
	parent := Box(*box)
	// Determine the name of the builtin to retrieve.
	builtinName := UINT + determine(
		width == NoWidthSpecifier,
		"", string(rune(width)),
	).(string)
	// Get the builtin.
	p := Builtin(builtinName)
	// If the builtin does not exist, return an error
	// describing that.
	if p == nil {
		return nil, makeError("there is no builtin for the given name " + builtinName)
	}
	// Check if the parent can assign to the requested builtin,
	// and that the box does not have erasure.
	// If there is an error, return it.
	if _, err := parent.CanAssign(*p); !box.HasErasure && err != nil {
		return nil, err
	}
	// Do the actual conversion.
	conversion := p.Convert(parent)
	return conversion, nil
}

// determine is a shorthand function utilised for handling ternary
// type assignments.
func determine(condition bool, option interface{}, other interface{}) interface{} {
	// The condition.
	if condition {
		// Return the first option.
		return option
	} else {
		// Otherwise, return the other option.
		return other
	}
}
