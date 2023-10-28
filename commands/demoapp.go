package commands

import "github.com/spf13/cobra"

var demoAppCmd *cobra.Command

func DemoAppCommand() *cobra.Command {
	if demoAppCmd != nil {
		return demoAppCmd
	}
	demoAppCmd = &cobra.Command{
		Use:   "demo",
		Short: "Demo CLI is an CLI tool built for demo purposes",
		Run: func(cmd *cobra.Command, args []string) {
			version, _ := cmd.Flags().GetBool(versionFlag.name)
			if version {
				showVersionInfo()
			} else {
				cmd.Help()
			}
		},
	}
	demoAppCmd.Flags().BoolP(versionFlag.name, versionFlag.shorthand, false, versionFlag.usage)
	demoAppCmd.AddCommand(versionCommand())
	demoAppCmd.AddCommand(helloCommand())
	return demoAppCmd
}
