package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type NetFlow struct {
	ReportingEnabled bool   `json:"reportingEnabled"`
	CollectorIP      string `json:"collectorIp"`
	CollectorPort    int    `json:"collectorPort"`
}

// Return The Net Flow Traffic Reporting Settings For A Network
func GetNetFlow(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/netflow", api.BaseUrl(), networkId)
	var datamodel = NetFlow{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
