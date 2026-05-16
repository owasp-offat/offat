// Package fuzz hosts payload corpora and schema-aware mutators used by
// checks to synthesize test bodies/params. Mutators include type confusion,
// null injection, prototype-pollution shapes, unicode normalization tricks,
// and JSON nesting bombs.
package fuzz
