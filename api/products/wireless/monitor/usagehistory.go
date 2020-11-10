package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type APUsageHistory []struct {
	StartTs      time.Time `json:"startTs"`
	EndTs        time.Time `json:"endTs"`
	TotalKbps    int       `json:"totalKbps"`
	SentKbps     int       `json:"sentKbps"`
	ReceivedKbps int       `json:"receivedKbps"`
}

// Return AP usage over time for a device or network client
func GetAPUsageHistory(networkId, t0, t1, timespan,
	resolution, autoResolution, clientId, deviceSerial, apTag,
	band, ssid string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/usageHistory", api.BaseUrl(), networkId)
	var datamodel = APUsageHistory{}

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
