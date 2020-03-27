package interfaces

import "reflect"

// Ensure that an interface{} is passed as a pointer.
// Panic is its not a pointer.
func EnsurePointer(obj interface{}) {
	if reflect.ValueOf(obj).Kind() != reflect.Ptr {
		panic("the argument must be passed as a pointer")
	}
}

// Take an object and enforce a pointer.
// Returning a pointer if its not the case.
func EnforcePointer(obj interface{}) interface{} {
	if reflect.ValueOf(obj).Kind() != reflect.Ptr {
		ptr := reflect.New(reflect.TypeOf(obj))
		ptr.Elem().Set(reflect.ValueOf(obj))
		return ptr.Interface()
	}
	return obj
}

// Return obj type.
func Typ(any interface{}) reflect.Type {
	return reflect.TypeOf(EnforcePointer(any)).Elem()
}
