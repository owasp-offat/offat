# OFFAT — Working notes for Claude Code

## What this project is

Agent-native automated API security scanner in Go. Rebuild of OWASP OFFAT on
the `revamp` branch. `main` is the legacy code — do not touch it. Ship
deterministic OWASP API Top 10 coverage first (M1); MCP server + Claude
provider land in M2. AI is an enhancement layer, never a dependency.

## Repo layout

- `cmd/offat/` — single binary entrypoint.
- `internal/cli/` — cobra/viper command tree: `scan`, `replay`, `diff`, `chain`, `report`, `serve-mcp`.
- `internal/spec/` — OpenAPI 3.x / Swagger 2 parsing.
- `internal/ingest/` — HAR, mitmproxy, Postman/Bruno, GraphQL introspection, gRPC reflection.
- `internal/httpx/` — `net/http` with a RoundTripper stack (auth, recording, OOB, retry, rate-limit, scope).
- `internal/auth/` — recipe library (bearer, API key, basic, OAuth2 cc in M1; SAML/mTLS/SigV4/HMAC later).
- `internal/browser/` — chromedp/rod for interactive auth flows (M3).
- `internal/check/<family>/` — one package per OWASP API Top 10 family. Every check `init()`s into `internal/check`'s registry; `internal/check/all` blank-imports them all.
- `internal/fuzz/` — payload corpora and schema-aware mutators.
- `internal/graph/` — rule-based chain planner; AI-augmented planner in M2.
- `internal/oob/` — DNS + HTTP collaborator (self-hosted, never a public service).
- `internal/corpus/` — replay corpus, regression export.
- `internal/scope/` — host allowlists, destructive-verb gates.
- `internal/ratelimit/` — token-bucket per host, `Retry-After` aware.
- `internal/store/` — SQLite via `modernc.org/sqlite` (CGO-free).
- `internal/report/` — SARIF, JSON, HTML, JUnit renderers.
- `internal/plugin/` — out-of-process gRPC plugin host (M3).
- `internal/ai/` — Provider interface + NoOp; Claude provider in M2.
- `internal/mcp/` — MCP server (M2).
- `internal/runner/` — scan orchestrator.
- `pkg/offat/` — public types for plugin authors (Check, Finding, Primitive, Endpoint, Chain).
- `docs/` — `PLAN.md`, `ARCHITECTURE.md`, `CHECKS.md`.
- `testdata/` — sample specs/HAR for unit tests.
- `test/` — integration suites (vulnerable-target docker-compose).

## Conventions

- Go 1.22+. `gofmt`, `goimports`, `golangci-lint run` clean before commit.
- Standard library `net/http` only. **No `fasthttp`.** HTTP/3 via `quic-go` opt-in.
- SQLite via `modernc.org/sqlite` (CGO-free) for portable builds.
- Tests: table-driven; integration tests run against docker-composed vAPI / crAPI / juice-shop.
  - `make test` runs unit; `make test-integration` boots containers.
- Errors: wrap with `%w`. No `panic` outside `main`. Context-cancelable I/O everywhere.
- Logging: `log/slog`, structured, leveled. No `fmt.Println` in library code.
- Public types live in `pkg/offat/`. Anything in `internal/` is not API.

## Safety rules (hard)

- Never run a scan without an explicit `--scope` allowlist.
- Destructive verbs (`DELETE`, non-idempotent `PUT`) gated behind `--allow-destructive`.
- OOB collaborator domain comes from config; never hardcoded to a public service.
- Every outbound test request carries `X-Offat-Run-ID` for log correlation.
- No active exploitation beyond proof-of-vulnerability — no shells, no data exfil beyond markers.

## When adding a new check

1. Pick (or create) a family package under `internal/check/<family>/`.
2. Implement the `offat.Check` interface (`ID`, `Family`, `Applicable`, `Generate`, `Verify`).
3. Declare the `offat.Primitive`s your findings expose so the chain planner can use them.
4. `init()` calls `check.Register(yourCheck{})`.
5. Add a table-driven unit test next to it; add an integration target under `test/` if relevant.
6. Update `docs/CHECKS.md` with the check id, family, primitives, and a one-line description.

## When adding a new primitive

The vocabulary in `pkg/offat/primitive.go` is intentionally small and stable.
Adding one means other checks may consume it — review in `docs/CHECKS.md` and
keep the name in the `verb:noun` shape (e.g. `leaks:user_uuid`).

## AI integration boundaries

All LLM calls go through `internal/ai/Provider`. The `NoOp` implementation must
always satisfy the same behavior contract — a feature that "only works with
AI" is a bug. Claude is for ranking, synthesis, and chain expansion, not for
correctness.

## MCP server (M2)

`offat serve-mcp` exposes `scan`, `replay`, `fuzz`, `chain`, `report`,
`record_traffic` as MCP tools. Any new CLI capability OFFAT ships should also
be exposed via MCP — agent-native parity is a project rule.

## Branches

- `main` — legacy OFFAT. Frozen until revamp ships.
- `revamp` — active rebuild. All PRs target this until M1 is feature-complete.

## Commands cheatsheet

```bash
make build         # build the offat binary
make test          # unit tests
make lint          # golangci-lint
make tidy          # go mod tidy
./offat --help
./offat scan -f testdata/openapi-petstore.yml --scope api.example.com --dry-run
```
