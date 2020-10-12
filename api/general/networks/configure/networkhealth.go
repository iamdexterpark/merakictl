package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)
type ChannelUtilizations []struct {
	ChannelUtilization
}
type ChannelUtilization struct {
	Serial string `json:"serial"`
	Model  string `json:"model"`
	Tags   string `json:"tags"`
	Wifi0  []struct {
		StartTime           time.Time `json:"startTime"`
		EndTime             time.Time `json:"endTime"`
		UtilizationTotal    float64   `json:"utilizationTotal"`
		Utilization80211    int       `json:"utilization80211"`
		UtilizationNon80211 float64   `json:"utilizationNon80211"`
	} `json:"wifi0"`
	Wifi1 []struct {
		StartTime           time.Time `json:"startTime"`
		EndTime             time.Time `json:"endTime"`
		UtilizationTotal    float64   `json:"utilizationTotal"`
		Utilization80211    int       `json:"utilization80211"`
		UtilizationNon80211 float64   `json:"utilizationNon80211"`
	} `json:"wifi1"`
}

// Get the channel utilization over each radio for all APs in a network.
func GetChannelUtilization(networkId, t0, t1, timespan, resolution, perPage, startingAfter, endingBefore string) (ChannelUtilizations, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/networkHealth/channelUtilization", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan", timespan)
	parameters.Add("resolution", resolution)
	parameters.Add("perPage",perPage)
	parameters.Add("startingAfter",startingAfter)
	parameters.Add("endingBefore", endingBefore)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = ChannelUtilizations{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

