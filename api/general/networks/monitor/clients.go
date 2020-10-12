package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type ClientNetworkTraffic []struct {
	Ts            time.Time `json:"ts"`
	Application   string    `json:"application"`
	Destination   string    `json:"destination"`
	Protocol      string    `json:"protocol"`
	Port          int       `json:"port"`
	Recv          int       `json:"recv"`
	Sent          int       `json:"sent"`
	NumFlows      int       `json:"numFlows"`
	ActiveSeconds int       `json:"activeSeconds"`
}

// Return the client's network traffic data over time
func GetClientNetworkTraffic(networkId, clientId string) (ClientNetworkTraffic, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s/trafficHistory", api.BaseUrl(), networkId, clientId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = ClientNetworkTraffic{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type ClientUsageHistory []struct {
	Sent     int       `json:"sent"`
	Received int       `json:"received"`
	Ts       time.Time `json:"ts"`
}

// Return the client's daily usage history
func GetClientUsageHistory(networkId, clientId string) (ClientUsageHistory, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s/usageHistory", api.BaseUrl(), networkId, clientId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = ClientUsageHistory{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}