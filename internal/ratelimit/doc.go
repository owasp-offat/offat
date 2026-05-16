// Package ratelimit implements per-host token-bucket rate limiting with
// Retry-After awareness. The engine consults a Limiter before every outbound
// request and yields if the target asks it to slow down.
package ratelimit
