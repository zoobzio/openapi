package openapi

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

const nullType = "null"

// SchemaType represents the type field in a JSON Schema, which can be either
// a single string (e.g., "string") or an array of strings (e.g., ["string", "null"]).
type SchemaType struct {
	value any // string or []string
}

// NewSchemaType creates a SchemaType pointer from a single type string.
func NewSchemaType(s string) *SchemaType {
	return &SchemaType{value: s}
}

// NewSchemaTypes creates a SchemaType pointer from multiple type strings.
func NewSchemaTypes(s []string) *SchemaType {
	if len(s) == 1 {
		return &SchemaType{value: s[0]}
	}
	return &SchemaType{value: s}
}

// IsEmpty returns true if no type is set.
func (t SchemaType) IsEmpty() bool {
	switch v := t.value.(type) {
	case string:
		return v == ""
	case []string:
		return len(v) == 0
	default:
		return true
	}
}

// IsArray returns true if the type is an array of types.
func (t SchemaType) IsArray() bool {
	_, ok := t.value.([]string)
	return ok
}

// String returns the type as a single string.
// If the type is an array, returns the first element.
// Returns empty string if no type is set.
func (t SchemaType) String() string {
	switch v := t.value.(type) {
	case string:
		return v
	case []string:
		if len(v) > 0 {
			return v[0]
		}
		return ""
	default:
		return ""
	}
}

// Strings returns the type as a slice of strings.
// If the type is a single string, returns a slice containing that string.
func (t SchemaType) Strings() []string {
	switch v := t.value.(type) {
	case string:
		if v == "" {
			return nil
		}
		return []string{v}
	case []string:
		return v
	default:
		return nil
	}
}

// Contains returns true if the given type is present.
func (t SchemaType) Contains(s string) bool {
	for _, v := range t.Strings() {
		if v == s {
			return true
		}
	}
	return false
}

// IsNullable returns true if "null" is one of the types.
func (t SchemaType) IsNullable() bool {
	return t.Contains(nullType)
}

// MarshalJSON implements json.Marshaler.
func (t SchemaType) MarshalJSON() ([]byte, error) {
	if t.IsEmpty() {
		return []byte(nullType), nil
	}
	return json.Marshal(t.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *SchemaType) UnmarshalJSON(data []byte) error {
	if string(data) == nullType {
		t.value = nil
		return nil
	}

	// Try string first
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		t.value = s
		return nil
	}

	// Try array of strings
	var arr []string
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}
	if len(arr) == 1 {
		t.value = arr[0]
	} else {
		t.value = arr
	}
	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (t SchemaType) MarshalYAML() (interface{}, error) {
	if t.IsEmpty() {
		return nil, nil
	}
	return t.value, nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (t *SchemaType) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind == yaml.ScalarNode {
		t.value = value.Value
		return nil
	}

	if value.Kind == yaml.SequenceNode {
		var arr []string
		if err := value.Decode(&arr); err != nil {
			return err
		}
		if len(arr) == 1 {
			t.value = arr[0]
		} else {
			t.value = arr
		}
		return nil
	}

	return nil
}
