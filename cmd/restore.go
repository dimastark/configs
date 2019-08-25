package cmd

import (
	"github.com/dimastark/configs/pkg/tool"
	"github.com/dimastark/configs/pkg/tools"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Runs all Restorers",
	Long:  "Command that runs all tools that can restore something",

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

			if restorer, ok := t.(tool.Restorer); ok {
				color.Blue("> Restore %s\n", t.Name())

				if err = restorer.Restore(); err != nil {
					return
				}
			}
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
