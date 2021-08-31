package test

import (
	"fmt"
	"reflect"
)

// ShouldHaveElements asserts if a slice has the expected elements,
// not considering their position like ShouldResemble does.
func ShouldHaveElements(actual interface{}, expected ...interface{}) string {
	v1 := reflect.ValueOf(actual)

	if !v1.IsValid() {
		return fmt.Sprintf("cannot assert type of %v", actual)
	}

	if v1.Kind() != reflect.Slice {
		return fmt.Sprintf("%v is not a slice", actual)
	}

	v2 := reflect.ValueOf(expected[0])

	if !v2.IsValid() {
		return fmt.Sprintf("cannot assert type of %v", expected[0])
	}

	if v2.Kind() != reflect.Slice {
		return fmt.Sprintf("%v is not a slice", expected[0])
	}

	if v1.IsNil() && v2.IsNil() {
		return ""
	}

	if v1.IsNil() && !v2.IsNil() {
		return fmt.Sprintf("cannot compare %v to nil", v2)
	}

	if v2.IsNil() && !v1.IsNil() {
		return fmt.Sprintf("cannot compare %v to nil", v1)
	}

	if v1.Len() != v2.Len() {
		return fmt.Sprintf("%v and %v have different lengths", v1, v2)
	}

	for i := 0; i < v1.Len(); i++ {
		match := false

		for j := 0; j < v2.Len(); j++ {
			if reflect.DeepEqual(v1.Index(i).Interface(), v2.Index(j).Interface()) {
				v2 = reflect.AppendSlice(v2.Slice(0, j), v2.Slice(j+1, v2.Len()))
				match = true

				break
			}
		}

		if !match {
			return fmt.Sprintf("%v is not expected", v1.Index(i))
		}
	}

	return ""
}
