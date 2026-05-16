package offat

import "context"

// TestCase is a single concrete probe a Check wants OFFAT to run.
type TestCase struct {
	Name    string  // human-readable test name, unique within the check
	Request Request // the request to send
	// Verify, if non-nil, is called with the actual Response. When nil, the
	// engine calls the parent Check's Verify with the original Request and the
	// Response.
	Verify func(req Request, resp Response) []Finding `json:"-"`
}

// CheckContext is the runtime view passed into Generate. It exposes the auth
// recipes, scope info, and any AI provider the engine has wired up. Concrete
// values live in internal/runner — this is the contract checks rely on.
type CheckContext interface {
	Ctx() context.Context
	RunID() string
	Scope() Scoper
	// Actor returns the configured actor by name (e.g. "admin", "user_a").
	Actor(name string) (Actor, bool)
}

// Scoper enforces target allowlists.
type Scoper interface {
	InScope(rawURL string) bool
}

// Actor is one authentication identity used during a scan.
type Actor struct {
	Name    string
	Headers map[string]string // applied to outbound requests (e.g. Authorization)
	Cookies map[string]string
	Role    string // free-form role label (e.g. "admin", "user")
}

// Check is the unit of testing OFFAT runs against an Endpoint.
type Check interface {
	// ID is a stable, dotted identifier (family.short-name).
	ID() string
	// Family groups the check under an OWASP API Top 10 category.
	Family() Family
	// Applicable returns true if this check has anything to do on the endpoint.
	Applicable(ep Endpoint) bool
	// Generate returns the probes to run. May be empty.
	Generate(ctx CheckContext, ep Endpoint) []TestCase
	// Verify is the default verifier used for TestCases whose own Verify is nil.
	Verify(req Request, resp Response) []Finding
}
