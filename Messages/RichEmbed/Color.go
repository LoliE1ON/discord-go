package Messages

import "github.com/bwmarrin/discordgo"

var ColorIncorrect = discordgo.MessageEmbed{
	Title:       "",
	Color:       0,
	Fields:      nil,
}

.setColor('#ff0000')
.setTitle('Color incorrect')
.addField('Color-HEX gives information about colors', 'You entered the wrong color. Need to enter Color-HEX', true)