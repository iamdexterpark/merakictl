package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type RadioSettings struct {
	Serial             string `json:"serial"`
	RfProfileID        string `json:"rfProfileId"`
	TwoFourGhzSettings struct {
		Channel     int `json:"channel"`
		TargetPower int `json:"targetPower"`
	} `json:"twoFourGhzSettings"`
	FiveGhzSettings struct {
		Channel      int `json:"channel"`
		ChannelWidth int `json:"channelWidth"`
		TargetPower  int `json:"targetPower"`
	} `json:"fiveGhzSettings"`
}

// Return The Radio Settings Of A Device
func GetRadioSettings(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/wireless/radio/settings",
		api.BaseUrl(), serial)
	var datamodel = RadioSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
