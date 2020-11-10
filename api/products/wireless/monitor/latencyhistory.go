package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type LatencyHistory []struct {
	StartTs      time.Time `json:"startTs"`
	EndTs        time.Time `json:"endTs"`
	AvgLatencyMs int       `json:"avgLatencyMs"`
}

// Return average wireless latency over time for a network, device, or network client
func GetLatencyHistory(networkId, t0, t1, timespan,
	resolution, autoResolution, clientId, deviceSerial, apTag, band, ssid, accessCategory string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/latencyHistory",
		api.BaseUrl(), networkId)
	var datamodel = LatencyHistory{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"resolution": resolution,
		"autoResolution": autoResolution,
		"clientId": clientId,
		"deviceSerial": deviceSerial,
		"apTag": apTag,
		"band": band,
		"ssid": ssid,
		"accessCategory": accessCategory}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}