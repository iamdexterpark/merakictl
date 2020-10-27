package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type AggregatedConnectivity []struct {
	Serial          string `json:"serial"`
	ConnectionStats struct {
		Assoc   int `json:"assoc"`
		Auth    int `json:"auth"`
		Dhcp    int `json:"dhcp"`
		DNS     int `json:"dns"`
		Success int `json:"success"`
	} `json:"connectionStats"`
}


// Aggregated connectivity info for this network, grouped by node
func GetAggregatedConnectivity(networkId, t0, t1, timespan,
	band, ssid, vlan, apTag string) (AggregatedConnectivity, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/connectionStats",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan", timespan)
	parameters.Add("band", band)
	parameters.Add("ssid", ssid)
	parameters.Add("vlan", vlan)
	parameters.Add("apTag", apTag)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = AggregatedConnectivity{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
// Aggregated latency info for this network, grouped by node
