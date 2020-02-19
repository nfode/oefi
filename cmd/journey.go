package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"

	"github.com/nfode/oefi/pkg/api"
)

func NewCmdJourney(api *api.Client) *cli.Command {
	cmd := &cli.Command{
		Name:         "journey",
		Aliases:      []string{"j"},
		BashComplete: completeJourneyCmd(api),
		Action:       runJourneyCmd(api),
	}
	cmd.Flags = append(cmd.Flags, &cli.StringFlag{
		Name:     "from",
		Aliases:  []string{"f"},
		Required: true,
	})
	cmd.Flags = append(cmd.Flags, &cli.StringFlag{
		Name:     "to",
		Aliases:  []string{"t"},
		Required: true,
	})
	return cmd
}

func runJourneyCmd(api *api.Client) cli.ActionFunc {
	return func(context *cli.Context) error {
		return nil
	}
}

func completeJourneyCmd(api *api.Client) cli.BashCompleteFunc {
	return func(ctx *cli.Context) {
		// This will complete if no args are passed
		if ctx.NArg() > 0 {
			return
		}
		if len(os.Args) > 2 {
			flag := os.Args[len(os.Args)-3]
			if flag == "--from" || flag == "-f" {
				search := os.Args[len(os.Args)-2]
				results := api.Search(search)
				for _, station := range results {
					fmt.Println(sanitizeStationName(station.Name))
				}

			} else {
				cli.DefaultCompleteWithFlags(ctx.App.Command("journey"))(ctx)
			}
		}

	}

}
