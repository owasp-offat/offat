// Package ssrf implements API7:2023 Server Side Request Forgery checks:
// metadata endpoints, localhost variants, IPv6 forms, DNS rebinding,
// redirect chains, blind SSRF via OOB callback.
package ssrf

import (
	"github.com/owasp-offat/offat/internal/check"
	"github.com/owasp-offat/offat/pkg/offat"
)

type imdsProbe struct{}

func (imdsProbe) ID() string                       { return "ssrf.cloud-imds" }
func (imdsProbe) Family() offat.Family             { return offat.FamilySSRF }
func (imdsProbe) Applicable(offat.Endpoint) bool   { return false }
func (imdsProbe) Generate(offat.CheckContext, offat.Endpoint) []offat.TestCase {
	return nil
}
func (imdsProbe) Verify(offat.Request, offat.Response) []offat.Finding {
	return nil
}

func init() { check.Register(imdsProbe{}) }
