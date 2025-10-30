package openapi

import (
	"testing"
)

func TestToYAML(t *testing.T) {
	spec := &OpenAPI{
		OpenAPI: "3.0.3",
		Info: Info{
			Title:   "Test API",
			Version: "1.0.0",
		},
		Paths: map[string]PathItem{
			"/users": {
				Get: &Operation{
					Summary: "List users",
					Responses: map[string]Response{
						"200": {Description: "Success"},
					},
				},
			},
		},
	}

	data, err := spec.ToYAML()
	if err != nil {
		t.Fatalf("ToYAML failed: %v", err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty YAML output")
	}
}

func TestFromYAML(t *testing.T) {
	yamlData := []byte(`openapi: 3.0.3
info:
  title: Test API
  version: 1.0.0
paths:
  /users:
    get:
      summary: List users
      responses:
        "200":
          description: Success
`)

	spec, err := FromYAML(yamlData)
	if err != nil {
		t.Fatalf("FromYAML failed: %v", err)
	}
	if spec.OpenAPI != "3.0.3" {
		t.Errorf("expected openapi '3.0.3', got %q", spec.OpenAPI)
	}
	if spec.Info.Title != "Test API" {
		t.Errorf("expected title 'Test API', got %q", spec.Info.Title)
	}
	if len(spec.Paths) != 1 {
		t.Errorf("expected 1 path, got %d", len(spec.Paths))
	}
}

func TestFromYAML_Invalid(t *testing.T) {
	invalidYAML := []byte(`invalid: [yaml`)

	_, err := FromYAML(invalidYAML)
	if err == nil {
		t.Error("expected error for invalid YAML")
	}
}

func TestYAML_RoundTrip(t *testing.T) {
	original := &OpenAPI{
		OpenAPI: "3.0.3",
		Info: Info{
			Title:       "Test API",
			Version:     "1.0.0",
			Description: "API for testing",
		},
		Servers: []Server{
			{URL: "https://api.example.com"},
		},
		Paths: map[string]PathItem{
			"/users/{id}": {
				Get: &Operation{
					OperationID: "getUserById",
					Parameters: []Parameter{
						{Name: "id", In: "path", Required: true, Schema: &Schema{Type: "string"}},
					},
					Responses: map[string]Response{
						"200": {Description: "Success"},
						"404": {Description: "Not found"},
					},
				},
			},
		},
	}

	yamlData, err := original.ToYAML()
	if err != nil {
		t.Fatalf("ToYAML failed: %v", err)
	}

	restored, err := FromYAML(yamlData)
	if err != nil {
		t.Fatalf("FromYAML failed: %v", err)
	}

	if restored.OpenAPI != original.OpenAPI {
		t.Errorf("openapi mismatch: expected %q, got %q", original.OpenAPI, restored.OpenAPI)
	}
	if restored.Info.Title != original.Info.Title {
		t.Errorf("title mismatch: expected %q, got %q", original.Info.Title, restored.Info.Title)
	}
	if len(restored.Servers) != len(original.Servers) {
		t.Errorf("servers length mismatch: expected %d, got %d", len(original.Servers), len(restored.Servers))
	}
	if len(restored.Paths) != len(original.Paths) {
		t.Errorf("paths length mismatch: expected %d, got %d", len(original.Paths), len(restored.Paths))
	}
}
