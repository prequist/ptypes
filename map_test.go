package ptypes

import "testing"

// The Map test runs the tests
// for an old proof of concept,
// the original idea of the use case of boxing.
//
// The usage showcased can be eventually implemented as a shorthand
// for easier mapping.
//
// Similar to an array, callback mapping, working on top of
// the type boxing as well as assertions.
//
// Playground: https://play.golang.org/p/LtrjgayWAHm


// The mapper function.
type mapper func(i interface{}) interface{}

// The mapFunction used to execute the mapper.
func mapFunction(arr []interface{}, m mapper) []interface{} {
	// Create an array.
	var newArr []interface{}
	// Iterate over the array.
	for _, v := range arr {
		// Append the mapped values.
		newArr = append(newArr, m(v))
	}
	// Return the new array.
	return newArr
}

// TestMap runs the test for mapping.
func TestMap(t *testing.T) {
	// The mapping function.
	mapFunc := func(i interface{}) interface{} {
		// Type assertion.
		value := i.(Box)
		// Conversion.
		change := *value.IntBox().Int() + 2
		// Return the change, back as a box.
		return FromInt(change)
	}
	// Create an int array.
	intArr := []int{1, 2, 3, 4}
	// Conversion into an interface[]
	var arr []interface{}
	for _, v := range intArr {
		arr = append(arr, FromInt(v))
	}
	// Apply the mappings.
	mapped := mapFunction(arr, mapFunc)
	// Iterate through the output.
	for _, v := range mapped {
		println(*v.(Box).IntBox().Int())
	}
}
