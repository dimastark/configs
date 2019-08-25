package cmd

import (
	"github.com/dimastark/configs/pkg/tool"
	"github.com/dimastark/configs/pkg/tools"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Runs all Backupers",
	Long:  "Command that runs all tools that can backup something",

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

			if backuper, ok := t.(tool.Backuper); ok {
				color.Blue("> Backup %s\n", t.Name())

				if err = backuper.Backup(); err != nil {
					return
				}
			}
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
