package Message

import (
	"log"

	"github.com/LoliE1ON/discord-go/Messages/RichEmbed"
	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Content == "!help" {
		_, err := s.ChannelMessageSendEmbed(message.ChannelID, &RichEmbed.Help)
		if err != nil {
			log.Println(err)
		}
	}

}
