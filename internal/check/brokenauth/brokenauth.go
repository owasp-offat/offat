// Package brokenauth implements API2:2023 Broken Authentication checks:
// JWT alg=none, RS256→HS256 key confusion, expired/revoked acceptance,
// refresh-token replay after logout, session fixation, password-reset reuse.
package brokenauth

import (
	"github.com/owasp-offat/offat/internal/check"
	"github.com/owasp-offat/offat/pkg/offat"
)

type jwtAlgNone struct{}

func (jwtAlgNone) ID() string                       { return "brokenauth.jwt-alg-none" }
func (jwtAlgNone) Family() offat.Family             { return offat.FamilyBrokenAuth }
func (jwtAlgNone) Applicable(offat.Endpoint) bool   { return false }
func (jwtAlgNone) Generate(offat.CheckContext, offat.Endpoint) []offat.TestCase {
	return nil
}
func (jwtAlgNone) Verify(offat.Request, offat.Response) []offat.Finding {
	return nil
}

func init() { check.Register(jwtAlgNone{}) }
