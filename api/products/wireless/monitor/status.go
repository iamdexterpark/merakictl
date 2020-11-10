package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetSSIDStatuses(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/wireless/status",
		api.BaseUrl(), serial)
	var datamodel = SSIDStatuses{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}