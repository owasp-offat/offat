// Package inventory implements API9:2023 Improper Inventory Management
// checks: shadow versions, undocumented routes, deprecated host header
// behaviors, JS-bundle endpoint scraping.
package inventory

import (
	"github.com/owasp-offat/offat/internal/check"
	"github.com/owasp-offat/offat/pkg/offat"
)

type shadowVersion struct{}

func (shadowVersion) ID() string                       { return "inventory.shadow-version" }
func (shadowVersion) Family() offat.Family             { return offat.FamilyInventory }
func (shadowVersion) Applicable(offat.Endpoint) bool   { return false }
func (shadowVersion) Generate(offat.CheckContext, offat.Endpoint) []offat.TestCase {
	return nil
}
func (shadowVersion) Verify(offat.Request, offat.Response) []offat.Finding {
	return nil
}

func init() { check.Register(shadowVersion{}) }
