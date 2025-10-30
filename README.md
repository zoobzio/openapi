# OpenAPI

A comprehensive Go implementation of the OpenAPI 3.0 specification.

## Features

- Complete OpenAPI 3.0 specification support
- Full JSON and YAML serialization/deserialization
- All core types: Info, Servers, Paths, Operations, Parameters, Request/Response bodies
- Advanced features: Security schemes, callbacks, links, discriminators, schema composition
- Reusable components for all object types
- Extensive validation support (min/max, patterns, enums, etc.)

## Installation

```bash
go get github.com/zoobzio/openapi
```

## Usage

### Creating an OpenAPI Specification

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/zoobzio/openapi"
)

func main() {
    spec := &openapi.OpenAPI{
        OpenAPI: "3.0.3",
        Info: openapi.Info{
            Title:       "My API",
            Version:     "1.0.0",
            Description: "A sample API",
        },
        Paths: map[string]openapi.PathItem{
            "/users": {
                Get: &openapi.Operation{
                    Summary:     "List users",
                    OperationID: "listUsers",
                    Responses: map[string]openapi.Response{
                        "200": {
                            Description: "Success",
                            Content: map[string]openapi.MediaType{
                                "application/json": {
                                    Schema: &openapi.Schema{
                                        Type: "array",
                                        Items: &openapi.Schema{
                                            Ref: "#/components/schemas/User",
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
        Components: &openapi.Components{
            Schemas: map[string]*openapi.Schema{
                "User": {
                    Type: "object",
                    Properties: map[string]*openapi.Schema{
                        "id":   {Type: "string"},
                        "name": {Type: "string"},
                    },
                    Required: []string{"id", "name"},
                },
            },
        },
    }

    // Output as JSON
    data, _ := json.MarshalIndent(spec, "", "  ")
    fmt.Println(string(data))
}
```

### Working with YAML

```go
import "gopkg.in/yaml.v3"

// Marshal to YAML
data, err := yaml.Marshal(spec)
if err != nil {
    log.Fatal(err)
}

// Unmarshal from YAML
var spec openapi.OpenAPI
err = yaml.Unmarshal(data, &spec)
if err != nil {
    log.Fatal(err)
}
```

### Security Schemes

```go
spec := &openapi.OpenAPI{
    OpenAPI: "3.0.3",
    Info: openapi.Info{
        Title:   "Secure API",
        Version: "1.0.0",
    },
    Components: &openapi.Components{
        SecuritySchemes: map[string]*openapi.SecurityScheme{
            "bearerAuth": {
                Type:         "http",
                Scheme:       "bearer",
                BearerFormat: "JWT",
            },
            "apiKey": {
                Type: "apiKey",
                Name: "X-API-Key",
                In:   "header",
            },
            "oauth2": {
                Type: "oauth2",
                Flows: &openapi.OAuthFlows{
                    AuthorizationCode: &openapi.OAuthFlow{
                        AuthorizationURL: "https://example.com/oauth/authorize",
                        TokenURL:         "https://example.com/oauth/token",
                        Scopes: map[string]string{
                            "read:users":  "Read user data",
                            "write:users": "Write user data",
                        },
                    },
                },
            },
        },
    },
    Security: []openapi.SecurityRequirement{
        {"bearerAuth": {}},
    },
}
```

### Schema Composition

```go
// Using allOf for inheritance
schema := &openapi.Schema{
    AllOf: []*openapi.Schema{
        {Ref: "#/components/schemas/Base"},
        {
            Type: "object",
            Properties: map[string]*openapi.Schema{
                "extra": {Type: "string"},
            },
        },
    },
}

// Using oneOf for polymorphism
schema := &openapi.Schema{
    OneOf: []*openapi.Schema{
        {Ref: "#/components/schemas/Cat"},
        {Ref: "#/components/schemas/Dog"},
    },
    Discriminator: &openapi.Discriminator{
        PropertyName: "petType",
        Mapping: map[string]string{
            "cat": "#/components/schemas/Cat",
            "dog": "#/components/schemas/Dog",
        },
    },
}
```

### Callbacks (Webhooks)

```go
operation := &openapi.Operation{
    Summary: "Subscribe to events",
    Callbacks: map[string]openapi.Callback{
        "onEvent": {
            "{$request.body#/callbackUrl}": openapi.PathItem{
                Post: &openapi.Operation{
                    Summary: "Event notification",
                    RequestBody: &openapi.RequestBody{
                        Content: map[string]openapi.MediaType{
                            "application/json": {
                                Schema: &openapi.Schema{
                                    Type: "object",
                                    Properties: map[string]*openapi.Schema{
                                        "event": {Type: "string"},
                                    },
                                },
                            },
                        },
                    },
                    Responses: map[string]openapi.Response{
                        "200": {Description: "Acknowledged"},
                    },
                },
            },
        },
    },
}
```

### Response Links

```go
responses := map[string]openapi.Response{
    "200": {
        Description: "User created",
        Links: map[string]*openapi.Link{
            "GetUserByUserId": {
                OperationID: "getUserById",
                Parameters: map[string]any{
                    "userId": "$response.body#/id",
                },
            },
        },
    },
}
```

## Complete Feature List

- **Core Objects**: OpenAPI, Info, Contact, License, Server, ServerVariable
- **Paths & Operations**: PathItem, Operation, Parameter, RequestBody, Response
- **Schemas**: Full JSON Schema support with validation, composition (allOf/oneOf/anyOf), discriminators
- **Media Types**: MediaType, Example, Encoding
- **Security**: SecurityScheme, SecurityRequirement, OAuthFlows
- **Advanced**: Callbacks, Links, ExternalDocumentation, Headers
- **Reusable Components**: All object types can be defined in Components for reuse
- **XML Support**: XML metadata for XML APIs
- **Validation**: Min/max, patterns, enums, formats, required fields, etc.

## License

MIT
