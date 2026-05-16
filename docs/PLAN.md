# OFFAT Revamp Plan

This is the canonical, committed milestone plan for the OFFAT rebuild on the
`revamp` branch. The companion design discussion lives outside the repo; this
document is what the codebase tracks against.

## Goals

- Cover **OWASP API Security Top 10 (2023)** as testable primitives.
- Support OpenAPI 3.x / Swagger 2, HAR / mitmproxy captures, GraphQL
  introspection, gRPC reflection as inputs.
- Generate tests four ways: spec-driven, traffic-driven, state-aware
  (CRUD graph), schema-aware fuzzing — with LLM synthesis pluggable on top.
- Model findings as a **chainable attack graph**, not a list.
- Provide an **MCP server** so Claude Code can drive scans, fuzz, replay,
  chain, and report as tools.
- **Deterministic mode** (no LLM) is always supported — AI is an enhancement
  layer, never a requirement.
- Outputs: SARIF, JSON, HTML, JUnit.
- Built-in **self-hosted OOB collaborator** (DNS + HTTP) for blind
  SSRF/XXE/log4shell-class detection.
- **Safe by default**: scope allowlists, destructive-verb gates, dry-run,
  idempotency awareness, per-host rate limits.

## Non-goals (initially)

- Web UI / dashboard (CLI + reports only).
- Multi-tenant SaaS hosting.
- Active exploitation beyond proof-of-vulnerability — no shells, no data
  exfil beyond markers.
- Network-layer (L3/L4) scanning — this is L7 API only.

## Milestones

### M1 — Deterministic Top 10 (no LLM dependency)

1. Repo scaffold: cobra/viper CLI, `internal/` layout, golangci-lint, `make test`. ✅ (scaffold)
2. OpenAPI 3.x + Swagger 2 ingest via `kin-openapi`; HAR ingest.
3. HTTP engine with auth rotation, recording, retry, rate limit, scope guard.
4. Auth recipes: bearer, API key (header/query/cookie), OAuth2 client
   credentials, basic. SAML/SigV4 deferred to M1.5.
5. Implement all 10 check families deterministically. Each ships with
   table-driven unit tests and an integration target against
   docker-composed vAPI / crAPI / juice-shop.
6. Finding graph + rule-based chain planner with the primitive vocabulary
   in `pkg/offat/primitive.go`.
7. OOB collaborator (DNS + HTTP), self-hostable; ship a `docker-compose.yml`.
8. Outputs: SARIF, JSON, HTML, JUnit.
9. `ai.Provider` interface defined; M1 ships `NoOp`. ✅
10. CI: GitHub Action that runs `offat scan` against a sample spec and
    uploads SARIF.

### M2 — Claude Code first-class

- `offat serve-mcp` MCP server.
- Claude provider behind `ai.Provider`: triage, payload synthesis, chain
  expansion, spec inference from traffic.
- Agentic crawler-steering loop using the Claude Agent SDK.

### M3 — Extensibility & ecosystem

- Out-of-process gRPC plugin host (avoids Go's plugin pkg pain).
- GraphQL deep checks (alias batching, depth/cost bypass, field suggestion
  leaks).
- gRPC reflection scanner + protobuf-aware probes.
- Browser-driven auth flows (chromedp/rod).
- PR/CI diff mode against previous spec.

## Verification (M1 acceptance)

1. `git branch --show-current` returns `revamp`; `main` HEAD unchanged.
2. `go build ./...` succeeds on a fresh clone.
3. `make test` runs the unit suite; `make test-integration` runs against
   vAPI in docker-compose and produces a SARIF file.
4. `offat scan -f testdata/openapi.yaml --scope localhost --dry-run` lists
   planned requests without sending.
5. `offat scan -f testdata/openapi.yaml --scope localhost -o out.sarif`
   finds at least the seeded BOLA + BOPLA + misconfig issues in vAPI.
6. `offat chain --from out.json` emits at least one multi-step chain on
   vAPI.
7. `CLAUDE.md` and `docs/PLAN.md` exist; `docs/ARCHITECTURE.md` and
   `docs/CHECKS.md` exist as skeletons. ✅
8. CI workflow on `revamp` runs the scan against `testdata/` and uploads
   SARIF as a workflow artifact.

## Open questions (revisit before/during M1)

- Final name on the binary: keep `offat`, or rename?
- License — keep MIT, or relicense to Apache-2.0 for a friendlier plugin
  ecosystem?
- Default OOB collaborator: ship a hosted one, or require self-host from day
  one?
- vAPI vs crAPI vs both as the integration test target?
