// Package misconfig implements API8:2023 Security Misconfiguration checks:
// verbose stack traces, CORS wildcards with credentials, missing security
// headers, default credentials, exposed admin/debug endpoints.
package misconfig

import (
	"github.com/owasp-offat/offat/internal/check"
	"github.com/owasp-offat/offat/pkg/offat"
)

type corsWildcardWithCredentials struct{}

func (corsWildcardWithCredentials) ID() string                       { return "misconfig.cors-wildcard-with-credentials" }
func (corsWildcardWithCredentials) Family() offat.Family             { return offat.FamilyMisconfig }
func (corsWildcardWithCredentials) Applicable(offat.Endpoint) bool   { return false }
func (corsWildcardWithCredentials) Generate(offat.CheckContext, offat.Endpoint) []offat.TestCase {
	return nil
}
func (corsWildcardWithCredentials) Verify(offat.Request, offat.Response) []offat.Finding {
	return nil
}

func init() { check.Register(corsWildcardWithCredentials{}) }
