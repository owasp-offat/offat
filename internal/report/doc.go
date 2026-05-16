// Package report renders findings into SARIF (GitHub code scanning), JSON,
// HTML, and JUnit. Each renderer implements the same interface so callers
// can swap formats without other changes.
package report

import (
	"io"

	"github.com/owasp-offat/offat/pkg/offat"
)

// Renderer writes findings (and optionally chains) to an io.Writer.
type Renderer interface {
	Render(w io.Writer, findings []offat.Finding, chains []offat.Chain) error
	Format() string // "sarif" | "json" | "html" | "junit"
}
