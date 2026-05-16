// Package bizflow implements API6:2023 Unrestricted Access to Sensitive
// Business Flows: signup loops, coupon stacking, OTP brute-force, rate-limit
// bypass via header rotation, race conditions on inventory/balance.
package bizflow

import (
	"github.com/owasp-offat/offat/internal/check"
	"github.com/owasp-offat/offat/pkg/offat"
)

type rateLimitHeaderRotation struct{}

func (rateLimitHeaderRotation) ID() string                       { return "bizflow.rate-limit-header-rotation" }
func (rateLimitHeaderRotation) Family() offat.Family             { return offat.FamilyBizFlow }
func (rateLimitHeaderRotation) Applicable(offat.Endpoint) bool   { return false }
func (rateLimitHeaderRotation) Generate(offat.CheckContext, offat.Endpoint) []offat.TestCase {
	return nil
}
func (rateLimitHeaderRotation) Verify(offat.Request, offat.Response) []offat.Finding {
	return nil
}

func init() { check.Register(rateLimitHeaderRotation{}) }
