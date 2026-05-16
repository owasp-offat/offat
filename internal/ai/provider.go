// Package ai gates all LLM calls behind a Provider interface. OFFAT always
// runs deterministically with the NoOp provider; concrete providers (Claude
// via the Anthropic SDK, etc.) land in M2.
//
// A feature that "only works with AI" is a bug: the NoOp provider must satisfy
// the same behavior contract every other provider does.
package ai

import (
	"context"

	"github.com/owasp-offat/offat/pkg/offat"
)

// Provider is the boundary OFFAT calls when it wants LLM assistance.
type Provider interface {
	Name() string
	// ClassifyEndpoint returns a free-form intent label (e.g. "money-moving").
	ClassifyEndpoint(ctx context.Context, ep offat.Endpoint) (string, error)
	// SynthesizePayloads returns context-aware payloads for a given family.
	SynthesizePayloads(ctx context.Context, ep offat.Endpoint, family offat.Family) ([]string, error)
	// TriageFindings filters and ranks findings; the returned slice is a subset.
	TriageFindings(ctx context.Context, findings []offat.Finding) ([]offat.Finding, error)
	// ExpandChains returns additional candidate chains beyond what the
	// deterministic planner produced.
	ExpandChains(ctx context.Context, findings []offat.Finding, existing []offat.Chain) ([]offat.Chain, error)
}

// NoOp is the default provider; it returns empty results and never errors.
type NoOp struct{}

func (NoOp) Name() string { return "noop" }
func (NoOp) ClassifyEndpoint(context.Context, offat.Endpoint) (string, error) {
	return "", nil
}
func (NoOp) SynthesizePayloads(context.Context, offat.Endpoint, offat.Family) ([]string, error) {
	return nil, nil
}
func (NoOp) TriageFindings(_ context.Context, findings []offat.Finding) ([]offat.Finding, error) {
	return findings, nil
}
func (NoOp) ExpandChains(_ context.Context, _ []offat.Finding, existing []offat.Chain) ([]offat.Chain, error) {
	return existing, nil
}
