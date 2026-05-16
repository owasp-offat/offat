# Checks Catalogue

Every check OFFAT ships is listed here with its id, OWASP family, the
primitives its findings expose to the chain planner, and a one-line
description. Updates land alongside the check itself.

## Primitive vocabulary

See `pkg/offat/primitive.go` for the authoritative list. Adding a new
primitive requires review; the planner consumes them, so the vocabulary is
kept small.

## Checks (M1 skeleton)

| Check ID                                  | Family       | Primitives produced | Notes |
| ----------------------------------------- | ------------ | ------------------- | ----- |
| `bola.cross-actor-id-replay`              | bola         | `enumerates:object` | Replay actor A's IDs as actor B across compatible endpoints. |
| `bopla.privileged-field-injection`        | bopla        | `auth:role_escalable` | Inject `role`, `isAdmin`, `tenant_id`, etc., into writes. |
| `brokenauth.jwt-alg-none`                 | brokenauth   | `auth:bypassable`   | Accept tokens with `alg=none`. |
| `bfla.verb-swap`                          | bfla         | `auth:bypassable`   | Switch `GET`→`PUT`/`DELETE` on protected paths. |
| `consumption.oversized-limit`             | consumption  | —                   | `?limit=999999` style amplification. |
| `bizflow.rate-limit-header-rotation`      | bizflow      | —                   | Bypass rate limits via `X-Forwarded-For` rotation. |
| `ssrf.cloud-imds`                         | ssrf         | `accepts:arbitrary_url`, `leaks:secret_material` | Probe `169.254.169.254` and cloud metadata variants. |
| `misconfig.cors-wildcard-with-credentials`| misconfig    | —                   | `Access-Control-Allow-Origin: *` with credentials. |
| `inventory.shadow-version`                | inventory    | —                   | `/v1` still alive when `/v2` is the documented version. |
| `upstream.webhook-response-fuzz`          | upstream     | —                   | Fuzz responses returned to the target's webhook receivers. |

The above is the M1 _skeleton_ list. Each check is currently a stub
registered with the global registry; concrete `Applicable`, `Generate`, and
`Verify` implementations are tracked per check.
