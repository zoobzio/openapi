# OpenAPI

[![CI Status](https://github.com/zoobz-io/openapi/workflows/CI/badge.svg)](https://github.com/zoobz-io/openapi/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/zoobz-io/openapi/graph/badge.svg?branch=main)](https://codecov.io/gh/zoobz-io/openapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/zoobz-io/openapi)](https://goreportcard.com/report/github.com/zoobz-io/openapi)
[![CodeQL](https://github.com/zoobz-io/openapi/workflows/CodeQL/badge.svg)](https://github.com/zoobz-io/openapi/security/code-scanning)
[![Go Reference](https://pkg.go.dev/badge/github.com/zoobz-io/openapi.svg)](https://pkg.go.dev/github.com/zoobz-io/openapi)
[![License](https://img.shields.io/github/license/zoobz-io/openapi)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/zoobz-io/openapi)](go.mod)
[![Release](https://img.shields.io/github/v/release/zoobz-io/openapi)](https://github.com/zoobz-io/openapi/releases)

OpenAPI 3.1 as native Go types. Build, read, and write API specifications with full type safety.

## OpenAPI in Go

The OpenAPI specification becomes Go structs—each type mirrors the spec exactly:

```go
spec := &openapi.OpenAPI{
    OpenAPI: "3.1.0",
    Info: openapi.Info{
        Title:   "Users API",
        Version: "1.0.0",
    },
    Paths: map[string]openapi.PathItem{
        "/users/{id}": {
            Get: &openapi.Operation{
                OperationID: "getUser",
                Parameters: []openapi.Parameter{{
                    Name:     "id",
                    In:       "path",
                    Required: true,
                    Schema:   &openapi.Schema{Type: openapi.NewSchemaType("string")},
                }},
                Responses: map[string]openapi.Response{
                    "200": {Description: "User found"},
                },
            },
        },
    },
}
```

No wrapper functions. No builder patterns. Just the specification as data.

## Install

```bash
go get github.com/zoobz-io/openapi
```

Requires Go 1.24 or higher.

## Quick Start

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/zoobz-io/openapi"
    "gopkg.in/yaml.v3"
)

func main() {
    // Build a specification
    spec := &openapi.OpenAPI{
        OpenAPI: "3.1.0",
        Info: openapi.Info{
            Title:       "Pet Store",
            Version:     "1.0.0",
            Description: "A sample pet store API",
        },
        Paths: map[string]openapi.PathItem{
            "/pets": {
                Get: &openapi.Operation{
                    Summary:     "List all pets",
                    OperationID: "listPets",
                    Responses: map[string]openapi.Response{
                        "200": {
                            Description: "A list of pets",
                            Content: map[string]openapi.MediaType{
                                "application/json": {
                                    Schema: &openapi.Schema{
                                        Type:  openapi.NewSchemaType("array"),
                                        Items: &openapi.Schema{Ref: "#/components/schemas/Pet"},
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
                "Pet": {
                    Type: openapi.NewSchemaType("object"),
                    Properties: map[string]*openapi.Schema{
                        "id":   {Type: openapi.NewSchemaType("integer")},
                        "name": {Type: openapi.NewSchemaType("string")},
                    },
                    Required: []string{"id", "name"},
                },
            },
        },
    }

    // Output as JSON
    json.NewEncoder(os.Stdout).Encode(spec)

    // Or YAML
    yaml.NewEncoder(os.Stdout).Encode(spec)

    // Read existing specs
    data, _ := os.ReadFile("api.yaml")
    var existing openapi.OpenAPI
    yaml.Unmarshal(data, &existing)
    fmt.Println(existing.Info.Title)
}
```

## Why OpenAPI?

- **Direct mapping**: Types match the OpenAPI 3.1 specification exactly—no translation layer
- **Full coverage**: Every spec construct supported, from basic paths to discriminators and callbacks
- **Dual format**: JSON and YAML serialization work identically via struct tags
- **Zero magic**: Standard Go marshalling, no reflection tricks or code generation
- **Minimal footprint**: Single dependency (yaml.v3), pure data structures

## Documentation

- [Overview](docs/1.overview.md) — Design philosophy and scope
- [Quick Start](docs/2.quickstart.md) — Common usage patterns
- [Types](docs/3.types.md) — Type hierarchy and relationships

## Contributing

Contributions welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

MIT
