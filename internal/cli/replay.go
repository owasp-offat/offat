package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newReplayCommand() *cobra.Command {
	var (
		source string
		actor  string
	)

	cmd := &cobra.Command{
		Use:   "replay",
		Short: "Replay a captured request under one or more auth contexts",
		Long: `Replay loads a previously captured request (from a corpus file or a HAR
entry) and re-issues it as the chosen actor. Use it to confirm a finding by
hand or to bulk-test cross-actor access.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if source == "" {
				return fmt.Errorf("--source is required")
			}
			// TODO(M1): wire corpus -> httpx -> store.
			fmt.Fprintf(cmd.OutOrStdout(), "replay: source=%s actor=%s\n", source, actor)
			return nil
		},
	}
	cmd.Flags().StringVar(&source, "source", "", "corpus file or HAR entry id")
	cmd.Flags().StringVar(&actor, "actor", "", "actor name to replay as (from --actors file)")
	return cmd
}
