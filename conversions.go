package ptypes

var (
	// The ConversionTypeMap holds the conversion methods with the
	// specified type names.
	ConversionTypeMap = make(map[string]conversionMethod, 0)

	// The Builtins map holds all the registered BuiltinType s
	Builtins    = make(map[string]BuiltinType, 0)

	// InterfaceMethod is the conversion method for the InterfaceAlias type
	// utilised for boxing and handling conversions between
	// pointers and `interface{}`s.
	InterfaceMethod = ConversionMethod("InterfaceAlias", func(ptr Box) interface{} {
		// Dereference the pointer, cast, and return.
		return *(*interface{})(ptr.Value)
	})

	// StringMethod is the conversion method for string types, and can support
	// type erasure.
	StringMethod = ConversionMethod("string", func(ptr Box) interface{} {
		// If the pointer has erasure, this routine will be followed.
		if ptr.HasErasure {
			// Dereference the pointer, and turn it into an interface{}.
			casted := *(*interface{})(ptr.Value)
			// Assert this value back into a string.
			return casted.(string)
		}
		// The pointer had no erasure, and could
		// be de-referenced back into it's original type, the string.
		return *(*string)(ptr.Value)
	})

	// UintMethod is the conversion method for uint types, and can support
	// type erasure.
	UintMethod = ConversionMethod("uint", func(ptr Box) interface{} {
		// If the pointer has erasure, the following flow
		// is executed.
		if ptr.HasErasure {
			// Converted into an interface.
			casted := *(*interface{})(ptr.Value)
			// Asserted into a uint.
			asserted := casted.(uint)
			// Return back a reference to the assertion.
			return &asserted
		}
		// The pointer had no erasure, and could
		// be de-referenced back into a reference of the pointer value.
		return (*uint)(ptr.Value)
	})

	// IntMethod is the conversion method for int types, and can support
	// type erasure.
	IntMethod = ConversionMethod("int", func(ptr Box) interface{} {
		// If the pointer has erasure, the following flow
		// is executed.
		if ptr.HasErasure {
			// De-referenced and converted into an interface.
			casted := *(*interface{})(ptr.Value)
			// Asserted back into an integer.
			asserted := casted.(int)
			// Return back a reference to the assertion.
			return &asserted
		}
		// The pointer had no erasure, and could
		// be de-referenced back into a reference of the pointer value.
		return (*int)(ptr.Value)
	})
)

// The conversion method type alias.
type conversionMethod func(ptr Box) interface{}

// ConversionMethod creates a conversion method and registers it into the map,
// then returns the value.
func ConversionMethod(name string, method func(ptr Box) interface{}) conversionMethod {
	// Convert the function into the conversion method type.
	m := conversionMethod(method)
	// Register the type into the map.
	ConversionTypeMap[name] = m
	// Return the method.
	return m
}

// Builtin gets a reference to a builtin type that can then
// be used for conversions.
func Builtin(name string) *BuiltinType {
	// Check if the built-in is registered.
	bi, ok := Builtins[name]
	// If it doesn't exist, return nil.
	if !ok {
		return nil
	}
	// Return a reference to the builtin type.
	return &bi
}