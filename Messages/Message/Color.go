package Message

import (
	"log"
	"strconv"
	"strings"

	"github.com/LoliE1ON/discord-go/Helpers/SliceHelper"

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

		// Remove old roles user
		err = removeUserRoles()
		if err != nil {
			log.Println(err)
			return
		}

		// Assign role
		err = session.GuildMemberRoleAdd(message.GuildID, message.Author.ID, roleId)
		if err != nil {
			log.Println(err)
			return
		}

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
func removeUserRoles() (err error) {

	// Fetching server roles
	serverRoles, err := session.GuildRoles(message.GuildID)
	if err != nil {
		err = errors.Wrap(err, "Error fetching server roles")
		return
	}

	// Fetching user roles
	userRoles := message.Member.Roles

	// Search role and remove role
	for _, serverRole := range serverRoles {

		_, err := HexHelper.ParseHexColor(serverRole.Name)
		if err == nil && SliceHelper.Contains(userRoles, serverRole.ID) {
			errRemoveRole := session.GuildMemberRoleRemove(message.GuildID, message.Author.ID, serverRole.ID)
			if errRemoveRole != nil {
				return errors.Wrap(errRemoveRole, "Error remove role")
			}
			log.Println("Remove role", serverRole.Name)
		}
	}

	return
}
