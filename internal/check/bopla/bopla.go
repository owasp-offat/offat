// Package bopla implements API3:2023 Broken Object Property Level
// Authorization (aka mass assignment) checks.
package bopla

import (
	"github.com/owasp-offat/offat/internal/check"
	"github.com/owasp-offat/offat/pkg/offat"
)

type privFieldInjection struct{}

func (privFieldInjection) ID() string             { return "bopla.privileged-field-injection" }
func (privFieldInjection) Family() offat.Family   { return offat.FamilyBOPLA }
func (privFieldInjection) Applicable(offat.Endpoint) bool { return false }
func (privFieldInjection) Generate(offat.CheckContext, offat.Endpoint) []offat.TestCase {
	return nil
}
func (privFieldInjection) Verify(offat.Request, offat.Response) []offat.Finding {
	return nil
}

func init() { check.Register(privFieldInjection{}) }
