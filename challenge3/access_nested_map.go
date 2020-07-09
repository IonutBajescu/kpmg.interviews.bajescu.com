package main

import (
	"errors"
	"strings"
)

var ValueNotString = errors.New("the value to be retrieved is not a string")
var KeyNotFound = errors.New("the key cannot be found in the map")

// I'm doing quite a few allocations here that are not entirely
// necessary, so technically we could further improve this function
// once we add a benchmark to the tests to prove the benefits.
func accessNestedStruct(object map[string]interface{}, key string) (string, error) {
	for {
		index := strings.Index(key, ".")
		if index == -1 {
			valueOfAnyType, ok := object[key]
			if !ok {
				return "", KeyNotFound
			}

			valueAsString, ok := valueOfAnyType.(string)
			if !ok {
				return "", ValueNotString
			}

			return valueAsString, nil
		}

		object = object[key[:index]].(map[string]interface{})
		key = key[index+1:]
	}
}
