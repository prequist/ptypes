package boxes

var (
	ConversionTypeMap = make(map[string]conversionMethod, 0)
	Builtins    = make(map[string]BuiltinType, 0)

	InterfaceMethod = ConversionMethod("InterfaceAlias", func(ptr Box) interface{} {
		return *(*interface{})(ptr.Value)
	})

	StringMethod = ConversionMethod("string", func(ptr Box) interface{} {
		if ptr.HasErasure {
			casted := *(*interface{})(ptr.Value)
			return casted.(string)
		}
		return *(*string)(ptr.Value)
	})
	UintMethod = ConversionMethod("uint", func(ptr Box) interface{} {
		if ptr.HasErasure {
			casted := *(*interface{})(ptr.Value)
			return casted.(*uint)
		}
		return (*uint)(ptr.Value)
	})
	IntMethod = ConversionMethod("int", func(ptr Box) interface{} {
		if ptr.HasErasure {
			casted := *(*interface{})(ptr.Value)
			return casted.(*int)
		}
		return (*int)(ptr.Value)
	})
)

type conversionMethod func(ptr Box) interface{}

func ConversionMethod(name string, method func(ptr Box) interface{}) conversionMethod {
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