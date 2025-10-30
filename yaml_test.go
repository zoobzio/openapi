package openapi

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestYAML_Marshal(t *testing.T) {
	spec := &OpenAPI{
		OpenAPI: "3.0.3",
		Info: Info{
			Title:       "Test API",
			Version:     "1.0.0",
			Description: "A test API",
			License: &License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
		},
		Servers: []Server{
			{
				URL:         "https://api.example.com",
				Description: "Production server",
			},
		},
		Paths: map[string]PathItem{
			"/users/{id}": {
				Get: &Operation{
					Summary:     "Get user by ID",
					OperationID: "getUserById",
					Parameters: []Parameter{
						{
							Name:     "id",
							In:       "path",
							Required: true,
							Schema:   &Schema{Type: "string"},
						},
					},
					Responses: map[string]Response{
						"200": {
							Description: "Success",
							Content: map[string]MediaType{
								"application/json": {
									Schema: &Schema{
										Ref: "#/components/schemas/User",
									},
								},
							},
						},
					},
				},
			},
		},
		Components: &Components{
			Schemas: map[string]*Schema{
				"User": {
					Type: "object",
					Properties: map[string]*Schema{
						"id":   {Type: "string"},
						"name": {Type: "string"},
					},
					Required: []string{"id", "name"},
				},
			},
		},
	}

	data, err := yaml.Marshal(spec)
	if err != nil {
		t.Fatalf("failed to marshal YAML: %v", err)
	}

	if len(data) == 0 {
		t.Error("marshaled YAML is empty")
	}

	t.Logf("Marshaled YAML:\n%s", string(data))
}

func TestYAML_Unmarshal(t *testing.T) {
	yamlData := `
openapi: 3.0.3
info:
  title: Test API
  version: 1.0.0
  description: A test API
  license:
    name: MIT
    url: https://opensource.org/licenses/MIT
servers:
  - url: https://api.example.com
    description: Production server
paths:
  /users/{id}:
    get:
      summary: Get user by ID
      operationId: getUserById
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
      responses:
        "200":
          description: Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/User"
components:
  schemas:
    User:
      type: object
      required:
        - id
        - name
      properties:
        id:
          type: string
        name:
          type: string
`

	var spec OpenAPI
	err := yaml.Unmarshal([]byte(yamlData), &spec)
	if err != nil {
		t.Fatalf("failed to unmarshal YAML: %v", err)
	}

	if spec.OpenAPI != "3.0.3" {
		t.Errorf("expected openapi '3.0.3', got %q", spec.OpenAPI)
	}
	if spec.Info.Title != "Test API" {
		t.Errorf("expected title 'Test API', got %q", spec.Info.Title)
	}
	if spec.Info.License == nil {
		t.Fatal("license should not be nil")
	}
	if spec.Info.License.Name != "MIT" {
		t.Errorf("expected license name 'MIT', got %q", spec.Info.License.Name)
	}
	if len(spec.Servers) != 1 {
		t.Fatalf("expected 1 server, got %d", len(spec.Servers))
	}
	if spec.Servers[0].URL != "https://api.example.com" {
		t.Errorf("expected server url 'https://api.example.com', got %q", spec.Servers[0].URL)
	}
	if len(spec.Paths) != 1 {
		t.Fatalf("expected 1 path, got %d", len(spec.Paths))
	}
	if spec.Components == nil {
		t.Fatal("components should not be nil")
	}
	if len(spec.Components.Schemas) != 1 {
		t.Fatalf("expected 1 schema, got %d", len(spec.Components.Schemas))
	}
}

func TestYAML_RoundTrip(t *testing.T) {
	original := &OpenAPI{
		OpenAPI: "3.0.3",
		Info: Info{
			Title:   "Round Trip Test",
			Version: "1.0.0",
		},
		Paths: map[string]PathItem{
			"/test": {
				Get: &Operation{
					Summary: "Test operation",
					Responses: map[string]Response{
						"200": {Description: "OK"},
					},
				},
			},
		},
	}

	// Marshal to YAML
	data, err := yaml.Marshal(original)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	// Unmarshal back to struct
	var result OpenAPI
	err = yaml.Unmarshal(data, &result)
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
	if len(result.Paths) != len(original.Paths) {
		t.Errorf("paths count mismatch: got %d, want %d", len(result.Paths), len(original.Paths))
	}
}

func TestYAML_SecuritySchemes(t *testing.T) {
	yamlData := `
openapi: 3.0.3
info:
  title: Security Test
  version: 1.0.0
components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT
    apiKey:
      type: apiKey
      name: X-API-Key
      in: header
    oauth2:
      type: oauth2
      flows:
        authorizationCode:
          authorizationUrl: https://example.com/oauth/authorize
          tokenUrl: https://example.com/oauth/token
          scopes:
            read:users: Read user data
            write:users: Write user data
`

	var spec OpenAPI
	err := yaml.Unmarshal([]byte(yamlData), &spec)
	if err != nil {
		t.Fatalf("failed to unmarshal YAML: %v", err)
	}

	if spec.Components == nil {
		t.Fatal("components should not be nil")
	}
	if len(spec.Components.SecuritySchemes) != 3 {
		t.Fatalf("expected 3 security schemes, got %d", len(spec.Components.SecuritySchemes))
	}

	bearer := spec.Components.SecuritySchemes["bearerAuth"]
	if bearer.Type != "http" {
		t.Errorf("expected bearer type 'http', got %q", bearer.Type)
	}
	if bearer.Scheme != "bearer" {
		t.Errorf("expected bearer scheme 'bearer', got %q", bearer.Scheme)
	}

	oauth := spec.Components.SecuritySchemes["oauth2"]
	if oauth.Type != "oauth2" {
		t.Errorf("expected oauth type 'oauth2', got %q", oauth.Type)
	}
	if oauth.Flows == nil || oauth.Flows.AuthorizationCode == nil {
		t.Fatal("oauth flows should not be nil")
	}
	if len(oauth.Flows.AuthorizationCode.Scopes) != 2 {
		t.Errorf("expected 2 scopes, got %d", len(oauth.Flows.AuthorizationCode.Scopes))
	}
}
