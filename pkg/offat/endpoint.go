package offat

// Endpoint is one operation on a target API, normalized from whatever input
// source produced it (OpenAPI op, HAR entry, captured request, etc).
type Endpoint struct {
	Method      string            // GET, POST, PUT, DELETE, PATCH, ...
	URL         string            // fully-qualified URL with templated path params still in place
	PathParams  map[string]string // declared path parameters and example values
	QueryParams map[string]string // declared query parameters and example values
	Headers     map[string]string // baseline request headers (excluding auth)
	Body        []byte            // example/seed body
	ContentType string            // request content type, if any
	OperationID string            // OpenAPI operationId or synthesized id
	Tags        []string          // OpenAPI tags or inferred labels
	AuthSchemes []string          // names of auth schemes that apply (referencing AuthRecipes)
	SourceSpec  string            // path or URL of the source spec/capture
}

// Request is a single HTTP request OFFAT will send (or has sent).
type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []byte
}

// Response is the captured response for a Request.
type Response struct {
	Status     int
	Headers    map[string]string
	Body       []byte
	DurationMS int64
}
