package Message

import (
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/LoliE1ON/discord-go/Helpers/HexHelper"

	"github.com/bwmarrin/discordgo"
)

var (
	color   string
	session *discordgo.Session
	message *discordgo.MessageCreate
)

func Color(s *discordgo.Session, m *discordgo.MessageCreate) {
	session, message = s, m

	var args []string = strings.Split(message.Content, " ")
	if args[0] == "!colorTest" && len(args) > 1 {

		color = args[1]

		// Validate hex value
		_, err := HexHelper.ParseHexColor(color)
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Invalid hex value")
			return
		}

		// Get role id
		roleId, err := getRole()
		if err != nil {
			log.Println(err)
			return
		}

		// Remove old roles
		err = removeRoles()
		if err != nil {
			log.Println(err)
			return
		}

		// Assign role
		session.GuildMemberRoleAdd(message.GuildID, message.Author.ID, roleId)

		session.ChannelMessageSend(message.ChannelID, "Role assign")
	}

}

// Create new role or return exist role
func getRole() (roleId string, err error) {

	// Fetching roles
	roles, err := session.GuildRoles(message.GuildID)
	if err != nil {
		err = errors.Wrap(err, "Error fetching roles")
		return
	}

	// Return old role
	for _, role := range roles {
		if role.Name == color {
			roleId = role.ID
			return
		}
	}

	// Create new role
	role, err := session.GuildRoleCreate(message.GuildID)
	if err != nil {
		err = errors.Wrap(err, "Error create new role")
		return
	}

	// Convert hex to int
	n, err := strconv.ParseInt(strings.Replace(color, "#", "", -1), 16, 32)
	if err != nil {
		err = errors.Wrap(err, "Error convert ")
		return
	}
	hexInt := int(n)

	// Edit new role
	newRole, err := session.GuildRoleEdit(message.GuildID, role.ID, color, hexInt, false, role.Permissions, false)
	if err != nil {
		err = errors.Wrap(err, "Error update role")
		return
	}

	roleId = newRole.ID
	return
}

// Remove old roles user
func removeRoles() (err error) {

	// Fetching server roles
	serverRoles, err := session.GuildRoles(message.GuildID)
	if err != nil {
		err = errors.Wrap(err, "Error fetching server roles")
		return
	}

	// Fetching user roles
	roles := message.Member.Roles

	// Govnokod
	for _, serverRole := range serverRoles {
		_, err := HexHelper.ParseHexColor(serverRole.Name)
		if err == nil && contains(roles, serverRole.ID) == true {
			session.GuildMemberRoleRemove(message.GuildID, message.Author.ID, serverRole.ID)
		}
	}

	return
}

// TODO: Put in helper
func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
