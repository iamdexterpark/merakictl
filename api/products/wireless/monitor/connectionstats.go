package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type AggregatedConnectivityPerAP struct {
	Serial          string `json:"serial"`
	ConnectionStats struct {
		Assoc   int `json:"assoc"`
		Auth    int `json:"auth"`
		Dhcp    int `json:"dhcp"`
		DNS     int `json:"dns"`
		Success int `json:"success"`
	} `json:"connectionStats"`
}

// Aggregated Connectivity Info For A Given AP On This Network
func GetAPAggregatedConnectivity(serial, t0, t1, timespan,
	band, ssid, vlan, apTag string) (AggregatedConnectivityPerAP, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/wireless/connectionStats",
		api.BaseUrl(), serial)
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

	var results = AggregatedConnectivityPerAP{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type AggregatedConnectivityPerNetwork struct {
	Assoc   int `json:"assoc"`
	Auth    int `json:"auth"`
	Dhcp    int `json:"dhcp"`
	DNS     int `json:"dns"`
	Success int `json:"success"`
}

// Aggregated Connectivity Info For This Network
func GetAggregatedConnectivityPerNetwork(networkId, t0, t1, timespan,
	band, ssid, vlan, apTag string) (AggregatedConnectivityPerNetwork, interface{}) {
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

	var results = AggregatedConnectivityPerNetwork{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
// Return PHY data rates over time for a network, device, or network client
