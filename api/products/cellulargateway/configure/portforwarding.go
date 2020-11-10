package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type PortForwardingRules struct {
	Rules []struct {
		LanIP      string   `json:"lanIp"`
		Name       string   `json:"name"`
		Access     string   `json:"access"`
		PublicPort string   `json:"publicPort"`
		LocalPort  string   `json:"localPort"`
		Uplink     string   `json:"uplink"`
		Protocol   string   `json:"protocol"`
		AllowedIps []string `json:"allowedIps,omitempty"`
	} `json:"rules"`
}

// Returns the port forwarding rules for a single MG
func GetPortForwardingRules(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/cellularGateway/portForwardingRules",
		api.BaseUrl(), serial)
	var datamodel = PortForwardingRules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}