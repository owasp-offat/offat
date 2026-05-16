package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newScanCommand() *cobra.Command {
	var (
		spec             string
		scopeFile        string
		output           string
		format           string
		rateLimit        int
		concurrency      int
		dryRun           bool
		allowDestructive bool
		actorsFile       string
		oobDomain        string
	)

	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan an API target using a spec or capture as input",
		Long: `Scan ingests an OpenAPI/Swagger spec (or HAR / Postman capture), runs the
configured OWASP API Top 10 checks against the target, and writes findings
to the chosen output format.

The --scope allowlist is required. Destructive verbs are blocked unless
--allow-destructive is set. --dry-run prints planned requests without sending
any traffic.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if scopeFile == "" {
				return fmt.Errorf("--scope is required; refusing to run without an allowlist")
			}
			if spec == "" {
				return fmt.Errorf("-f/--file is required (OpenAPI spec, HAR, or collection)")
			}
			// TODO(M1): wire ingest -> spec -> scope -> runner -> store -> report.
			fmt.Fprintf(cmd.OutOrStdout(),
				"scan: spec=%s scope=%s format=%s output=%s rate=%d conc=%d dryRun=%v destructive=%v actors=%s oob=%s\n",
				spec, scopeFile, format, output, rateLimit, concurrency, dryRun, allowDestructive, actorsFile, oobDomain)
			return nil
		},
	}

	cmd.Flags().StringVarP(&spec, "file", "f", "", "spec or capture (OpenAPI YAML/JSON, HAR, Postman) — path or URL")
	cmd.Flags().StringVar(&scopeFile, "scope", "", "scope allowlist file (required)")
	cmd.Flags().StringVarP(&output, "output", "o", "", "output file (defaults to stdout)")
	cmd.Flags().StringVar(&format, "format", "json", "output format: json, sarif, html, junit")
	cmd.Flags().IntVarP(&rateLimit, "rate", "r", 50, "max requests per second per host")
	cmd.Flags().IntVarP(&concurrency, "concurrency", "c", 16, "global worker concurrency")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "print planned requests; send nothing")
	cmd.Flags().BoolVar(&allowDestructive, "allow-destructive", false, "permit DELETE / unsafe PUT probes")
	cmd.Flags().StringVar(&actorsFile, "actors", "", "actors/auth file with role-scoped credentials")
	cmd.Flags().StringVar(&oobDomain, "oob-domain", "", "OOB collaborator domain (self-hosted)")
	return cmd
}
