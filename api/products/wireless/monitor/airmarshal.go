package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetAirMarshalScanResults(serial, t0, timespan string) (AirMarshalScanResults, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/airMarshal",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("timespan",timespan)

	session.Request.URL.RawQuery = parameters.Encode()
	var results = AirMarshalScanResults{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
