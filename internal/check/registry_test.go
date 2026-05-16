package check

import (
	"testing"

	"github.com/owasp-offat/offat/pkg/offat"
)

type stubCheck struct {
	id     string
	family offat.Family
}

func (s stubCheck) ID() string                                          { return s.id }
func (s stubCheck) Family() offat.Family                                { return s.family }
func (s stubCheck) Applicable(offat.Endpoint) bool                      { return true }
func (s stubCheck) Generate(offat.CheckContext, offat.Endpoint) []offat.TestCase {
	return nil
}
func (s stubCheck) Verify(offat.Request, offat.Response) []offat.Finding { return nil }

func TestRegistry(t *testing.T) {
	// Snapshot registry to avoid clashing with init() registrations.
	before := len(All())

	Register(stubCheck{id: "test.alpha", family: offat.FamilyBOLA})
	Register(stubCheck{id: "test.beta", family: offat.FamilyBFLA})

	if got := len(All()) - before; got != 2 {
		t.Fatalf("expected 2 new checks, got %d", got)
	}
	if got := len(ByFamily(offat.FamilyBOLA)); got < 1 {
		t.Fatalf("expected BOLA family non-empty, got %d", got)
	}
}
