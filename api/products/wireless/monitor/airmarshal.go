package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type AirMarshalScanResults []struct {
	Ssid   string `json:"ssid"`
	Bssids []struct {
		Bssid      string `json:"bssid"`
		Contained  bool   `json:"contained"`
		DetectedBy []struct {
			Device string `json:"device"`
			Rssi   int    `json:"rssi"`
		} `json:"detectedBy"`
	} `json:"bssids"`
	Channels      []int    `json:"channels"`
	FirstSeen     int      `json:"firstSeen"`
	LastSeen      int      `json:"lastSeen"`
	WiredMacs     []string `json:"wiredMacs"`
	WiredVlans    []int    `json:"wiredVlans"`
	WiredLastSeen int      `json:"wiredLastSeen"`
}

// List Air Marshal scan results from a network
func GetAirMarshalScanResults(serial, t0, timespan string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/airMarshal",
		api.BaseUrl(), serial)
	var datamodel = AirMarshalScanResults{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"timespan": timespan}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
