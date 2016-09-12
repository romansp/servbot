package models

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// ChannelInfo describes data of
type ChannelInfo struct {
	Channel      string
	StreamStatus StreamStatus
	Mods         []string
	Commands     []string
}

// GetCommands Helper Command for mustashe
func (channelInfo ChannelInfo) GetCommands() string {
	return strings.Join(channelInfo.Commands, ", ")
}

// GetStreamDuration Helper Command for time for mustashe
func (channelInfo ChannelInfo) GetStreamDuration() string {

	if !channelInfo.StreamStatus.Online {
		return ""
	}
	minutePrefix := "минут"
	hourPrefix := "часов"
	duration := time.Now().Sub(channelInfo.StreamStatus.Start)
	minutes := float64(int(duration.Minutes() - math.Floor(duration.Minutes()/60)*60))
	hours := float64(int(duration.Hours()))
	if math.Floor(minutes/10) != 1 {
		switch int(minutes - math.Floor(minutes/10)*10) {
		case 1:
			minutePrefix = "минуту"
			break
		case 2:
		case 3:
		case 4:
			minutePrefix = "минуты"
		}
	}

	if math.Floor(hours/10) != 1 {
		switch int(hours - math.Floor(hours/10)*10) {
		case 1:
			hourPrefix = "час"
			break
		case 2:
		case 3:
		case 4:
			hourPrefix = "часа"
		}
	}
	if int(minutes) == 0 {
		return fmt.Sprintf("%d %s", int(hours), hourPrefix)

	}
	if int(hours) == 0 {
		return fmt.Sprintf("%d %s", int(minutes), minutePrefix)
	}
	return fmt.Sprintf("%d %s %d %s", int(hours), hourPrefix, int(minutes), minutePrefix)

}
