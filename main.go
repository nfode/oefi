package main

import (
	"github.com/nfode/oefi/cmd"
	"github.com/nfode/oefi/pkg/api"
)

func main() {
	client := api.Client{Adress: "https://2.db.transport.rest"}
	cmd.AddSubcommand(cmd.NewCmdSearch(&client))
	cmd.AddSubcommand(cmd.NewCmdCompletion())
	cmd.AddSubcommand(cmd.NewCmdJourney(&client))
	cmd.AddSubcommand(cmd.NewCmdDeparture(&client))
	cmd.Execute()
}
