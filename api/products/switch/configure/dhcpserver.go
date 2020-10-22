package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type DHCPServerPolicy struct {
	DefaultPolicy  string   `json:"defaultPolicy"`
	AllowedServers []string `json:"allowedServers"`
}

// Return The DHCP Server Policy
func GetDHCPServerPolicy(networkId string) (DHCPServerPolicy , interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/dhcpServerPolicy",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = DHCPServerPolicy {}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

