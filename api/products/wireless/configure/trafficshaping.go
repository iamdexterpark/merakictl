package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetTrafficShapingSettings(networkId, number string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/trafficShaping/rules",
		api.BaseUrl(), networkId, number)
	var datamodel TrafficShapingSettings
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}