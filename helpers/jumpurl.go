package helpers

import (
	"github.com/bwmarrin/discordgo"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/config"
)

// JumpURL returns the URL to jump to a message
func JumpURL(m *discordgo.Message) string {
	return "https://discordapp.com/channels/" + config.GuildID + "/" + m.ChannelID + "/" + m.ID
}
