package cmd

import (
	"fmt"
	"github.com/nfode/oefi/internal/stationcache"
	"github.com/nfode/oefi/pkg/api"
	"github.com/urfave/cli/v2"
	"os"
	"text/tabwriter"
)

func NewCmdDeparture(api *api.Client) *cli.Command {
	return &cli.Command{
		Name:         "departures",
		BashComplete: completeDeparturesCmd(api),
		Action:       run(api),
	}
}

func run(api *api.Client) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		station := ctx.Args().First()
		possibleStation := api.Search(station)
		if len(possibleStation) > 1 {
			panic("should not happen")
		}
		result := api.Departures(possibleStation[0].Id)
		// initialize tabwriter
		w := new(tabwriter.Writer)

		// minwidth, tabwidth, padding, padchar, flags
		w.Init(os.Stdout, 8, 8, 0, '\t', 0)
		fmt.Fprintf(w, "%s\t%s\t%s\t", "Line", "Departure", "To")
		fmt.Fprintf(w, "\n %s\t%s\t%s\t", "----", "----", "----")
		for _, departures := range result {
			fmt.Fprintf(w, "\n %v\t%v\t%v\t", departures.Line.Name, departures.When, departures.Direction)
		}
		w.Flush()
		fmt.Print("\n")
		return nil
	}
}

func completeDeparturesCmd(api *api.Client) cli.BashCompleteFunc {
	return func(ctx *cli.Context) {
		search := os.Args[len(os.Args)-2]
		if search != "" {
			possibleStation := api.Search(search)
			cachedStations := stationcache.Find(search)
			stations := append([]string{}, cachedStations...)
			for _, station := range possibleStation {
				name := sanitizeStationName(station.Name)
				if !contains(stations, name) {
					stations = append(stations, name)
				}
			}
			for _, station := range stations {
				fmt.Fprintf(ctx.App.Writer, fmt.Sprintf("%s:\n", station))
			}
			var result []string
			for _, station := range possibleStation {
				name := sanitizeStationName(station.Name)
				if ! contains(cachedStations, name) {
					result = append(result, name)
				}
			}
			stationcache.Write(result)
		}
	}
}

func contains(arr []string, s string) bool {
	for _, item := range arr {
		if item == s {
			return true
		}
	}
	return false
}
