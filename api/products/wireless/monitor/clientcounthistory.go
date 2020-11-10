package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type WirelessClientCount []struct {
	StartTs     time.Time `json:"startTs"`
	EndTs       time.Time `json:"endTs"`
	ClientCount int       `json:"clientCount"`
}

// Return wireless client counts over time for a network, device, or network client
func GetWirelessClientCount(networkId, t0, t1, timespan,
	resolution, autoResolution, clientId, deviceSerial, apTag,
	band, ssid string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clientCountHistory",
		api.BaseUrl(), networkId)
	var datamodel = WirelessClientCount{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":             t0,
		"t1":             t1,
		"timespan":       timespan,
		"resolution":     resolution,
		"autoResolution": autoResolution,
		"clientId":       clientId,
		"deviceSerial":   deviceSerial,
		"apTag":          apTag,
		"band":           band,
		"ssid":           ssid}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
