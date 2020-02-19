package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"os"
)

type RootCmd struct {
	root *cli.App
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = RootCmd{
	root: newCmdRoot(),
}

func newCmdRoot() *cli.App {
	return &cli.App{Name: "oefi", EnableBashCompletion: true,}
}

func AddSubcommand(cmd *cli.Command) {
	rootCmd.root.Commands = append(rootCmd.root.Commands, cmd)
}

func Execute() {
	if err := rootCmd.root.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
