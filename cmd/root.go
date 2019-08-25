package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "configs",
	Short: "dimastark configs",
	Long:  "My configs utilities for Mac OS X",

	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
