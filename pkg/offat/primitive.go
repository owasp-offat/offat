package offat

// Primitive is a capability a finding exposes to the chain planner. The
// vocabulary is intentionally small and stable so checks and plugins can
// interoperate. New primitives must be reviewed in docs/CHECKS.md.
type Primitive string

const (
	// Information leaks — values usable as inputs to other probes.
	PrimLeaksUserUUID        Primitive = "leaks:user_uuid"
	PrimLeaksTenantID        Primitive = "leaks:tenant_id"
	PrimLeaksInternalHost    Primitive = "leaks:internal_hostname"
	PrimLeaksStackTrace      Primitive = "leaks:stack_trace"
	PrimLeaksCredentials     Primitive = "leaks:credentials"
	PrimLeaksSecretMaterial  Primitive = "leaks:secret_material"
	PrimLeaksSchemaInternals Primitive = "leaks:schema_internals"

	// Input handling — what the endpoint accepts.
	PrimAcceptsArbitraryURL Primitive = "accepts:arbitrary_url"
	PrimAcceptsFilePath     Primitive = "accepts:file_path"
	PrimAcceptsHTML         Primitive = "accepts:html"
	PrimAcceptsTemplate     Primitive = "accepts:template_expression"
	PrimAcceptsSQLFragment  Primitive = "accepts:sql_fragment"

	// Reflection / persistence.
	PrimReflectsInput Primitive = "reflects:input"
	PrimStoresInput   Primitive = "stores:input"

	// Authorization / authentication state.
	PrimAuthBypassable    Primitive = "auth:bypassable"
	PrimAuthDowngradable  Primitive = "auth:downgradable"
	PrimAuthMissing       Primitive = "auth:missing"
	PrimAuthRoleEscalable Primitive = "auth:role_escalable"

	// Enumeration.
	PrimEnumeratesUser   Primitive = "enumerates:user"
	PrimEnumeratesTenant Primitive = "enumerates:tenant"
	PrimEnumeratesObject Primitive = "enumerates:object"
)
