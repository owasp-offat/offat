// Package graph models OFFAT's finding graph and runs the rule-based chain
// planner. Findings are nodes; edges are induced by the Primitives a finding
// exposes and the primitives the next check needs.
package graph

import "github.com/owasp-offat/offat/pkg/offat"

// RulePlanner is M1's deterministic planner. It walks a small set of known
// chain templates over the primitive vocabulary.
type RulePlanner struct{}

// Propose returns chains derived from the deterministic templates. M1 returns
// nil until templates are implemented.
func (RulePlanner) Propose(_ []offat.Finding) []offat.Chain { return nil }

var _ offat.Planner = RulePlanner{}
