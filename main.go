package main

import (
	"os"

	"github.com/shibme/demo-cli/commands"
)

func main() {
	if err := commands.DemoAppCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
