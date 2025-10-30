package openapi

import (
	"testing"
)

func TestOpenAPI_Structure(t *testing.T) {
	spec := &OpenAPI{
		OpenAPI: "3.0.3",
		Info: Info{
			Title:       "Test API",
			Version:     "1.0.0",
			Description: "A test API",
		},
		Paths: make(map[string]PathItem),
		Components: &Components{
			Schemas: make(map[string]*Schema),
		},
	}

	if spec.OpenAPI != "3.0.3" {
		t.Errorf("expected openapi '3.0.3', got %q", spec.OpenAPI)
	}
	if spec.Info.Title != "Test API" {
		t.Errorf("expected title 'Test API', got %q", spec.Info.Title)
	}
	if spec.Paths == nil {
		t.Error("paths should not be nil")
	}
	if spec.Components == nil {
		t.Error("components should not be nil")
	}
}

func TestInfo(t *testing.T) {
	info := Info{
		Title:       "My API",
		Version:     "2.0.0",
		Description: "API Description",
		Contact: &Contact{
			Name:  "Support",
			URL:   "https://example.com",
			Email: "support@example.com",
		},
	}

	if info.Title != "My API" {
		t.Errorf("expected title 'My API', got %q", info.Title)
	}
	if info.Version != "2.0.0" {
		t.Errorf("expected version '2.0.0', got %q", info.Version)
	}
	if info.Contact == nil {
		t.Fatal("contact should not be nil")
	}
	if info.Contact.Email != "support@example.com" {
		t.Errorf("expected email 'support@example.com', got %q", info.Contact.Email)
	}
}

func TestServer(t *testing.T) {
	server := Server{
		URL:         "https://api.example.com",
		Description: "Production server",
	}

	if server.URL != "https://api.example.com" {
		t.Errorf("expected url 'https://api.example.com', got %q", server.URL)
	}
	if server.Description != "Production server" {
		t.Errorf("expected description 'Production server', got %q", server.Description)
	}
}

func TestPathItem(t *testing.T) {
	pathItem := PathItem{
		Summary:     "User operations",
		Description: "Operations for user management",
		Get: &Operation{
			OperationID: "getUser",
			Summary:     "Get user",
		},
		Post: &Operation{
			OperationID: "createUser",
			Summary:     "Create user",
		},
	}

	if pathItem.Summary != "User operations" {
		t.Errorf("expected summary 'User operations', got %q", pathItem.Summary)
	}
	if pathItem.Get == nil {
		t.Fatal("GET operation should not be nil")
	}
	if pathItem.Get.OperationID != "getUser" {
		t.Errorf("expected operation ID 'getUser', got %q", pathItem.Get.OperationID)
	}
}

func TestOperation(t *testing.T) {
	operation := Operation{
		Tags:        []string{"users"},
		Summary:     "Get user by ID",
		Description: "Returns a single user",
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
			},
		},
	}

	if len(operation.Tags) != 1 || operation.Tags[0] != "users" {
		t.Errorf("expected tags ['users'], got %v", operation.Tags)
	}
	if operation.OperationID != "getUserById" {
		t.Errorf("expected operation ID 'getUserById', got %q", operation.OperationID)
	}
	if len(operation.Parameters) != 1 {
		t.Errorf("expected 1 parameter, got %d", len(operation.Parameters))
	}
	if len(operation.Responses) != 1 {
		t.Errorf("expected 1 response, got %d", len(operation.Responses))
	}
}

func TestParameter(t *testing.T) {
	param := Parameter{
		Name:        "userId",
		In:          "path",
		Description: "User ID",
		Required:    true,
		Schema:      &Schema{Type: "string"},
	}

	if param.Name != "userId" {
		t.Errorf("expected name 'userId', got %q", param.Name)
	}
	if param.In != "path" {
		t.Errorf("expected in 'path', got %q", param.In)
	}
	if !param.Required {
		t.Error("expected required to be true")
	}
	if param.Schema.Type != "string" {
		t.Errorf("expected schema type 'string', got %q", param.Schema.Type)
	}
}

func TestRequestBody(t *testing.T) {
	reqBody := RequestBody{
		Description: "User object",
		Required:    true,
		Content: map[string]MediaType{
			"application/json": {
				Schema: &Schema{Ref: "#/components/schemas/User"},
			},
		},
	}

	if reqBody.Description != "User object" {
		t.Errorf("expected description 'User object', got %q", reqBody.Description)
	}
	if !reqBody.Required {
		t.Error("expected required to be true")
	}
	if len(reqBody.Content) != 1 {
		t.Errorf("expected 1 content type, got %d", len(reqBody.Content))
	}
}

func TestResponse(t *testing.T) {
	response := Response{
		Description: "Successful operation",
		Content: map[string]MediaType{
			"application/json": {
				Schema: &Schema{Type: "object"},
			},
		},
	}

	if response.Description != "Successful operation" {
		t.Errorf("expected description 'Successful operation', got %q", response.Description)
	}
	if len(response.Content) != 1 {
		t.Errorf("expected 1 content type, got %d", len(response.Content))
	}
}

func TestSchema_BasicTypes(t *testing.T) {
	tests := []struct {
		name   string
		schema Schema
	}{
		{"string", Schema{Type: "string"}},
		{"integer", Schema{Type: "integer"}},
		{"number", Schema{Type: "number"}},
		{"boolean", Schema{Type: "boolean"}},
		{"array", Schema{Type: "array", Items: &Schema{Type: "string"}}},
		{"object", Schema{Type: "object", Properties: map[string]*Schema{
			"name": {Type: "string"},
		}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.schema.Type != tt.name {
				t.Errorf("expected type %q, got %q", tt.name, tt.schema.Type)
			}
		})
	}
}

func TestSchema_WithRef(t *testing.T) {
	schema := Schema{
		Ref: "#/components/schemas/User",
	}

	if schema.Ref != "#/components/schemas/User" {
		t.Errorf("expected ref '#/components/schemas/User', got %q", schema.Ref)
	}
}

func TestSchema_WithRequired(t *testing.T) {
	schema := Schema{
		Type: "object",
		Properties: map[string]*Schema{
			"name":  {Type: "string"},
			"email": {Type: "string"},
		},
		Required: []string{"name", "email"},
	}

	if len(schema.Required) != 2 {
		t.Errorf("expected 2 required fields, got %d", len(schema.Required))
	}
}

func TestSchema_WithFormat(t *testing.T) {
	schema := Schema{
		Type:   "string",
		Format: "date-time",
	}

	if schema.Format != "date-time" {
		t.Errorf("expected format 'date-time', got %q", schema.Format)
	}
}

func TestComponents(t *testing.T) {
	components := Components{
		Schemas: map[string]*Schema{
			"User": {
				Type: "object",
				Properties: map[string]*Schema{
					"id":   {Type: "integer"},
					"name": {Type: "string"},
				},
			},
		},
		Responses: map[string]*Response{
			"NotFound": {
				Description: "Resource not found",
			},
		},
	}

	if len(components.Schemas) != 1 {
		t.Errorf("expected 1 schema, got %d", len(components.Schemas))
	}
	if len(components.Responses) != 1 {
		t.Errorf("expected 1 response, got %d", len(components.Responses))
	}
}

func TestTag(t *testing.T) {
	tag := Tag{
		Name:        "users",
		Description: "User management operations",
	}

	if tag.Name != "users" {
		t.Errorf("expected name 'users', got %q", tag.Name)
	}
	if tag.Description != "User management operations" {
		t.Errorf("expected description 'User management operations', got %q", tag.Description)
	}
}

func TestMediaType(t *testing.T) {
	mediaType := MediaType{
		Schema: &Schema{Type: "object"},
		Example: map[string]any{
			"name": "John",
			"age":  30,
		},
	}

	if mediaType.Schema.Type != "object" {
		t.Errorf("expected schema type 'object', got %q", mediaType.Schema.Type)
	}
	if mediaType.Example == nil {
		t.Error("expected example to be set")
	}
}

func TestSecurityScheme_APIKey(t *testing.T) {
	scheme := SecurityScheme{
		Type:        "apiKey",
		Name:        "api_key",
		In:          "header",
		Description: "API key authentication",
	}

	if scheme.Type != "apiKey" {
		t.Errorf("expected type 'apiKey', got %q", scheme.Type)
	}
	if scheme.Name != "api_key" {
		t.Errorf("expected name 'api_key', got %q", scheme.Name)
	}
	if scheme.In != "header" {
		t.Errorf("expected in 'header', got %q", scheme.In)
	}
}

func TestSecurityScheme_OAuth2(t *testing.T) {
	scheme := SecurityScheme{
		Type: "oauth2",
		Flows: &OAuthFlows{
			AuthorizationCode: &OAuthFlow{
				AuthorizationURL: "https://example.com/oauth/authorize",
				TokenURL:         "https://example.com/oauth/token",
				Scopes: map[string]string{
					"read:users":  "Read user data",
					"write:users": "Write user data",
				},
			},
		},
	}

	if scheme.Type != "oauth2" {
		t.Errorf("expected type 'oauth2', got %q", scheme.Type)
	}
	if scheme.Flows == nil {
		t.Fatal("flows should not be nil")
	}
	if scheme.Flows.AuthorizationCode == nil {
		t.Fatal("authorization code flow should not be nil")
	}
	if len(scheme.Flows.AuthorizationCode.Scopes) != 2 {
		t.Errorf("expected 2 scopes, got %d", len(scheme.Flows.AuthorizationCode.Scopes))
	}
}

func TestSecurityScheme_HTTP(t *testing.T) {
	scheme := SecurityScheme{
		Type:         "http",
		Scheme:       "bearer",
		BearerFormat: "JWT",
	}

	if scheme.Type != "http" {
		t.Errorf("expected type 'http', got %q", scheme.Type)
	}
	if scheme.Scheme != "bearer" {
		t.Errorf("expected scheme 'bearer', got %q", scheme.Scheme)
	}
	if scheme.BearerFormat != "JWT" {
		t.Errorf("expected bearer format 'JWT', got %q", scheme.BearerFormat)
	}
}

func TestSecurityRequirement(t *testing.T) {
	req := SecurityRequirement{
		"oauth2": {"read:users", "write:users"},
		"apiKey": {},
	}

	if len(req) != 2 {
		t.Errorf("expected 2 security requirements, got %d", len(req))
	}
	if len(req["oauth2"]) != 2 {
		t.Errorf("expected 2 oauth2 scopes, got %d", len(req["oauth2"]))
	}
}

func TestExternalDocumentation(t *testing.T) {
	docs := ExternalDocumentation{
		Description: "Find more info here",
		URL:         "https://example.com/docs",
	}

	if docs.URL != "https://example.com/docs" {
		t.Errorf("expected url 'https://example.com/docs', got %q", docs.URL)
	}
	if docs.Description != "Find more info here" {
		t.Errorf("expected description 'Find more info here', got %q", docs.Description)
	}
}

func TestSchema_Composition(t *testing.T) {
	schema := Schema{
		AllOf: []*Schema{
			{Ref: "#/components/schemas/Base"},
			{
				Type: "object",
				Properties: map[string]*Schema{
					"extra": {Type: "string"},
				},
			},
		},
	}

	if len(schema.AllOf) != 2 {
		t.Errorf("expected 2 allOf schemas, got %d", len(schema.AllOf))
	}
}

func TestSchema_OneOf(t *testing.T) {
	schema := Schema{
		OneOf: []*Schema{
			{Ref: "#/components/schemas/Cat"},
			{Ref: "#/components/schemas/Dog"},
		},
		Discriminator: &Discriminator{
			PropertyName: "petType",
			Mapping: map[string]string{
				"cat": "#/components/schemas/Cat",
				"dog": "#/components/schemas/Dog",
			},
		},
	}

	if len(schema.OneOf) != 2 {
		t.Errorf("expected 2 oneOf schemas, got %d", len(schema.OneOf))
	}
	if schema.Discriminator == nil {
		t.Fatal("discriminator should not be nil")
	}
	if schema.Discriminator.PropertyName != "petType" {
		t.Errorf("expected property name 'petType', got %q", schema.Discriminator.PropertyName)
	}
}

func TestSchema_Default(t *testing.T) {
	schema := Schema{
		Type:    "integer",
		Default: 10,
	}

	if schema.Default != 10 {
		t.Errorf("expected default 10, got %v", schema.Default)
	}
}

func TestHeader(t *testing.T) {
	header := Header{
		Description: "Request ID",
		Required:    true,
		Schema:      &Schema{Type: "string", Format: "uuid"},
	}

	if header.Description != "Request ID" {
		t.Errorf("expected description 'Request ID', got %q", header.Description)
	}
	if !header.Required {
		t.Error("expected required to be true")
	}
	if header.Schema.Format != "uuid" {
		t.Errorf("expected schema format 'uuid', got %q", header.Schema.Format)
	}
}

func TestExample(t *testing.T) {
	example := Example{
		Summary:     "Success example",
		Description: "A successful response",
		Value: map[string]any{
			"id":   123,
			"name": "Test User",
		},
	}

	if example.Summary != "Success example" {
		t.Errorf("expected summary 'Success example', got %q", example.Summary)
	}
	if example.Value == nil {
		t.Error("expected value to be set")
	}
}

func TestLink(t *testing.T) {
	link := Link{
		OperationID: "getUserById",
		Parameters: map[string]any{
			"userId": "$response.body#/id",
		},
		Description: "Get user by ID from response",
	}

	if link.OperationID != "getUserById" {
		t.Errorf("expected operation ID 'getUserById', got %q", link.OperationID)
	}
	if len(link.Parameters) != 1 {
		t.Errorf("expected 1 parameter, got %d", len(link.Parameters))
	}
}

func TestCallback(t *testing.T) {
	callback := Callback{
		"{$request.body#/callbackUrl}": PathItem{
			Post: &Operation{
				Summary: "Webhook notification",
				RequestBody: &RequestBody{
					Content: map[string]MediaType{
						"application/json": {
							Schema: &Schema{Type: "object"},
						},
					},
				},
				Responses: map[string]Response{
					"200": {Description: "Acknowledged"},
				},
			},
		},
	}

	if len(callback) != 1 {
		t.Errorf("expected 1 callback path, got %d", len(callback))
	}
}

func TestDiscriminator(t *testing.T) {
	disc := Discriminator{
		PropertyName: "type",
		Mapping: map[string]string{
			"user":  "#/components/schemas/User",
			"admin": "#/components/schemas/Admin",
		},
	}

	if disc.PropertyName != "type" {
		t.Errorf("expected property name 'type', got %q", disc.PropertyName)
	}
	if len(disc.Mapping) != 2 {
		t.Errorf("expected 2 mappings, got %d", len(disc.Mapping))
	}
}

func TestXML(t *testing.T) {
	xml := XML{
		Name:      "User",
		Namespace: "http://example.com/schema",
		Prefix:    "ex",
		Wrapped:   true,
	}

	if xml.Name != "User" {
		t.Errorf("expected name 'User', got %q", xml.Name)
	}
	if xml.Namespace != "http://example.com/schema" {
		t.Errorf("expected namespace 'http://example.com/schema', got %q", xml.Namespace)
	}
	if !xml.Wrapped {
		t.Error("expected wrapped to be true")
	}
}

func TestLicense(t *testing.T) {
	license := License{
		Name: "Apache 2.0",
		URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
	}

	if license.Name != "Apache 2.0" {
		t.Errorf("expected name 'Apache 2.0', got %q", license.Name)
	}
	if license.URL != "https://www.apache.org/licenses/LICENSE-2.0.html" {
		t.Errorf("expected url 'https://www.apache.org/licenses/LICENSE-2.0.html', got %q", license.URL)
	}
}

func TestServerVariable(t *testing.T) {
	variable := ServerVariable{
		Enum:        []string{"v1", "v2", "v3"},
		Default:     "v2",
		Description: "API version",
	}

	if variable.Default != "v2" {
		t.Errorf("expected default 'v2', got %q", variable.Default)
	}
	if len(variable.Enum) != 3 {
		t.Errorf("expected 3 enum values, got %d", len(variable.Enum))
	}
}

func TestEncoding(t *testing.T) {
	encoding := Encoding{
		ContentType: "application/json",
		Headers: map[string]*Header{
			"X-Custom": {
				Schema: &Schema{Type: "string"},
			},
		},
	}

	if encoding.ContentType != "application/json" {
		t.Errorf("expected content type 'application/json', got %q", encoding.ContentType)
	}
	if len(encoding.Headers) != 1 {
		t.Errorf("expected 1 header, got %d", len(encoding.Headers))
	}
}

func TestComponents_Extended(t *testing.T) {
	components := Components{
		SecuritySchemes: map[string]*SecurityScheme{
			"bearerAuth": {
				Type:   "http",
				Scheme: "bearer",
			},
		},
		Parameters: map[string]*Parameter{
			"PageParam": {
				Name:   "page",
				In:     "query",
				Schema: &Schema{Type: "integer"},
			},
		},
		RequestBodies: map[string]*RequestBody{
			"UserBody": {
				Required: true,
				Content: map[string]MediaType{
					"application/json": {
						Schema: &Schema{Ref: "#/components/schemas/User"},
					},
				},
			},
		},
		Links: map[string]*Link{
			"GetUserByUserId": {
				OperationID: "getUser",
			},
		},
	}

	if len(components.SecuritySchemes) != 1 {
		t.Errorf("expected 1 security scheme, got %d", len(components.SecuritySchemes))
	}
	if len(components.Parameters) != 1 {
		t.Errorf("expected 1 parameter, got %d", len(components.Parameters))
	}
	if len(components.RequestBodies) != 1 {
		t.Errorf("expected 1 request body, got %d", len(components.RequestBodies))
	}
	if len(components.Links) != 1 {
		t.Errorf("expected 1 link, got %d", len(components.Links))
	}
}
