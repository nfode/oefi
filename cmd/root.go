package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var completion = `
__oefi_search_station()
{
    local oefi_out
    if oefi_out=$(__oefi_debug_out "oefi search \""${COMP_WORDS[COMP_CWORD]}"\""); then
		__oefi_debug "${oefi_out[*]}"
        COMPREPLY=( $( compgen -W "${oefi_out[*]}" -- "${COMP_WORDS[COMP_CWORD]}" ) )
    fi
}

__custom_func() {
	__oefi_debug "last command ${last_command}"
    case ${last_command} in
        oefi_departures)
            __oefi_search_station
            return
            ;;
        *)
            ;;
    esac
}
`

type RootCmd struct {
	root *cobra.Command
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = RootCmd{
	root: newCmdRoot(),
}

func newCmdRoot() *cobra.Command {
	return &cobra.Command{
		Use:                    "oefi",
		Short:                  "A brief description of your application",
		Long:                   "",
		BashCompletionFunction: completion,
	}
}

func AddSubcommand(cmd *cobra.Command) {
	rootCmd.root.AddCommand(cmd)
}

func AddFlagCompletion(cmd *cobra.Command, name string, complete string) {
	if cmd.Flag(name) != nil {
		if cmd.Flag(name).Annotations == nil {
			cmd.Flag(name).Annotations = map[string][]string{}
		}
		cmd.Flag(name).Annotations[cobra.BashCompCustom] = append(
			cmd.Flag(name).Annotations[cobra.BashCompCustom],
			complete,
		)
	}
}

func Execute() {
	if err := rootCmd.root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
