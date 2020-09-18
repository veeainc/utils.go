package types

import (
	"fmt"
	"reflect"
)

// Check if its a slice.
func isSlice(v interface{}) bool {
	if v == nil {
		return false
	}
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

// Check if its a map.
func isMap(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

// Check if an interface{} is a map and contains the provided key.
func hasKey(obj interface{}, key string) bool {
	if m, ok := obj.(map[string]interface{}); ok {
		if _, found := m[key]; found {
			return true
		}
	}
	return false
}

// Return an interface{} key value if its a map
func getKey(obj interface{}, key string) interface{} {
	if m, ok := obj.(map[string]interface{}); ok {
		if k, found := m[key]; found {
			return k
		}
	}
	return nil
}

// Convert any types to string
func toString(any interface{}) string {
	return fmt.Sprintf("%v", any)
}
