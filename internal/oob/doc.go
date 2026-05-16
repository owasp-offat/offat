// Package oob is OFFAT's out-of-band collaborator: a self-hosted DNS + HTTP
// listener used to detect blind SSRF, XXE, and log4shell-class issues.
//
// The collaborator domain is required to be configured; OFFAT never falls
// back to a public hosted service to avoid leaking target data.
package oob
