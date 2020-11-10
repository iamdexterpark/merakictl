package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type DHCPSettings struct {
	DhcpLeaseTime        string   `json:"dhcpLeaseTime"`
	DNSNameservers       string   `json:"dnsNameservers"`
	DNSCustomNameservers []string `json:"dnsCustomNameservers"`
}

// List common DHCP settings of MGs
func GetDHCPSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/dhcp",
		api.BaseUrl(), networkId)
	var datamodel = DHCPSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
