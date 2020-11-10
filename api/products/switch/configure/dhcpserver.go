package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type DHCPServerPolicy struct {
	DefaultPolicy  string   `json:"defaultPolicy"`
	AllowedServers []string `json:"allowedServers"`
}

// Return The DHCP Server Policy
func GetDHCPServerPolicy(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/dhcpServerPolicy",
		api.BaseUrl(), networkId)
	var datamodel = DHCPServerPolicy{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
