package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newChainCommand() *cobra.Command {
	var (
		from string
		out  string
	)

	cmd := &cobra.Command{
		Use:   "chain",
		Short: "Run the chain planner over existing findings",
		Long: `Chain consumes a findings file produced by "offat scan" and runs the
rule-based planner over the primitive graph to propose multi-step exploits.

In M2 the planner can be augmented by Claude Code via the MCP server.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if from == "" {
				return fmt.Errorf("--from is required (path to findings JSON)")
			}
			// TODO(M1): load findings -> graph.Planner -> emit chains.
			fmt.Fprintf(cmd.OutOrStdout(), "chain: from=%s out=%s\n", from, out)
			return nil
		},
	}
	cmd.Flags().StringVar(&from, "from", "", "findings JSON file produced by `offat scan`")
	cmd.Flags().StringVarP(&out, "output", "o", "", "chains output file (defaults to stdout)")
	return cmd
}
