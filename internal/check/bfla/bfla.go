// Package bfla implements API5:2023 Broken Function Level Authorization:
// replay admin-tier endpoints with low-priv tokens, verb-swap, cross-role
// path traversal.
package bfla

import (
	"github.com/owasp-offat/offat/internal/check"
	"github.com/owasp-offat/offat/pkg/offat"
)

type verbSwap struct{}

func (verbSwap) ID() string                       { return "bfla.verb-swap" }
func (verbSwap) Family() offat.Family             { return offat.FamilyBFLA }
func (verbSwap) Applicable(offat.Endpoint) bool   { return false }
func (verbSwap) Generate(offat.CheckContext, offat.Endpoint) []offat.TestCase {
	return nil
}
func (verbSwap) Verify(offat.Request, offat.Response) []offat.Finding {
	return nil
}

func init() { check.Register(verbSwap{}) }
