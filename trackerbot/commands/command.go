package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type Command struct {
	Execute  	func (ctx *CommandContext)
	Triggers 	*[]string
	Name     	string
	Description string
}

type CommandContext struct {
	Trigger string
	Args 	[]string
	Author 	*discordgo.User
	Session *discordgo.Session
	Event 	*discordgo.MessageCreate
}

func (ctx *CommandContext) Reply(message string) *discordgo.Message {
	msg, err := ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, message)
	if err != nil {
		log.Fatal("Error sending message,", err)
		return nil
	}
	return msg
}

func CreateContext(trigger string, args []string, session *discordgo.Session, event *discordgo.MessageCreate) *CommandContext {
	return &CommandContext{
		Trigger: trigger,
		Args: args,
		Author: event.Author,
		Session: session,
		Event: event,
	}
}
