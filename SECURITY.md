# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| Latest  | :white_check_mark: |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security issue, please follow these steps:

1. **DO NOT** create a public GitHub issue
2. Email security details to the maintainers
3. Include:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if available)

## Security Best Practices

When using OpenAPI:

1. **Input Validation**: When unmarshaling untrusted data, validate the resulting structures before use.

2. **Schema Definitions**: Be cautious when programmatically generating schemas from user input, as this could lead to injection attacks.

3. **Reference Resolution**: If implementing `$ref` resolution, ensure proper bounds checking to prevent infinite loops or resource exhaustion.

4. **Example Data**: Example values in schemas may be rendered in documentation. Ensure they don't contain sensitive information.

5. **External URLs**: When processing `url` fields (servers, external docs, etc.), validate them to prevent SSRF attacks.

## Security Features

OpenAPI is designed with security in mind:

- Minimal dependencies (only gopkg.in/yaml.v3)
- No network operations
- No file system operations beyond normal Go imports
- Pure data structures with no executable code
- Thread-safe for concurrent use

## Acknowledgments

We appreciate responsible disclosure of security vulnerabilities.
