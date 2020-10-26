package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type TrafficShapingSettings struct {
	TrafficShapingEnabled bool `json:"trafficShapingEnabled"`
	DefaultRulesEnabled   bool `json:"defaultRulesEnabled"`
	Rules                 []struct {
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
		PcpTagValue  interface{} `json:"pcpTagValue"`
	} `json:"rules"`
}

// Display The Traffic Shaping Settings For A SSID On An MR Network
func GetTrafficShapingSettings(networkId, number string) (TrafficShapingSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/trafficShaping/rules",
		api.BaseUrl(), networkId, number)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var ssid TrafficShapingSettings
	user_agent.UnMarshalJSON(session.Body, &ssid)
	traceback := user_agent.TraceBack(session)
	return ssid, traceback
}