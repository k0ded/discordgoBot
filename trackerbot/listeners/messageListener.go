package listeners

import (
	"discord/trackerbot/commands"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (listener *Listener) ListenToCommand(session *discordgo.Session, event *discordgo.MessageCreate)  {
	msg := &event.Content

	if event.Author.Bot || !strings.HasPrefix(strings.ToLower(*msg), strings.ToLower(listener.prefix)) {
		return
	}
	trimmed := strings.TrimPrefix(event.Content[len(listener.prefix):]," ")
	split := strings.Split(trimmed, " ")
	trig := &split[0]

	var args []string
	if len(split) > 1 {
		args = split[1:]
	}

	for _, cmd := range commands.Commands {
		for _, trigger := range *cmd.Triggers {

			if strings.HasPrefix(strings.ToLower(*trig), strings.ToLower(trigger)) {
				cmd.Execute(
					commands.CreateContext(
						*trig, args, session, event,
						),
					)
				return
			}
		}
	}
}
