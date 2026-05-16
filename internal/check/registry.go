// Package check hosts the registry that all check-family packages register
// into via init(). The runner iterates the registry to build a scan plan.
package check

import (
	"sort"
	"sync"

	"github.com/owasp-offat/offat/pkg/offat"
)

var (
	mu       sync.RWMutex
	registry = map[string]offat.Check{}
)

// Register adds a check to the global registry. Intended for use from a
// check family's init().
func Register(c offat.Check) {
	mu.Lock()
	defer mu.Unlock()
	registry[c.ID()] = c
}

// All returns every registered check, sorted by ID for stable iteration.
func All() []offat.Check {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]offat.Check, 0, len(registry))
	for _, c := range registry {
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID() < out[j].ID() })
	return out
}

// ByFamily returns every registered check in the given family.
func ByFamily(f offat.Family) []offat.Check {
	out := []offat.Check{}
	for _, c := range All() {
		if c.Family() == f {
			out = append(out, c)
		}
	}
	return out
}
