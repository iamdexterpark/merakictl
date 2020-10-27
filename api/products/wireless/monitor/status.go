package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type SSIDStatuses struct {
	BasicServiceSets []struct {
		SsidName     string `json:"ssidName"`
		SsidNumber   int    `json:"ssidNumber"`
		Enabled      bool   `json:"enabled"`
		Band         string `json:"band"`
		Bssid        string `json:"bssid"`
		Channel      int    `json:"channel"`
		ChannelWidth string `json:"channelWidth"`
		Power        string `json:"power"`
		Visible      bool   `json:"visible"`
		Broadcasting bool   `json:"broadcasting"`
	} `json:"basicServiceSets"`
}

// Return the SSID statuses of an access point
func GetSSIDStatuses(serial string) (SSIDStatuses, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/wireless/status",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SSIDStatuses{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}