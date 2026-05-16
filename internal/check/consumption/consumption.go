// Package consumption implements API4:2023 Unrestricted Resource Consumption
// checks: oversized limit params, deeply-nested queries, ReDoS payloads,
// multipart bombs, slow-loris timeouts.
package consumption

import (
	"github.com/owasp-offat/offat/internal/check"
	"github.com/owasp-offat/offat/pkg/offat"
)

type oversizedLimit struct{}

func (oversizedLimit) ID() string                       { return "consumption.oversized-limit" }
func (oversizedLimit) Family() offat.Family             { return offat.FamilyConsumption }
func (oversizedLimit) Applicable(offat.Endpoint) bool   { return false }
func (oversizedLimit) Generate(offat.CheckContext, offat.Endpoint) []offat.TestCase {
	return nil
}
func (oversizedLimit) Verify(offat.Request, offat.Response) []offat.Finding {
	return nil
}

func init() { check.Register(oversizedLimit{}) }
