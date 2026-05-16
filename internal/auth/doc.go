// Package auth is OFFAT's auth recipe library: bearer, API key (header/query/
// cookie), basic, OAuth2 client credentials in M1; SAML, mTLS, AWS SigV4, HMAC
// in M1.5+.
//
// A Recipe transforms an outgoing request to authenticate as a given Actor.
package auth

import "github.com/owasp-offat/offat/pkg/offat"

// Recipe applies an actor's credentials to a request.
type Recipe interface {
	Name() string
	Apply(req *offat.Request, actor offat.Actor) error
}
