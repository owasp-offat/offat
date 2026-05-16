package scope

import "testing"

func TestAllowlist(t *testing.T) {
	a := NewAllowlist([]string{"api.example.com", "internal.test"})

	cases := map[string]bool{
		"https://api.example.com/v1/users": true,
		"https://API.example.com/v1":       true, // case-insensitive
		"http://internal.test:8080/health": true,
		"https://evil.example.com/":        false,
		"not a url":                        false,
	}
	for raw, want := range cases {
		if got := a.InScope(raw); got != want {
			t.Errorf("InScope(%q) = %v, want %v", raw, got, want)
		}
	}
}
