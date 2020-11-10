package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

// TrafficAnalysis - Return The Traffic Analysis Settings For A Network
type TrafficAnalysis struct {
	Mode                string `json:"mode"`
	CustomPieChartItems []struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"customPieChartItems"`
}

// TrafficShaping - Returns the application categories for traffic shaping rules.
type TrafficShaping struct {
	ApplicationCategories []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Applications []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"applications"`
	} `json:"applicationCategories"`
}

// TrafficShapingDSCP - Returns the available DSCP tagging options for your traffic shaping rules.
type TrafficShapingDSCP []struct {
	DscpTagValue int    `json:"dscpTagValue"`
	Description  string `json:"description"`
}

// GetTrafficAnalysis - Return The Traffic Analysis Settings For A Network
func GetTrafficAnalysis(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/trafficAnalysis", api.BaseUrl(), networkId)
	var datamodel = TrafficAnalysis{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// GetTrafficShaping - Returns the application categories for traffic shaping rules.
func GetTrafficShaping(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/trafficShaping/applicationCategories", api.BaseUrl(), networkId)
	var datamodel = TrafficShaping{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// GetTrafficShapingDSCP - Returns the available DSCP tagging options for your traffic shaping rules.
func GetTrafficShapingDSCP(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/trafficShaping/dscpTaggingOptions", api.BaseUrl(), networkId)
	var datamodel = TrafficShapingDSCP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
