package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type APChannelUtilization []struct {
	StartTs             time.Time `json:"startTs"`
	EndTs               time.Time `json:"endTs"`
	UtilizationTotal    float64   `json:"utilizationTotal"`
	Utilization80211    float64   `json:"utilization80211"`
	UtilizationNon80211 float64   `json:"utilizationNon80211"`
}

// Return AP channel utilization over time for a device or network client
func GetAPChannelUtilization(networkId, t0, t1, timespan,
	resolution, autoResolution, clientId, deviceSerial, apTag,
	band string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/channelUtilizationHistory", api.BaseUrl(), networkId)
	var datamodel = APChannelUtilization{}

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
		"band": band}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}