package offat

import "time"

// Evidence is the minimal data captured to justify a finding.
type Evidence struct {
	Request    Request
	Response   Response
	Notes      string   // short human-readable summary of why this is a finding
	Indicators []string // signal names that fired (e.g. "cors-wildcard-with-credentials")
}

// ReproSteps lists the steps a human or planner can follow to reproduce the finding.
type ReproSteps struct {
	CurlCommand string   // single-shot reproducer where applicable
	Steps       []string // ordered prose steps; chains may have several
}

// Finding is OFFAT's atomic result. Findings are nodes in the attack graph;
// the Primitives slice is the edge vocabulary the planner walks.
type Finding struct {
	ID         string     // stable id for this finding occurrence
	CheckID    string     // e.g. "bola.cross-actor-id-replay"
	Family     Family     // OWASP API Top 10 grouping
	Endpoint   Endpoint   // the operation tested
	Severity   Severity   // info..critical
	Title      string     // short title for reports
	Evidence   Evidence   // request/response and indicator names
	Primitives []Primitive
	Repro      ReproSteps
	DetectedAt time.Time
	RunID      string // matches the X-Offat-Run-ID marker
}
