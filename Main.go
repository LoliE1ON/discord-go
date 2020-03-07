package main

import (
	"fmt"
	"log"

	"github.com/LoliE1ON/discord-go/Helpers/ConfigHelper"

	"github.com/bwmarrin/discordgo"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	// Get config file
	config, err := ConfigHelper.Get()
	if err != nil {
		log.Println(err)
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	//dg.AddHandler()

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

}
