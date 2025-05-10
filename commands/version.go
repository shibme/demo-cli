package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev"
var ProjectName = "app"

func showVersionInfo() {
	fmt.Println(ProjectName, Version)
}

var versionCmd *cobra.Command

func versionCommand() *cobra.Command {
	if versionCmd != nil {
		return versionCmd
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Shows the version information",
		Run: func(cmd *cobra.Command, args []string) {
			showVersionInfo()
		},
	}
	return versionCmd
}
