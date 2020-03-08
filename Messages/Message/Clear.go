package Message

import (
	"log"

	"github.com/LoliE1ON/discord-go/Messages/RichEmbed"

	"github.com/bwmarrin/discordgo"
)

// Clear role color
func Clear(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Content == "!clear" {

		// Remove roles
		err := RemoveUserRoles(s, m)
		if err != nil {
			log.Println("Clear role", err)
			return
		}

		// Reply
		_, err = s.ChannelMessageSendEmbed(m.ChannelID, &RichEmbed.Clear)
		if err != nil {
			log.Println("Clear role", err)
			return
		}
	}

}
