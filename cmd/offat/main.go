// Command offat is the OFFAT (OFFensive Api Tester) CLI entrypoint.
//
// The single binary hosts every subcommand: scan, replay, diff, chain, report,
// serve-mcp, version. See `offat --help`.
package main

import (
	"fmt"
	"os"

	_ "github.com/owasp-offat/offat/internal/check/all"
	"github.com/owasp-offat/offat/internal/cli"
)

func main() {
	if err := cli.NewRootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "offat:", err)
		os.Exit(1)
	}
}
