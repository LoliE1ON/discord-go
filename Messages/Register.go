package Messages

import (
	"github.com/LoliE1ON/discord-go/Messages/Message"
	"github.com/bwmarrin/discordgo"
)

// Register messages
func Register(dg *discordgo.Session) {
	dg.AddHandler(Message.About)
}
