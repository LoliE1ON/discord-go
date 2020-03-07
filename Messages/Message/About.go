package Message

import "github.com/bwmarrin/discordgo"

func About(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!miku" {
		s.ChannelMessageSend(m.ChannelID, "Baka!")
	}

}
