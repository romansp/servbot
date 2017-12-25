package commandHandlers

import (
	"unicode/utf8"

	"github.com/khades/servbot/ircClient"
	"github.com/khades/servbot/models"
	"github.com/khades/servbot/repos"
)

func Vote(online bool, chatMessage *models.ChatMessage, chatCommand models.ChatCommand, ircClient *ircClient.IrcClient) {
	channelInfo, error := repos.GetChannelInfo(&chatMessage.ChannelID)
	if error != nil || channelInfo.SubdayEnabled == false {
		return
	}

	if utf8.RuneCountInString(chatMessage.MessageBody) < 2 {
		return
	}
	game := chatMessage.MessageBody[1:]
	subday, subdayError := repos.GetLastActiveSubday(&chatMessage.ChannelID)
	if subdayError != nil {
		ircClient.SendPublic(&models.OutgoingMessage{
			Channel: chatMessage.Channel,
			Body:    "Сабдей ещё не запущен SMOrc",
			User:    chatMessage.User})
		return
	}
	if subday.SubsOnly == true && chatMessage.IsSub == false {
		ircClient.SendPublic(&models.OutgoingMessage{
			Channel: chatMessage.Channel,
			Body:    "Ты не саб SMOrc",
			User:    chatMessage.User})
		return
	}
	repos.VoteForSubday(&chatMessage.User, &chatMessage.UserID, &subday.ID, &game)
}
