// Package ingest reads traffic captures (HAR, mitmproxy dumps,
// Postman/Bruno/Insomnia collections), GraphQL introspection responses, and
// gRPC reflection output, and normalizes them into Endpoints.
//
// Spec files are handled by package spec; ingest is for captured traffic.
package ingest

import "github.com/owasp-offat/offat/pkg/offat"

// Ingester reads a capture source and returns the Endpoints it discovered.
type Ingester interface {
	Ingest(source string) ([]offat.Endpoint, error)
}
