package cmd

import (
    "fmt"
    "strings"
    "unicode"

    "github.com/nfode/oefi/pkg/api"
    "github.com/spf13/cobra"
)

func NewCmdSearch(api *api.Client) *cobra.Command {
    return &cobra.Command{
        Use:   "search",
        Short: "",
        Long:  ``,
        Run: func(cmd *cobra.Command, args []string) {
            result := api.Search(args[0])
            for _, result := range result {
                name := strings.Map(func(r rune) rune {
                    if unicode.IsSpace(r) {
                        return '-'
                    }
                    return r
                },result.Name)
                fmt.Println(name)
            }
        },
        Args: cobra.MaximumNArgs(1),
    }
}
