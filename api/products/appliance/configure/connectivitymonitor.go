package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type ConnectivityMonitoringDestinations struct {
	Destinations []struct {
		IP          string `json:"ip"`
		Description string `json:"description"`
		Default     bool   `json:"default"`
	} `json:"destinations"`
}

// Return the connectivity testing destinations for an MX network
func GetConnectivityMonitoringDestinations(networkId string ) (ConnectivityMonitoringDestinations, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/connectivityMonitoringDestinations", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = ConnectivityMonitoringDestinations{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
