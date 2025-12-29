// Package openapi provides a comprehensive Go implementation of the OpenAPI 3.1 specification.
package openapi

import "encoding/json"

// ToJSON marshals the OpenAPI spec to JSON bytes.
func (o *OpenAPI) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}

// FromJSON unmarshals JSON bytes into an OpenAPI spec.
func FromJSON(data []byte) (*OpenAPI, error) {
	var spec OpenAPI
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, err
	}
	return &spec, nil
}
