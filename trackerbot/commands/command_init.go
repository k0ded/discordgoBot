package commands

var Commands = [...]Command{
	{
		Execute:     PingCommand,
		Triggers:    &[]string{"ping", "p"},
		Name:        "Ping",
		Description: "Displays the heartbeat ping",
	},
	{
		Execute:     QuoteCommand,
		Triggers:    &[]string{"quote"},
		Name:        "Quote",
		Description: "Sends a random quote from the one and only",
	},
}
