package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type AggregatedConnectivityNetwork struct {
	Mac             string `json:"mac"`
	ConnectionStats struct {
		Assoc   int `json:"assoc"`
		Auth    int `json:"auth"`
		Dhcp    int `json:"dhcp"`
		DNS     int `json:"dns"`
		Success int `json:"success"`
	} `json:"connectionStats"`
}

// Aggregated Connectivity Info For A Given Client On This Network
func GetAggregatedConnectivity(devices, serial, t0, t1, timespan,
	band, ssid, vlan, apTag string) (AggregatedConnectivityNetwork, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/wireless/clients/%s/connectionStats",
		api.BaseUrl(), devices, serial)
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

	var results = AggregatedConnectivityNetwork{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


type AggregatedConnectivityClients struct {
	Assoc   int `json:"assoc"`
	Auth    int `json:"auth"`
	Dhcp    int `json:"dhcp"`
	DNS     int `json:"dns"`
	Success int `json:"success"`
}

// Aggregated Connectivity Info For This Network Grouped By Clients
func GetAggregatedConnectivityClients(networkId, clientId, t0, t1, timespan,
	band, ssid, vlan, apTag string) (AggregatedConnectivityClients , interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clients/%s/connectionStats",
		api.BaseUrl(), networkId, clientId)
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

	var results = AggregatedConnectivityClients {}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type WirelessConnectivityEvents []struct {
	OccurredAt   int     `json:"occurredAt"`
	DeviceSerial string  `json:"deviceSerial"`
	Band         int     `json:"band"`
	SsidNumber   int     `json:"ssidNumber"`
	Type         string  `json:"type"`
	Subtype      string  `json:"subtype"`
	Severity     string  `json:"severity"`
	DurationMs   float64 `json:"durationMs"`
	Channel      int     `json:"channel"`
	Rssi         int     `json:"rssi"`
	EventData    struct {
		ClientIP string `json:"clientIp"`
	} `json:"eventData"`
}

// List the wireless connectivity events for a client within a network in the timespan
func GetWirelessConnectivityEvents(networkId, clientId, perPage, startingAfter,
	endingBefore, t0, t1, timespan, types, includedSeverities, band, ssidNumber,
	deviceSerial string) (WirelessConnectivityEvents , interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clients/%s/connectivityEvents",
		api.BaseUrl(), networkId, clientId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("perPage", perPage)
	parameters.Add("startingAfter", startingAfter)
	parameters.Add("endingBefore", endingBefore)
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan", timespan)
	parameters.Add("types", types)
	parameters.Add("includedSeverities", includedSeverities)
	parameters.Add("band", band)
	parameters.Add("ssidNumber", ssidNumber)
	parameters.Add("deviceSerial", deviceSerial)

	session.Request.URL.RawQuery = parameters.Encode()

	var results = WirelessConnectivityEvents{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return the latency history for a client
