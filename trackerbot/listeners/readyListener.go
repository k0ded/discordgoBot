package listeners

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (listener *Listener) ListenToReady(session *discordgo.Session, ready *discordgo.Ready) {
	_, err := session.UserUpdateStatus(discordgo.StatusOnline)
	if err != nil {
		return 
	}

	log.Println("Bot ready!")
}
