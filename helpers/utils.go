package helpers

import (
	"encoding/json"
	"fmt"
)

// ToJSON converts any JSON-serializable value to JSON bytes.
// Returns the JSON bytes and any error that occurred during marshaling.
func ToJSON(v interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
	}
	return jsonBytes, nil
}

// ToJSONString converts any JSON-serializable value to a JSON string.
// Returns the JSON string and any error that occurred during marshaling.
func ToJSONString(v interface{}) (int, string, error) {
	jsonBytes, err := ToJSON(v)
	if err != nil {
		return 0, "", err
	}
	return len(jsonBytes), string(jsonBytes), nil
}

// Must variants that panic instead of returning errors
// Useful for cases where you know the data is valid

// MustToJSON converts any JSON-serializable value to JSON bytes.
// Panics if marshaling fails.
func MustToJSON(v interface{}) []byte {
	jsonBytes, err := ToJSON(v)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// MustToJSONString converts any JSON-serializable value to a JSON string.
// Panics if marshaling fails.
func MustToJSONString(v interface{}) (int, string) {
	contentLength, jsonStr, err := ToJSONString(v)
	if err != nil {
		panic(err)
	}
	return contentLength, jsonStr
}
