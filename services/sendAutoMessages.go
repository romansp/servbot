package services

import (
	"strings"

	"github.com/khades/servbot/bot"
	"github.com/khades/servbot/models"
	"github.com/khades/servbot/repos"
)

func SendAutoMessages() {
	messages, error := repos.GetCurrentAutoMessages()
	if error != nil {
		return
	}
	for _, message := range *messages {
		repos.ResetAutoMessageThreshold(&message)
		channel, error := repos.GetUsernameByID(&message.ChannelID)
		if error == nil && strings.TrimSpace(message.Message) != "" {
			bot.IrcClientInstance.SendPublic(&models.OutgoingMessage{Channel: *channel, Body: message.Message})
		}
	}
}
