package openapi

import (
	"encoding/json"
	"testing"
)

func TestJSON_Marshal(t *testing.T) {
	spec := &OpenAPI{
		OpenAPI: "3.0.3",
		Info: Info{
			Title:       "Test API",
			Version:     "1.0.0",
			Description: "A test API",
		},
		Paths: map[string]PathItem{
			"/users": {
				Get: &Operation{
					Summary:     "List users",
					OperationID: "listUsers",
					Responses: map[string]Response{
						"200": {
							Description: "Success",
						},
					},
				},
			},
		},
	}

	data, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}

	if len(data) == 0 {
		t.Error("marshaled JSON is empty")
	}

	t.Logf("Marshaled JSON:\n%s", string(data))
}

func TestJSON_Unmarshal(t *testing.T) {
	jsonData := `{
  "openapi": "3.0.3",
  "info": {
    "title": "Test API",
    "version": "1.0.0"
  },
  "paths": {
    "/users": {
      "get": {
        "summary": "List users",
        "operationId": "listUsers",
        "responses": {
          "200": {
            "description": "Success"
          }
        }
      }
    }
  }
}`

	var spec OpenAPI
	err := json.Unmarshal([]byte(jsonData), &spec)
	if err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if spec.OpenAPI != "3.0.3" {
		t.Errorf("expected openapi '3.0.3', got %q", spec.OpenAPI)
	}
	if spec.Info.Title != "Test API" {
		t.Errorf("expected title 'Test API', got %q", spec.Info.Title)
	}
}

func TestJSON_RoundTrip(t *testing.T) {
	original := &OpenAPI{
		OpenAPI: "3.0.3",
		Info: Info{
			Title:   "Round Trip Test",
			Version: "1.0.0",
		},
		Components: &Components{
			SecuritySchemes: map[string]*SecurityScheme{
				"bearerAuth": {
					Type:   "http",
					Scheme: "bearer",
				},
			},
		},
	}

	// Marshal to JSON
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	// Unmarshal back to struct
	var result OpenAPI
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	// Verify
	if result.OpenAPI != original.OpenAPI {
		t.Errorf("openapi mismatch: got %q, want %q", result.OpenAPI, original.OpenAPI)
	}
	if result.Info.Title != original.Info.Title {
		t.Errorf("title mismatch: got %q, want %q", result.Info.Title, original.Info.Title)
	}
}
