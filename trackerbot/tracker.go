package main

import (
	"discord/trackerbot/configuration"
	"discord/trackerbot/listeners"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	config *configuration.Config

)

func init()  {
	config = configuration.LoadConfig("Resources/config.json")
}

func main() {
	log.SetPrefix("BOT: ")
	log.SetFlags(log.Lmsgprefix | log.Ltime)

	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal("Error creating session session,", err)
		return
	}

	bot, err := session.GatewayBot()
	if err != nil {
		log.Fatal("Error obtaining recommended amount of shards,", err)
		return
	}

	if config.UseSharding {
		session.ShardID = config.ShardID
		session.ShardCount = bot.Shards
	}

	usr, err := session.User("@me")
	if err != nil {
		log.Fatal("Error obtaining account details,", err)
		return
	}

	listener := listeners.New(config.Prefix)
	session.AddHandler(listener.ListenToCommand)
	session.AddHandler(listener.ListenToReady)

	err = session.Open()
	if err != nil {
		log.Fatal("Error opening connection,", err)
		return
	}

	session.LogLevel = 1

	log.Println("Started under the name: " + usr.Username)

	<-make(chan bool)
}
