package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var helloCmd *cobra.Command

func helloCommand() *cobra.Command {
	if helloCmd != nil {
		return helloCmd
	}
	helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "Says hello",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString(nameFlag.name)
			if name == "" {
				fmt.Println("Hello, World!")
			} else {
				fmt.Println("Hello " + color.GreenString(name) + "!")
			}
			os.Exit(0)
		},
	}
	helloCmd.Flags().StringP(nameFlag.name, nameFlag.shorthand, "", nameFlag.usage)
	return helloCmd
}
