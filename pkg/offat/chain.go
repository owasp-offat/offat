package offat

// Chain is a sequence of findings the planner believes compose into a higher
// impact attack than any single finding on its own.
type Chain struct {
	ID       string    // stable id for this chain instance
	Title    string    // short headline (e.g. "SSRF → IMDS → AWS keys")
	Severity Severity  // strongest severity in the chain, escalated if appropriate
	Findings []string  // Finding.IDs in order
	Steps    []string  // human-readable steps, one per hop
	Notes    string    // why this chain matters
}

// Planner walks the finding graph and proposes chains.
type Planner interface {
	Propose(findings []Finding) []Chain
}
