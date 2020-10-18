package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type PerformanceClasses []struct {
	PerformanceClass
}
type PerformanceClass struct {
	CustomPerformanceClassID string `json:"customPerformanceClassId"`
	Name                     string `json:"name"`
	MaxLatency               int    `json:"maxLatency"`
	MaxJitter                int    `json:"maxJitter"`
	MaxLossPercentage        int    `json:"maxLossPercentage"`
}

// List all custom performance classes for an MX network
func GetPerformanceClasses(networkId string ) (PerformanceClasses, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/trafficShaping/customPerformanceClasses", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = PerformanceClasses{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return a custom performance class for an MX network
func GetPerformanceClass(networkId, customPerformanceClassId string ) (PerformanceClass, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/trafficShaping/customPerformanceClasses/%s", api.BaseUrl(), networkId, customPerformanceClassId)
	var payload io.ReadSeeker
	var results = PerformanceClass{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


type TrafficShapingSettings struct {
	DefaultRulesEnabled bool `json:"defaultRulesEnabled"`
	Rules               []struct {
		Definitions []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"definitions"`
		PerClientBandwidthLimits struct {
			Settings        string `json:"settings"`
			BandwidthLimits struct {
				LimitUp   int `json:"limitUp"`
				LimitDown int `json:"limitDown"`
			} `json:"bandwidthLimits"`
		} `json:"perClientBandwidthLimits"`
		DscpTagValue interface{} `json:"dscpTagValue"`
		Priority     string      `json:"priority"`
	} `json:"rules"`
}

// Display the traffic shaping settings rules for an MX network
func GetTrafficShapingSettings(networkId string ) (TrafficShapingSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/trafficShaping/rules", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = TrafficShapingSettings{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
