package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetRadioSettings(serial string) (RadioSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/wireless/radio/settings",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = RadioSettings{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}