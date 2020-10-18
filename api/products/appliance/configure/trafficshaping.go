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


type TrafficShapingRules struct {
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
func GetTrafficShapingRules(networkId string ) (TrafficShapingRules, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/trafficShaping/rules", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = TrafficShapingRules{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type UplinkBandwidthSettings struct {
	BandwidthLimits struct {
		Wan1 struct {
			LimitUp   int `json:"limitUp"`
			LimitDown int `json:"limitDown"`
		} `json:"wan1"`
		Wan2 struct {
			LimitUp   int `json:"limitUp"`
			LimitDown int `json:"limitDown"`
		} `json:"wan2"`
		Cellular struct {
			LimitUp   int `json:"limitUp"`
			LimitDown int `json:"limitDown"`
		} `json:"cellular"`
	} `json:"bandwidthLimits"`
}

// Returns the uplink bandwidth settings for your MX network.
func GetUplinkBandwidthSettings(networkId string ) (UplinkBandwidthSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/trafficShaping/uplinkBandwidth", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = UplinkBandwidthSettings{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


type UplinkSelectionSettings struct {
	ActiveActiveAutoVpnEnabled  bool   `json:"activeActiveAutoVpnEnabled"`
	DefaultUplink               string `json:"defaultUplink"`
	LoadBalancingEnabled        bool   `json:"loadBalancingEnabled"`
	WanTrafficUplinkPreferences []struct {
		TrafficFilters []struct {
			Type  string `json:"type"`
			Value struct {
				Protocol string `json:"protocol"`
				Source   struct {
					Port string `json:"port"`
					Cidr string `json:"cidr"`
				} `json:"source"`
				Destination struct {
					Port string `json:"port"`
					Cidr string `json:"cidr"`
				} `json:"destination"`
			} `json:"value"`
		} `json:"trafficFilters"`
		PreferredUplink string `json:"preferredUplink"`
	} `json:"wanTrafficUplinkPreferences"`
	VpnTrafficUplinkPreferences []struct {
		TrafficFilters []struct {
			Type  string `json:"type"`
			Value struct {
				ID string `json:"id"`
			} `json:"value"`
		} `json:"trafficFilters"`
		PreferredUplink   string `json:"preferredUplink"`
		FailOverCriterion string `json:"failOverCriterion,omitempty"`
		PerformanceClass  struct {
			Type                     string `json:"type"`
			CustomPerformanceClassID string `json:"customPerformanceClassId"`
		} `json:"performanceClass,omitempty"`
	} `json:"vpnTrafficUplinkPreferences"`
}

// Show uplink selection settings for an MX network
func GetUplinkSelectionSettings(networkId string ) (UplinkSelectionSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/trafficShaping/uplinkSelection", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = UplinkSelectionSettings{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type TrafficShapingSettings struct {
	GlobalBandwidthLimits struct {
		LimitUp   int `json:"limitUp"`
		LimitDown int `json:"limitDown"`
	} `json:"globalBandwidthLimits"`
}

// Display the traffic shaping settings for an MX network
func GetTrafficShapingSettings(networkId string ) (TrafficShapingSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/trafficShaping", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = TrafficShapingSettings{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}