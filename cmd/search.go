package cmd

import (
	"fmt"
	"github.com/nfode/oefi/pkg/api"
	"github.com/urfave/cli/v2"
)

func NewCmdSearch(api *api.Client) *cli.Command {
	return &cli.Command{
		Name:    "search",
		Aliases: []string{"s"},
		Action: func(ctx *cli.Context) error {
			result := api.Search(ctx.Args().First())
			for _, result := range result {
				name := sanitizeStationName(result.Name)
				fmt.Println(name)
			}
			return nil
		},
	}
}
