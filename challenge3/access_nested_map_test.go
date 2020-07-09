package main

import (
	"fmt"
	"testing"
)

func TestStringValues(t *testing.T) {
	scenarios := []struct {
		nestedMap map[string]interface{}
		key string
		expected string
		expectedErr error
	}{
		{
			map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{"c": "d"},
				},
			},
			"a.b.c",
			"d",
			nil,
		},
		{
			map[string]interface{}{
				"a": 1,
			},
			"a",
			"",
			ValueNotString,
		},
		{
			map[string]interface{}{
				"a": "z",
			},
			"doesn't exist",
			"",
			KeyNotFound,
		},
	}

	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("scenario %d", i), func(t *testing.T) {
			value, err := accessNestedStruct(scenario.nestedMap, scenario.key)
			if value != scenario.expected {
				t.Errorf("The value returned is incorrect, expected 'd' got '%v'", value)
			}

			if err != scenario.expectedErr {
				t.Errorf("Expected error %v, got %v instead.", scenario.expectedErr, err)
			}
		})
	}
}