package openapi

// OpenAPI 3.0 specification structs
// Minimal implementation for generating API documentation from handlers

// OpenAPI represents the root document object.
type OpenAPI struct {
	Paths        map[string]PathItem    `json:"paths" yaml:"paths"`
	Components   *Components            `json:"components,omitempty" yaml:"components,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Info         Info                   `json:"info" yaml:"info"`
	OpenAPI      string                 `json:"openapi" yaml:"openapi"`
	Servers      []Server               `json:"servers,omitempty" yaml:"servers,omitempty"`
	Security     []SecurityRequirement  `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []Tag                  `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// Info provides metadata about the API.
type Info struct {
	Title          string   `json:"title" yaml:"title"`
	Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version" yaml:"version"`
}

// Contact information for the API.
type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

// License information for the API.
type License struct {
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

// Server represents a server.
type Server struct {
	Variables   map[string]ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
	URL         string                    `json:"url" yaml:"url"`
	Description string                    `json:"description,omitempty" yaml:"description,omitempty"`
}

// ServerVariable represents a server variable for server URL template substitution.
type ServerVariable struct {
	Default     string   `json:"default" yaml:"default"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
}

// PathItem describes operations available on a single path.
type PathItem struct {
	Ref         string      `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string      `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Get         *Operation  `json:"get,omitempty" yaml:"get,omitempty"`
	Post        *Operation  `json:"post,omitempty" yaml:"post,omitempty"`
	Put         *Operation  `json:"put,omitempty" yaml:"put,omitempty"`
	Delete      *Operation  `json:"delete,omitempty" yaml:"delete,omitempty"`
	Patch       *Operation  `json:"patch,omitempty" yaml:"patch,omitempty"`
	Options     *Operation  `json:"options,omitempty" yaml:"options,omitempty"`
	Head        *Operation  `json:"head,omitempty" yaml:"head,omitempty"`
	Servers     []Server    `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation describes a single API operation on a path.
type Operation struct {
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	RequestBody  *RequestBody           `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    map[string]Response    `json:"responses" yaml:"responses"`
	Callbacks    map[string]Callback    `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Summary      string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	OperationID  string                 `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Tags         []string               `json:"tags,omitempty" yaml:"tags,omitempty"`
	Parameters   []Parameter            `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Security     []SecurityRequirement  `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      []Server               `json:"servers,omitempty" yaml:"servers,omitempty"`
	Deprecated   bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
}

// Parameter describes a single operation parameter.
type Parameter struct {
	Example         any                  `json:"example,omitempty" yaml:"example,omitempty"`
	Schema          *Schema              `json:"schema,omitempty" yaml:"schema,omitempty"`
	Content         map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Examples        map[string]*Example  `json:"examples,omitempty" yaml:"examples,omitempty"`
	Explode         *bool                `json:"explode,omitempty" yaml:"explode,omitempty"`
	Name            string               `json:"name,omitempty" yaml:"name,omitempty"`
	In              string               `json:"in,omitempty" yaml:"in,omitempty"`
	Description     string               `json:"description,omitempty" yaml:"description,omitempty"`
	Ref             string               `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Style           string               `json:"style,omitempty" yaml:"style,omitempty"`
	AllowEmptyValue bool                 `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	AllowReserved   bool                 `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Deprecated      bool                 `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Required        bool                 `json:"required,omitempty" yaml:"required,omitempty"`
}

// RequestBody describes a request body.
type RequestBody struct {
	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Ref         string               `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool                 `json:"required,omitempty" yaml:"required,omitempty"`
}

// Response describes a single response from an API operation.
type Response struct {
	Headers     map[string]*Header   `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Ref         string               `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
}

// MediaType provides schema and examples for the media type.
type MediaType struct {
	Schema   *Schema             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  any                 `json:"example,omitempty" yaml:"example,omitempty"`
	Examples map[string]*Example `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding map[string]Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

// Schema represents a JSON Schema.
type Schema struct {
	AdditionalProperties any                    `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	Example              any                    `json:"example,omitempty" yaml:"example,omitempty"`
	Default              any                    `json:"default,omitempty" yaml:"default,omitempty"`
	ExclusiveMaximum     *bool                  `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Not                  *Schema                `json:"not,omitempty" yaml:"not,omitempty"`
	Items                *Schema                `json:"items,omitempty" yaml:"items,omitempty"`
	Deprecated           *bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Discriminator        *Discriminator         `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	Nullable             *bool                  `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	WriteOnly            *bool                  `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	ReadOnly             *bool                  `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	ExternalDocs         *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	MaxLength            *int                   `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	XML                  *XML                   `json:"xml,omitempty" yaml:"xml,omitempty"`
	UniqueItems          *bool                  `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	Minimum              *float64               `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum              *float64               `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMinimum     *bool                  `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxProperties        *int                   `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	Properties           map[string]*Schema     `json:"properties,omitempty" yaml:"properties,omitempty"`
	MinLength            *int                   `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	MultipleOf           *float64               `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	MinProperties        *int                   `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	MinItems             *int                   `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	MaxItems             *int                   `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	Format               string                 `json:"format,omitempty" yaml:"format,omitempty"`
	Pattern              string                 `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	Type                 string                 `json:"type,omitempty" yaml:"type,omitempty"`
	Title                string                 `json:"title,omitempty" yaml:"title,omitempty"`
	Description          string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Ref                  string                 `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	OneOf                []*Schema              `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	Enum                 []any                  `json:"enum,omitempty" yaml:"enum,omitempty"`
	AnyOf                []*Schema              `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	AllOf                []*Schema              `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	Required             []string               `json:"required,omitempty" yaml:"required,omitempty"`
}

// Components holds reusable objects for the specification.
type Components struct {
	Schemas         map[string]*Schema         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]*Response       `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]*Parameter      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]*Example        `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]*RequestBody    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         map[string]*Header         `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Links           map[string]*Link           `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       map[string]*Callback       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
}

// Tag adds metadata to a single tag.
type Tag struct {
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Name         string                 `json:"name" yaml:"name"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
}

// ExternalDocumentation allows referencing an external resource.
type ExternalDocumentation struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url" yaml:"url"`
}

// SecurityRequirement lists required security schemes to execute an operation.
type SecurityRequirement map[string][]string

// SecurityScheme defines a security scheme for the API.
type SecurityScheme struct {
	Type             string      `json:"type" yaml:"type"` // "apiKey", "http", "oauth2", "openIdConnect"
	Description      string      `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string      `json:"name,omitempty" yaml:"name,omitempty"`                         // for apiKey
	In               string      `json:"in,omitempty" yaml:"in,omitempty"`                             // for apiKey: "query", "header", "cookie"
	Scheme           string      `json:"scheme,omitempty" yaml:"scheme,omitempty"`                     // for http: "bearer", "basic", etc.
	BearerFormat     string      `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`         // for http bearer
	Flows            *OAuthFlows `json:"flows,omitempty" yaml:"flows,omitempty"`                       // for oauth2
	OpenIDConnectURL string      `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"` // for openIdConnect
}

// OAuthFlows allows configuration of the supported OAuth flows.
type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

// OAuthFlow configuration details for a supported OAuth flow.
type OAuthFlow struct {
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshURL       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
}

// Header follows the structure of Parameter but without name and in fields.
type Header struct {
	Example         any                  `json:"example,omitempty" yaml:"example,omitempty"`
	Explode         *bool                `json:"explode,omitempty" yaml:"explode,omitempty"`
	Schema          *Schema              `json:"schema,omitempty" yaml:"schema,omitempty"`
	Examples        map[string]*Example  `json:"examples,omitempty" yaml:"examples,omitempty"`
	Content         map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Ref             string               `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Description     string               `json:"description,omitempty" yaml:"description,omitempty"`
	Style           string               `json:"style,omitempty" yaml:"style,omitempty"`
	Required        bool                 `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      bool                 `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool                 `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	AllowReserved   bool                 `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}

// Example object for example values.
type Example struct {
	Ref           string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary       string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   string `json:"description,omitempty" yaml:"description,omitempty"`
	Value         any    `json:"value,omitempty" yaml:"value,omitempty"`
	ExternalValue string `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
}

// Link represents a possible design-time link for a response.
type Link struct {
	RequestBody  any            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Parameters   map[string]any `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Server       *Server        `json:"server,omitempty" yaml:"server,omitempty"`
	Ref          string         `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	OperationRef string         `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	OperationID  string         `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Description  string         `json:"description,omitempty" yaml:"description,omitempty"`
}

// Callback is a map of possible out-of band callbacks.
type Callback map[string]PathItem

// Discriminator supports polymorphism.
type Discriminator struct {
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
	PropertyName string            `json:"propertyName" yaml:"propertyName"`
}

// XML metadata for array/object schemas.
type XML struct {
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
}

// Encoding property for multipart media types.
type Encoding struct {
	Headers       map[string]*Header `json:"headers,omitempty" yaml:"headers,omitempty"`
	Explode       *bool              `json:"explode,omitempty" yaml:"explode,omitempty"`
	ContentType   string             `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Style         string             `json:"style,omitempty" yaml:"style,omitempty"`
	AllowReserved bool               `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}
