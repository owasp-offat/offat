// Package all blank-imports every check family so their init() registrations
// fire when the binary starts. Import this from cmd/offat (and from tests
// that need the full registry).
package all

import (
	_ "github.com/owasp-offat/offat/internal/check/bfla"
	_ "github.com/owasp-offat/offat/internal/check/bizflow"
	_ "github.com/owasp-offat/offat/internal/check/bola"
	_ "github.com/owasp-offat/offat/internal/check/bopla"
	_ "github.com/owasp-offat/offat/internal/check/brokenauth"
	_ "github.com/owasp-offat/offat/internal/check/consumption"
	_ "github.com/owasp-offat/offat/internal/check/inventory"
	_ "github.com/owasp-offat/offat/internal/check/misconfig"
	_ "github.com/owasp-offat/offat/internal/check/ssrf"
	_ "github.com/owasp-offat/offat/internal/check/upstream"
)
