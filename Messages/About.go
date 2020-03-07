package Messages

import "github.com/bwmarrin/discordgo"

func About(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "!miku" {
		s.ChannelMessageSend(m.ChannelID, "Baka!")
	}

}
