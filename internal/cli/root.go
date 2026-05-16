// Package cli wires up the cobra command tree for the offat binary.
package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Version is set at build time via -ldflags.
var Version = "0.0.0-dev"

// NewRootCommand builds the root command and registers all subcommands.
func NewRootCommand() *cobra.Command {
	var configFile string

	root := &cobra.Command{
		Use:           "offat",
		Short:         "OFFensive Api Tester — agent-native API security scanner",
		Long:          rootLongHelp,
		Version:       Version,
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return initConfig(configFile)
		},
	}

	root.PersistentFlags().StringVar(&configFile, "config", "", "config file (default: ./offat.yaml or $HOME/.offat.yaml)")
	root.PersistentFlags().Bool("verbose", false, "enable verbose logging")
	_ = viper.BindPFlag("verbose", root.PersistentFlags().Lookup("verbose"))

	root.AddCommand(
		newScanCommand(),
		newReplayCommand(),
		newDiffCommand(),
		newChainCommand(),
		newReportCommand(),
		newServeMCPCommand(),
	)
	return root
}

func initConfig(configFile string) error {
	viper.SetEnvPrefix("OFFAT")
	viper.AutomaticEnv()
	if configFile == "" {
		return nil
	}
	viper.SetConfigFile(configFile)
	return viper.ReadInConfig()
}

const rootLongHelp = `OFFAT scans HTTP APIs for OWASP API Top 10 vulnerabilities, models findings
as a chainable attack graph, and exposes its primitives over MCP for use by
Claude Code and other agents.

Run "offat <subcommand> --help" for details on each command.`
