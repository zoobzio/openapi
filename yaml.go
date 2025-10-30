package openapi

import "gopkg.in/yaml.v3"

// ToYAML marshals the OpenAPI spec to YAML bytes.
func (o *OpenAPI) ToYAML() ([]byte, error) {
	return yaml.Marshal(o)
}

// FromYAML unmarshals YAML bytes into an OpenAPI spec.
func FromYAML(data []byte) (*OpenAPI, error) {
	var spec OpenAPI
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, err
	}
	return &spec, nil
}
