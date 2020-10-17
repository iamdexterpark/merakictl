package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type OneToManyNatRules struct {
	Rules []struct {
		PublicIP  string `json:"publicIp"`
		Uplink    string `json:"uplink"`
		PortRules []struct {
			Name       string   `json:"name"`
			Protocol   string   `json:"protocol"`
			PublicPort string   `json:"publicPort"`
			LocalIP    string   `json:"localIp"`
			LocalPort  string   `json:"localPort"`
			AllowedIps []string `json:"allowedIps"`
		} `json:"portRules"`
	} `json:"rules"`
}

// Return the 1:Many NAT mapping rules for an MX network
func GetOneToManyNatRules(networkId string ) (OneToManyNatRules, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/oneToManyNatRules", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = OneToManyNatRules{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}