package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type ConnectivityTesting struct {
	Destinations []struct {
		IP          string `json:"ip"`
		Description string `json:"description"`
		Default     bool   `json:"default"`
	} `json:"destinations"`
}

// Return the connectivity testing destinations for an MG network
func GetConnectivityTesting(networkId string) (ConnectivityTesting, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/connectivityMonitoringDestinations",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = ConnectivityTesting{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}