package cmd

import (
	"github.com/dimastark/configs/pkg/tool"
	"github.com/dimastark/configs/pkg/tools"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Runs all Cleanupers",
	Long:  "Command that runs all tools that can cleanup something",

	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		for _, t := range tools.Registry {
			if installer, ok := t.(tool.Installer); ok {
				if installed := installer.Installed(); !installed {
					color.Blue("> Install %s\n", t.Name())

					if err = installer.Install(); err != nil {
						return
					}
				}
			}

			if cleanuper, ok := t.(tool.Cleanuper); ok {
				color.Blue("> Cleanup %s\n", t.Name())

				if err = cleanuper.Cleanup(); err != nil {
					return
				}
			}
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
