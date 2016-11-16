package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/khades/servbot/models"
	"github.com/khades/servbot/repos"
)

type tdjTrack struct {
	Title string
}

// CheckTwitchDJTrack checks last playing track
func CheckTwitchDJTrack() {
	channels := repos.GetTwitchDJEnabledChannels()
	log.Println(len(channels))
	for _, channel := range channels {
		checkOneTwitchDJTrack(channel)
	}
}

func checkOneTwitchDJTrack(channel *models.ChannelInfo) {
	status := models.TwitchDJ{ID: channel.TwitchDJ.ID}
	defer repos.PushTwitchDJ(&channel.Channel, &status)
	resp, error := http.Get("https://twitch-dj.ru/includes/back.php?func=get_track&channel=" + channel.TwitchDJ.ID)
	defer resp.Body.Close()

	if error != nil {
		log.Println(error)
		return
	}
	track := tdjTrack{}
	marshallError := json.NewDecoder(resp.Body).Decode(&track)
	if marshallError != nil {
		log.Println(marshallError)
		return
	}
	if track.Title != "" {
		status.Playing = true
		status.Track = track.Title
	}
}