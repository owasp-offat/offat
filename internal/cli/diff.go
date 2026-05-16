package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newDiffCommand() *cobra.Command {
	var (
		base    string
		current string
	)

	cmd := &cobra.Command{
		Use:   "diff",
		Short: "Diff two specs/captures and surface only new attack surface",
		Long: `Diff compares a baseline spec or capture against the current one and
restricts subsequent scans to endpoints, parameters, or response shapes that
have changed. Useful in PR/CI mode.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if base == "" || current == "" {
				return fmt.Errorf("--base and --current are both required")
			}
			// TODO(M3): implement spec diff against the inventory family.
			fmt.Fprintf(cmd.OutOrStdout(), "diff: base=%s current=%s\n", base, current)
			return nil
		},
	}
	cmd.Flags().StringVar(&base, "base", "", "baseline spec or capture")
	cmd.Flags().StringVar(&current, "current", "", "current spec or capture")
	return cmd
}
