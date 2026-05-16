// Package runner is the scan orchestrator: it pulls Endpoints from
// spec/ingest, asks each registered Check for TestCases, dispatches them
// through httpx with rate limiting and scope enforcement, records findings,
// and runs the graph.Planner over the result.
package runner
