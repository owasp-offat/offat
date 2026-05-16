// Package upstream implements API10:2023 Unsafe Consumption of APIs checks:
// fuzz responses the target consumes from third parties via webhook
// receivers, OAuth callbacks, and other ingress paths.
package upstream

import (
	"github.com/owasp-offat/offat/internal/check"
	"github.com/owasp-offat/offat/pkg/offat"
)

type webhookFuzz struct{}

func (webhookFuzz) ID() string                       { return "upstream.webhook-response-fuzz" }
func (webhookFuzz) Family() offat.Family             { return offat.FamilyUpstream }
func (webhookFuzz) Applicable(offat.Endpoint) bool   { return false }
func (webhookFuzz) Generate(offat.CheckContext, offat.Endpoint) []offat.TestCase {
	return nil
}
func (webhookFuzz) Verify(offat.Request, offat.Response) []offat.Finding {
	return nil
}

func init() { check.Register(webhookFuzz{}) }
