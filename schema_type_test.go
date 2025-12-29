package openapi

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestSchemaType_NewSchemaType(t *testing.T) {
	st := NewSchemaType("string")
	if st.String() != "string" {
		t.Errorf("expected 'string', got %q", st.String())
	}
	if st.IsArray() {
		t.Error("expected IsArray to be false")
	}
}

func TestSchemaType_NewSchemaTypes_Single(t *testing.T) {
	st := NewSchemaTypes([]string{"string"})
	if st.String() != "string" {
		t.Errorf("expected 'string', got %q", st.String())
	}
	if st.IsArray() {
		t.Error("single element should not be array")
	}
}

func TestSchemaType_NewSchemaTypes_Multiple(t *testing.T) {
	st := NewSchemaTypes([]string{"string", "null"})
	if !st.IsArray() {
		t.Error("expected IsArray to be true")
	}
	if st.String() != "string" {
		t.Errorf("expected first element 'string', got %q", st.String())
	}
	types := st.Strings()
	if len(types) != 2 {
		t.Errorf("expected 2 types, got %d", len(types))
	}
}

func TestSchemaType_IsEmpty(t *testing.T) {
	var empty SchemaType
	if !empty.IsEmpty() {
		t.Error("zero value should be empty")
	}

	st := NewSchemaType("")
	if !st.IsEmpty() {
		t.Error("empty string should be empty")
	}

	st = NewSchemaType("string")
	if st.IsEmpty() {
		t.Error("non-empty string should not be empty")
	}
}

func TestSchemaType_Contains(t *testing.T) {
	st := NewSchemaTypes([]string{"string", "null"})

	if !st.Contains("string") {
		t.Error("should contain 'string'")
	}
	if !st.Contains("null") {
		t.Error("should contain 'null'")
	}
	if st.Contains("integer") {
		t.Error("should not contain 'integer'")
	}
}

func TestSchemaType_IsNullable(t *testing.T) {
	st := NewSchemaType("string")
	if st.IsNullable() {
		t.Error("string type should not be nullable")
	}

	st = NewSchemaTypes([]string{"string", "null"})
	if !st.IsNullable() {
		t.Error("type with 'null' should be nullable")
	}
}

func TestSchemaType_Strings_Single(t *testing.T) {
	st := NewSchemaType("integer")
	types := st.Strings()
	if len(types) != 1 || types[0] != "integer" {
		t.Errorf("expected ['integer'], got %v", types)
	}
}

func TestSchemaType_Strings_Empty(t *testing.T) {
	var empty SchemaType
	types := empty.Strings()
	if types != nil {
		t.Errorf("expected nil, got %v", types)
	}
}

func TestSchemaType_MarshalJSON_String(t *testing.T) {
	st := NewSchemaType("string")
	data, err := json.Marshal(st)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	if string(data) != `"string"` {
		t.Errorf("expected '\"string\"', got %s", data)
	}
}

func TestSchemaType_MarshalJSON_Array(t *testing.T) {
	st := NewSchemaTypes([]string{"string", "null"})
	data, err := json.Marshal(st)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	if string(data) != `["string","null"]` {
		t.Errorf("expected '[\"string\",\"null\"]', got %s", data)
	}
}

func TestSchemaType_MarshalJSON_Empty(t *testing.T) {
	var empty SchemaType
	data, err := json.Marshal(empty)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	if string(data) != "null" {
		t.Errorf("expected 'null', got %s", data)
	}
}

func TestSchemaType_UnmarshalJSON_String(t *testing.T) {
	var st SchemaType
	if err := json.Unmarshal([]byte(`"integer"`), &st); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if st.String() != "integer" {
		t.Errorf("expected 'integer', got %q", st.String())
	}
	if st.IsArray() {
		t.Error("should not be array")
	}
}

func TestSchemaType_UnmarshalJSON_Array(t *testing.T) {
	var st SchemaType
	if err := json.Unmarshal([]byte(`["number", "null"]`), &st); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if !st.IsArray() {
		t.Error("should be array")
	}
	if !st.IsNullable() {
		t.Error("should be nullable")
	}
	types := st.Strings()
	if len(types) != 2 || types[0] != "number" || types[1] != "null" {
		t.Errorf("unexpected types: %v", types)
	}
}

func TestSchemaType_UnmarshalJSON_SingleElementArray(t *testing.T) {
	var st SchemaType
	if err := json.Unmarshal([]byte(`["string"]`), &st); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	// Single element array should be normalized to string
	if st.IsArray() {
		t.Error("single element array should be normalized to string")
	}
	if st.String() != "string" {
		t.Errorf("expected 'string', got %q", st.String())
	}
}

func TestSchemaType_UnmarshalJSON_Null(t *testing.T) {
	var st SchemaType
	if err := json.Unmarshal([]byte(`null`), &st); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if !st.IsEmpty() {
		t.Error("null should result in empty")
	}
}

func TestSchemaType_MarshalYAML_String(t *testing.T) {
	st := NewSchemaType("boolean")
	data, err := yaml.Marshal(st)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	if string(data) != "boolean\n" {
		t.Errorf("expected 'boolean\\n', got %q", data)
	}
}

func TestSchemaType_MarshalYAML_Array(t *testing.T) {
	st := NewSchemaTypes([]string{"object", "null"})
	data, err := yaml.Marshal(st)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	expected := "- object\n- \"null\"\n"
	if string(data) != expected {
		t.Errorf("expected %q, got %q", expected, data)
	}
}

func TestSchemaType_UnmarshalYAML_String(t *testing.T) {
	var st SchemaType
	if err := yaml.Unmarshal([]byte("array"), &st); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if st.String() != "array" {
		t.Errorf("expected 'array', got %q", st.String())
	}
}

func TestSchemaType_UnmarshalYAML_Array(t *testing.T) {
	var st SchemaType
	yamlData := `
- string
- "null"
`
	if err := yaml.Unmarshal([]byte(yamlData), &st); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if !st.IsArray() {
		t.Error("should be array")
	}
	if !st.IsNullable() {
		t.Error("should be nullable")
	}
}

func TestSchemaType_JSON_RoundTrip_String(t *testing.T) {
	original := NewSchemaType("string")
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var restored SchemaType
	if err := json.Unmarshal(data, &restored); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if original.String() != restored.String() {
		t.Errorf("round trip failed: %q != %q", original.String(), restored.String())
	}
}

func TestSchemaType_JSON_RoundTrip_Array(t *testing.T) {
	original := NewSchemaTypes([]string{"integer", "null"})
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var restored SchemaType
	if err := json.Unmarshal(data, &restored); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	origTypes := original.Strings()
	restoredTypes := restored.Strings()
	if len(origTypes) != len(restoredTypes) {
		t.Errorf("round trip failed: length mismatch")
	}
	for i := range origTypes {
		if origTypes[i] != restoredTypes[i] {
			t.Errorf("round trip failed at index %d: %q != %q", i, origTypes[i], restoredTypes[i])
		}
	}
}
