package openapi

import (
	"encoding/json"
	"testing"
)

func TestToJSON(t *testing.T) {
	spec := &OpenAPI{
		OpenAPI: "3.1.0",
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

	data, err := spec.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty JSON output")
	}

	// Verify it's valid JSON
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		t.Errorf("output is not valid JSON: %v", err)
	}
}

func TestFromJSON(t *testing.T) {
	jsonData := []byte(`{
		"openapi": "3.1.0",
		"info": {
			"title": "Test API",
			"version": "1.0.0"
		},
		"paths": {
			"/users": {
				"get": {
					"summary": "List users",
					"responses": {
						"200": {
							"description": "Success"
						}
					}
				}
			}
		}
	}`)

	spec, err := FromJSON(jsonData)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}
	if spec.OpenAPI != "3.1.0" {
		t.Errorf("expected openapi '3.1.0', got %q", spec.OpenAPI)
	}
	if spec.Info.Title != "Test API" {
		t.Errorf("expected title 'Test API', got %q", spec.Info.Title)
	}
	if len(spec.Paths) != 1 {
		t.Errorf("expected 1 path, got %d", len(spec.Paths))
	}
}

func TestFromJSON_Invalid(t *testing.T) {
	invalidJSON := []byte(`{"invalid": json}`)

	_, err := FromJSON(invalidJSON)
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestJSON_RoundTrip(t *testing.T) {
	original := &OpenAPI{
		OpenAPI: "3.1.0",
		Info: Info{
			Title:       "Test API",
			Version:     "1.0.0",
			Description: "API for testing",
			Contact: &Contact{
				Name:  "Support",
				Email: "support@example.com",
			},
		},
		Servers: []Server{
			{URL: "https://api.example.com", Description: "Production"},
		},
		Paths: map[string]PathItem{
			"/users/{id}": {
				Get: &Operation{
					OperationID: "getUserById",
					Tags:        []string{"users"},
					Parameters: []Parameter{
						{Name: "id", In: "path", Required: true, Schema: &Schema{Type: NewSchemaType("string")}},
					},
					Responses: map[string]Response{
						"200": {Description: "Success"},
						"404": {Description: "Not found"},
					},
				},
			},
		},
		Components: &Components{
			Schemas: map[string]*Schema{
				"User": {
					Type: NewSchemaType("object"),
					Properties: map[string]*Schema{
						"id":   {Type: NewSchemaType("string")},
						"name": {Type: NewSchemaType("string")},
					},
					Required: []string{"id", "name"},
				},
			},
		},
	}

	jsonData, err := original.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	restored, err := FromJSON(jsonData)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}

	if restored.OpenAPI != original.OpenAPI {
		t.Errorf("openapi mismatch: expected %q, got %q", original.OpenAPI, restored.OpenAPI)
	}
	if restored.Info.Title != original.Info.Title {
		t.Errorf("title mismatch: expected %q, got %q", original.Info.Title, restored.Info.Title)
	}
	if restored.Info.Contact.Email != original.Info.Contact.Email {
		t.Errorf("contact email mismatch: expected %q, got %q", original.Info.Contact.Email, restored.Info.Contact.Email)
	}
	if len(restored.Servers) != len(original.Servers) {
		t.Errorf("servers length mismatch: expected %d, got %d", len(original.Servers), len(restored.Servers))
	}
	if len(restored.Paths) != len(original.Paths) {
		t.Errorf("paths length mismatch: expected %d, got %d", len(original.Paths), len(restored.Paths))
	}
	if len(restored.Components.Schemas) != len(original.Components.Schemas) {
		t.Errorf("schemas length mismatch: expected %d, got %d", len(original.Components.Schemas), len(restored.Components.Schemas))
	}
}
