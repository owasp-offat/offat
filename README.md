# OFFAT — OFFensive Api Tester (revamp)

> **Status:** Pre-alpha rebuild on the `revamp` branch. `main` carries the
> legacy code and is frozen until M1 ships.

OFFAT is an agent-native API security scanner in Go. It reads OpenAPI specs,
HAR captures, GraphQL introspection, or gRPC reflection, runs the OWASP API
Security Top 10 (2023) as concrete test primitives, models the resulting
findings as a chainable attack graph, and (in M2) exposes everything to
Claude Code via MCP.

## Why this rebuild

The 2024 codebase was built on `fasthttp` and a small fixed set of injection
checks. The revamp is:

- **Comprehensive** — every OWASP API Top 10 (2023) class as testable primitives.
- **Chain-aware** — findings declare composable primitives; a planner walks the graph to surface multi-step exploits.
- **Agent-native** — first-class MCP server so Claude Code (and other agents) can drive scans, fuzz, replay, chain, and report.
- **Safe by default** — required scope allowlist, destructive-verb gating, dry-run, self-hosted OOB collaborator.
- **Deterministic without AI** — the LLM provider is pluggable; the tool runs unaided.

## OWASP API Top 10 coverage (M1)

| API# | Class                                    | Package                          |
| ---- | ---------------------------------------- | -------------------------------- |
| API1 | Broken Object Level Authorization        | `internal/check/bola`            |
| API2 | Broken Authentication                    | `internal/check/brokenauth`      |
| API3 | Broken Object Property Level Auth (BOPLA)| `internal/check/bopla`           |
| API4 | Unrestricted Resource Consumption        | `internal/check/consumption`     |
| API5 | Broken Function Level Authorization      | `internal/check/bfla`            |
| API6 | Unrestricted Access to Sensitive Flows   | `internal/check/bizflow`         |
| API7 | Server Side Request Forgery              | `internal/check/ssrf`            |
| API8 | Security Misconfiguration                | `internal/check/misconfig`       |
| API9 | Improper Inventory Management            | `internal/check/inventory`       |
| API10| Unsafe Consumption of APIs               | `internal/check/upstream`        |

## Quick start

```bash
make build
./offat --help
./offat scan -f testdata/openapi-petstore.yml --scope api.example.com --dry-run
```

## Documentation

- `docs/PLAN.md` — milestone plan.
- `docs/ARCHITECTURE.md` — module layout, interfaces, dataflow.
- `docs/CHECKS.md` — per-check spec (id, family, primitives, behavior).
- `CLAUDE.md` — working notes for Claude Code contributors.

## License

MIT — see `LICENSE`. (Open question for M1: relicense to Apache-2.0 for a
friendlier plugin ecosystem; not decided.)
