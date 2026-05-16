// Package plugin hosts out-of-process plugins over gRPC. We avoid Go's
// in-process plugin package because it is brittle across toolchains and
// platforms. Plugin authors ship a binary that speaks the OFFAT plugin
// protocol; the host launches it as a subprocess and routes Check calls
// across the wire.
package plugin
