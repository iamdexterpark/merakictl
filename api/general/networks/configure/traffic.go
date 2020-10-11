package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type TrafficAnalysis struct {
	Mode                string `json:"mode"`
	CustomPieChartItems []struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"customPieChartItems"`
}

// Return The Traffic Analysis Settings For A Network
func GetTrafficAnalysis(networkId string) (TrafficAnalysis, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/trafficAnalysis", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = TrafficAnalysis{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

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

// Returns the application categories for traffic shaping rules.
func GetTrafficShaping(networkId string) (TrafficShaping, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/trafficShaping/applicationCategories", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = TrafficShaping{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}