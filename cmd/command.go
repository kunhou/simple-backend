package cmd

import (
	"github.com/urfave/cli/v2"
)

var (
	cliApp *cli.App
)

func init() {
	cliApp = cli.NewApp()
	cliApp = &cli.App{
		Name:  "run",
		Usage: "Run the application",
		Action: func(*cli.Context) error {
			runApp()
			return nil
		},
	}
}
