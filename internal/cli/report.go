package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newReportCommand() *cobra.Command {
	var (
		from   string
		out    string
		format string
	)

	cmd := &cobra.Command{
		Use:   "report",
		Short: "Render findings in the chosen format",
		Long: `Report converts a findings JSON file to SARIF, HTML, JUnit, or
pretty-printed JSON. Use it to re-render reports without re-running a scan.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if from == "" {
				return fmt.Errorf("--from is required")
			}
			// TODO(M1): pick renderer based on --format.
			fmt.Fprintf(cmd.OutOrStdout(), "report: from=%s out=%s format=%s\n", from, out, format)
			return nil
		},
	}
	cmd.Flags().StringVar(&from, "from", "", "findings JSON file")
	cmd.Flags().StringVarP(&out, "output", "o", "", "output file (defaults to stdout)")
	cmd.Flags().StringVar(&format, "format", "sarif", "json, sarif, html, junit")
	return cmd
}
