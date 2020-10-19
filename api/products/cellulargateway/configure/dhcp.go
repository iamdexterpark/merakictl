package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type DHCPSettings  struct {
	DhcpLeaseTime        string   `json:"dhcpLeaseTime"`
	DNSNameservers       string   `json:"dnsNameservers"`
	DNSCustomNameservers []string `json:"dnsCustomNameservers"`
}

// List common DHCP settings of MGs
func GetDHCPSettings(networkId string) (DHCPSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/dhcp",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = DHCPSettings{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
