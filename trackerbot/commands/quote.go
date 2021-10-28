package commands

import (
	"discord/trackerbot/configuration"
	"math/rand"
	"time"
)

var (
	quotes []string
	random *rand.Rand
)

func init() {
	quotes = *configuration.LoadQuotes("Resources/quotes.json")
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func QuoteCommand(ctx *CommandContext) {
	r := random.Intn(len(quotes))
	ctx.Reply(quotes[r])
}
