package cmd

import (
    "fmt"

    "github.com/nfode/oefi/pkg/api"
    "github.com/spf13/cobra"
)

func NewCmdJourney(api *api.Client) *cobra.Command {
    var from string
    var to string
    cmd := &cobra.Command{
        Use:   "journey",
        Short: "",
        Long:  ``,
        Run: func(cmd *cobra.Command, args []string) {
           fmt.Println(from) 
        },
    }
    cmd.Flags().StringVar(&from, "from", from, "from")
    AddFlagCompletion(cmd,"from","__oefi_search_station")
    cmd.Flags().StringVar(&to, "to", to, "to")
    AddFlagCompletion(cmd,"to","__oefi_search_station")
    return cmd
}
