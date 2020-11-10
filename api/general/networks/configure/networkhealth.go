package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetChannelUtilization(networkId, t0, t1, timespan, resolution, perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/networkHealth/channelUtilization", api.BaseUrl(), networkId)
	var datamodel = ChannelUtilizations{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"resolution": resolution,
		"perPage": perPage,
		"startingAfter": startingAfter,
		"endingBefore": endingBefore}


	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions

}

