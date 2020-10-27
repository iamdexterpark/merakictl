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
	band, ssid, vlan, apTag string) (AggregatedConnectivityClients, interface{}) {
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

	var results = AggregatedConnectivityClients{}
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
	deviceSerial string) (WirelessConnectivityEvents, interface{}) {
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

type ClientLatencyHistory []struct {
	T0                    int `json:"t0"`
	T1                    int `json:"t1"`
	LatencyBinsByCategory struct {
		BackgroundTraffic struct {
			Zero5   int `json:"0.5"`
			One1    int `json:"1.1"`
			Two1    int `json:"2.1"`
			Four1   int `json:"4.1"`
			Eight1  int `json:"8.1"`
			One61   int `json:"16.1"`
			Three21 int `json:"32.1"`
			Six41   int `json:"64.1"`
			One281  int `json:"128.1"`
			Two561  int `json:"256.1"`
			Five121 int `json:"512.1"`
			One0241 int `json:"1024.1"`
			Two0481 int `json:"2048.1"`
		} `json:"backgroundTraffic"`
		BestEffortTraffic struct {
			Zero5   int `json:"0.5"`
			One1    int `json:"1.1"`
			Two1    int `json:"2.1"`
			Four1   int `json:"4.1"`
			Eight1  int `json:"8.1"`
			One61   int `json:"16.1"`
			Three21 int `json:"32.1"`
			Six41   int `json:"64.1"`
			One281  int `json:"128.1"`
			Two561  int `json:"256.1"`
			Five121 int `json:"512.1"`
			One0241 int `json:"1024.1"`
			Two0481 int `json:"2048.1"`
		} `json:"bestEffortTraffic"`
		VideoTraffic struct {
			Zero5   int `json:"0.5"`
			One1    int `json:"1.1"`
			Two1    int `json:"2.1"`
			Four1   int `json:"4.1"`
			Eight1  int `json:"8.1"`
			One61   int `json:"16.1"`
			Three21 int `json:"32.1"`
			Six41   int `json:"64.1"`
			One281  int `json:"128.1"`
			Two561  int `json:"256.1"`
			Five121 int `json:"512.1"`
			One0241 int `json:"1024.1"`
			Two0481 int `json:"2048.1"`
		} `json:"videoTraffic"`
		VoiceTraffic struct {
			Zero5   int `json:"0.5"`
			One1    int `json:"1.1"`
			Two1    int `json:"2.1"`
			Four1   int `json:"4.1"`
			Eight1  int `json:"8.1"`
			One61   int `json:"16.1"`
			Three21 int `json:"32.1"`
			Six41   int `json:"64.1"`
			One281  int `json:"128.1"`
			Two561  int `json:"256.1"`
			Five121 int `json:"512.1"`
			One0241 int `json:"1024.1"`
			Two0481 int `json:"2048.1"`
		} `json:"voiceTraffic"`
	} `json:"latencyBinsByCategory"`
}

// Return the latency history for a client
func GetClientLatencyHistory(networkId, clientId, t0, t1, timespan,
	resolution string) (ClientLatencyHistory, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clients/%s/latencyHistory",
		api.BaseUrl(), networkId, clientId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan", timespan)
	parameters.Add("resolution", resolution)

	session.Request.URL.RawQuery = parameters.Encode()

	var results = ClientLatencyHistory{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type AggregatedLatencies struct {
	AggregatedLatency
}
type AggregatedLatency struct {
	Mac          string `json:"mac"`
	LatencyStats struct {
		BackgroundTraffic struct {
			RawDistribution struct {
				Num0    int `json:"0"`
				Num1    int `json:"1"`
				Num2    int `json:"2"`
				Num4    int `json:"4"`
				Num8    int `json:"8"`
				Num16   int `json:"16"`
				Num32   int `json:"32"`
				Num64   int `json:"64"`
				Num128  int `json:"128"`
				Num256  int `json:"256"`
				Num512  int `json:"512"`
				Num1024 int `json:"1024"`
				Num2048 int `json:"2048"`
			} `json:"rawDistribution"`
			Avg float64 `json:"avg"`
		} `json:"backgroundTraffic"`
		BestEffortTraffic string `json:"bestEffortTraffic"`
		VideoTraffic      string `json:"videoTraffic"`
		VoiceTraffic      string `json:"voiceTraffic"`
	} `json:"latencyStats"`
}

// Aggregated Latency Info For This Network Grouped By Clients
func GetAggregatedLatencies(networkId, t0, t1, timespan,
	band, ssid, vlan, apTag, fields string) (AggregatedLatencies, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clients/latencyStats",
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
	parameters.Add("fields", fields)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = AggregatedLatencies{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Aggregated Latency Info For A Given Client On This Network
func GetAggregatedLatency(networkId, clientId, t0, t1, timespan,
	band, ssid, vlan, apTag, fields string) (AggregatedLatency, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clients/%s/latencyStats",
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
	parameters.Add("fields", fields)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = AggregatedLatency{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

