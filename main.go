package main

import (
	"os"

	"dev.shib.me/demo-cli/commands"
)

func main() {
	if err := commands.DemoAppCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
