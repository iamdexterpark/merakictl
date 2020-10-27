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

type AggregatedLatencyInfo []struct {
	Serial       string `json:"serial"`
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

// Aggregated latency info for this network, grouped by node
func GetAggregatedLatencyInfo(networkId, t0, t1, timespan,
	band, ssid, vlan, apTag, fields string) (AggregatedLatencyInfo, interface{}) {
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

	var results = AggregatedLatencyInfo{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}