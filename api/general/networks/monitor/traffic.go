package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type TrafficAnalysis []struct {
	Application string      `json:"application"`
	Destination interface{} `json:"destination"`
	Protocol    string      `json:"protocol"`
	Port        int         `json:"port"`
	Sent        int         `json:"sent"`
	Recv        int         `json:"recv"`
	NumClients  int         `json:"numClients"`
	ActiveTime  int         `json:"activeTime"`
	Flows       int         `json:"flows"`
}

// Return the traffic analysis data for this network
func GetTrafficAnalysis(networkId, t0, timespan, deviceType string) (TrafficAnalysis, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/traffic", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("timespan", timespan)
	parameters.Add("deviceType", deviceType)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = TrafficAnalysis{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}