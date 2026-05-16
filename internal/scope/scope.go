// Package scope enforces target allowlists and destructive-verb gates.
// A scan refuses to run without a scope.
package scope

import (
	"net/url"
	"strings"

	"github.com/owasp-offat/offat/pkg/offat"
)

// Allowlist is a Scoper backed by a static host allowlist.
type Allowlist struct {
	hosts map[string]struct{}
}

// NewAllowlist returns an Allowlist containing the given hosts (case-insensitive).
func NewAllowlist(hosts []string) *Allowlist {
	m := make(map[string]struct{}, len(hosts))
	for _, h := range hosts {
		m[strings.ToLower(strings.TrimSpace(h))] = struct{}{}
	}
	return &Allowlist{hosts: m}
}

// InScope returns true if the URL's host is on the allowlist.
func (a *Allowlist) InScope(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	_, ok := a.hosts[strings.ToLower(u.Hostname())]
	return ok
}

var _ offat.Scoper = (*Allowlist)(nil)
