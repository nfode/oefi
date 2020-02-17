package cmd

import (
    "fmt"
    "github.com/nfode/oefi/pkg/api"
    "github.com/spf13/cobra"
    "os"
    "text/tabwriter"
)

func NewCmdDeparture(api *api.Client) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "departures",
        Short: "",
        Long:  ``,
        Run: func(cmd *cobra.Command, args []string) {
            possibleStation := api.Search(args[0])
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
        },
    }
    return cmd
}
