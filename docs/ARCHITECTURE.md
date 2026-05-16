# Architecture

## Dataflow (M1)

```
spec/ingest ──► []Endpoint ──► runner ──► check.Generate ──► []TestCase
                                                 │
                                                 ▼
                                              httpx (auth + rate-limit + scope + record)
                                                 │
                                                 ▼
                                            check.Verify ──► []Finding ──► store
                                                                            │
                                                                            ▼
                                                                      graph.Planner ──► []Chain
                                                                            │
                                                                            ▼
                                                                       report (SARIF/JSON/HTML/JUnit)
```

The `ai.Provider` interface sits beside `runner` and `graph.Planner`. In M1
the `NoOp` provider is wired in; in M2 the Claude provider augments
classification, payload synthesis, triage, and chain expansion.

## Package map

| Package                        | Responsibility                                                    |
| ------------------------------ | ----------------------------------------------------------------- |
| `cmd/offat`                    | binary entrypoint                                                 |
| `internal/cli`                 | cobra/viper command tree                                          |
| `internal/spec`                | OpenAPI / Swagger parsing                                         |
| `internal/ingest`              | HAR / Postman / GraphQL / gRPC ingest                             |
| `internal/httpx`               | RoundTripper stack (auth, recording, OOB, retry, scope)           |
| `internal/auth`                | recipe library                                                    |
| `internal/browser`             | chromedp/rod for SSO/captcha auth flows (M3)                      |
| `internal/check/<family>`      | per-family checks; register into the global registry              |
| `internal/check/all`           | blank-imports every family so `init()` registrations fire         |
| `internal/fuzz`                | payload corpora and schema-aware mutators                         |
| `internal/graph`               | rule-based chain planner (M1); LLM-augmented (M2)                 |
| `internal/oob`                 | DNS + HTTP collaborator                                           |
| `internal/corpus`              | replay corpus, regression export                                  |
| `internal/scope`               | host allowlist, destructive-verb gates                            |
| `internal/ratelimit`           | per-host token-bucket                                             |
| `internal/store`               | SQLite (modernc.org/sqlite, CGO-free)                             |
| `internal/report`              | SARIF / JSON / HTML / JUnit renderers                             |
| `internal/plugin`              | out-of-process gRPC plugin host (M3)                              |
| `internal/ai`                  | LLM provider interface + NoOp                                     |
| `internal/mcp`                 | MCP server (M2)                                                   |
| `internal/runner`              | scan orchestrator                                                 |
| `pkg/offat`                    | public types: Check, Finding, Primitive, Endpoint, Chain          |

## Key interfaces

```go
type Check interface {
    ID() string
    Family() Family
    Applicable(ep Endpoint) bool
    Generate(ctx CheckContext, ep Endpoint) []TestCase
    Verify(req Request, resp Response) []Finding
}

type Planner interface {
    Propose(findings []Finding) []Chain
}

type Provider interface {       // internal/ai
    Name() string
    ClassifyEndpoint(ctx, ep) (string, error)
    SynthesizePayloads(ctx, ep, family) ([]string, error)
    TriageFindings(ctx, findings) ([]Finding, error)
    ExpandChains(ctx, findings, existing) ([]Chain, error)
}
```

## Concurrency model

- Bounded worker pool sized by `--concurrency`.
- Per-host token-bucket limiter; respects `Retry-After`.
- Backoff on `429` and `5xx`.
- All I/O is `context.Context`-cancellable.

## HTTP stack

`net/http` with a custom `RoundTripper` chain:

1. **Scope guard** — short-circuits requests to out-of-scope hosts.
2. **Auth recipe** — applies the current actor's credentials.
3. **OOB header injection** — adds the configured collaborator hostname to
   selected fields when relevant.
4. **Rate limiter** — blocks until the bucket grants.
5. **Recorder** — captures request/response into the corpus.
6. **Retry** — with `Retry-After` and bounded backoff.
7. **Transport** — HTTP/2 by default; HTTP/3 opt-in via `quic-go` when the
   target advertises ALPN `h3`.

## Storage

SQLite via `modernc.org/sqlite` (CGO-free) for portability. Schemas:

- `runs` — one row per scan invocation.
- `findings` — finding rows, linked to the run.
- `requests` — recorded request/response pairs (the corpus).
- `chains` — planner output, linked to constituent findings.
