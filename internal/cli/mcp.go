package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newServeMCPCommand() *cobra.Command {
	var addr string

	cmd := &cobra.Command{
		Use:   "serve-mcp",
		Short: "Run OFFAT as an MCP server (M2)",
		Long: `serve-mcp exposes OFFAT's primitives (scan, replay, fuzz, chain, report,
record_traffic) as MCP tools so Claude Code and other agents can drive the
scanner. Stubbed in M1; implementation lands in M2.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// TODO(M2): start MCP server over stdio or addr.
			fmt.Fprintf(cmd.OutOrStdout(), "serve-mcp: not yet implemented (M2). addr=%s\n", addr)
			return nil
		},
	}
	cmd.Flags().StringVar(&addr, "addr", "", "listen address; empty means stdio")
	return cmd
}
