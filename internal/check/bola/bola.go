// Package bola implements API1:2023 Broken Object Level Authorization
// checks: capture object IDs from actor A's session and replay every
// reference shape against actor B's token.
package bola

import (
	"github.com/owasp-offat/offat/internal/check"
	"github.com/owasp-offat/offat/pkg/offat"
)

type crossActorIDReplay struct{}

func (crossActorIDReplay) ID() string             { return "bola.cross-actor-id-replay" }
func (crossActorIDReplay) Family() offat.Family   { return offat.FamilyBOLA }
func (crossActorIDReplay) Applicable(offat.Endpoint) bool { return false } // M1: implement
func (crossActorIDReplay) Generate(offat.CheckContext, offat.Endpoint) []offat.TestCase {
	return nil
}
func (crossActorIDReplay) Verify(offat.Request, offat.Response) []offat.Finding {
	return nil
}

func init() { check.Register(crossActorIDReplay{}) }
