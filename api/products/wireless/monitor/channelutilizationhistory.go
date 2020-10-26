package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
	band string) (APChannelUtilization, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/channelUtilizationHistory", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan", timespan)
	parameters.Add("resolution", resolution)
	parameters.Add("autoResolution", autoResolution)
	parameters.Add("clientId", clientId)
	parameters.Add("deviceSerial", deviceSerial)
	parameters.Add("apTag", apTag)
	parameters.Add("band", band)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = APChannelUtilization{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}