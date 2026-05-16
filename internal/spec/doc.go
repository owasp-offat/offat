// Package spec parses OpenAPI 3.x / Swagger 2 documents into the engine's
// normalized Endpoint model. It can also infer a spec from captured traffic
// (HAR) when no OpenAPI is available.
//
// Implementation lands in M1 using github.com/getkin/kin-openapi.
package spec

import "github.com/owasp-offat/offat/pkg/offat"

// Parser is the contract all spec sources implement.
type Parser interface {
	Parse(source string) ([]offat.Endpoint, error)
}
