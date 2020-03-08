package RichEmbed

import (
	"github.com/LoliE1ON/discord-go/Helpers/HexHelper"
	"github.com/bwmarrin/discordgo"
)

var ColorIncorrect discordgo.MessageEmbed = discordgo.MessageEmbed{
	Title: "Color incorrect",
	Fields: []*discordgo.MessageEmbedField{{
		Name:   "Color-HEX gives information about colors",
		Value:  "You entered the wrong color. Need to enter Color-HEX",
		Inline: true,
	}},
}

func ColorAssign(color string) (message discordgo.MessageEmbed, err error) {

	hexInt, err := HexHelper.HexToInt(color)
	if err != nil {
		return
	}

	message = discordgo.MessageEmbed{
		Title: "Color successfully assigned",
		Color: hexInt,
	}

	return
}
