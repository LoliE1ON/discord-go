package RichEmbed

import (
	"github.com/bwmarrin/discordgo"
)

var Help discordgo.MessageEmbed = discordgo.MessageEmbed{
	Title: "Commands for color selection:",
	Color: 16711680,
	Fields: []*discordgo.MessageEmbedField{{
		Name:   "Color-hex gives information about colors",
		Value:  "```!color #d81e4c``` ```!color #00ff15```",
		Inline: true,
	}},
}
