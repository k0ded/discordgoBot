package commands

import "fmt"

func PingCommand(ctx *CommandContext)  {
	if ctx.Trigger == "p" {
		ctx.Reply(
			fmt.Sprintf(
				"Pong, %vms",
				ctx.Session.HeartbeatLatency().Milliseconds(),
				),
			)
	}else {
		ctx.Reply("Pong!")
	}
}
