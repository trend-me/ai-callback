package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func JsonEqual(a, b string) bool {
	var j1, j2 map[string]interface{}

	if err := json.Unmarshal([]byte(a), &j1); err != nil {
		fmt.Printf("ERROR unmarshalling a: %v\n", err)
		return false
	}
	if err := json.Unmarshal([]byte(b), &j2); err != nil {
		fmt.Printf("ERROR unmarshalling b: %v\n", err)
		return false
	}

	return reflect.DeepEqual(j1, j2)
}
