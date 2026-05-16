// Package httpx is OFFAT's HTTP engine. It wraps net/http with a chain of
// RoundTrippers for auth rotation, request recording, OOB header injection,
// rate-limiting, retries with Retry-After awareness, and scope enforcement.
//
// HTTP/2 is on by default via net/http. HTTP/3 is opt-in via quic-go when the
// target advertises ALPN "h3".
package httpx

import (
	"context"

	"github.com/owasp-offat/offat/pkg/offat"
)

// Client is the engine's view of an HTTP client. Concrete implementation
// builds the RoundTripper stack from configuration.
type Client interface {
	Do(ctx context.Context, req offat.Request) (offat.Response, error)
}
