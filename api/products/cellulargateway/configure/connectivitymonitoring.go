package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type ConnectivityTesting struct {
	Destinations []struct {
		IP          string `json:"ip"`
		Description string `json:"description"`
		Default     bool   `json:"default"`
	} `json:"destinations"`
}

// Return the connectivity testing destinations for an MG network
func GetConnectivityTesting(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/connectivityMonitoringDestinations",
		api.BaseUrl(), networkId)
	var datamodel = ConnectivityTesting{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}