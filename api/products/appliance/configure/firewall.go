package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type CellularFirewallRules struct {
	Rules []struct {
		Comment       string `json:"comment"`
		Policy        string `json:"policy"`
		Protocol      string `json:"protocol"`
		DestPort      int    `json:"destPort"`
		DestCidr      string `json:"destCidr"`
		SrcPort       string `json:"srcPort"`
		SrcCidr       string `json:"srcCidr"`
		SyslogEnabled bool   `json:"syslogEnabled"`
	} `json:"rules"`
}

// Return the cellular firewall rules for an MX network
func GetCellularFirewallRules(networkId string ) (CellularFirewallRules, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/cellularFirewallRules", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = CellularFirewallRules{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


type FirewalledServices []struct {
	FirewalledService
}
type FirewalledService struct {
	Service    string   `json:"service"`
	Access     string   `json:"access"`
	AllowedIps []string `json:"allowedIps"`
}

// List the appliance services and their accessibility rules
func GetFirewalledServices(networkId string ) (FirewalledServices, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/firewalledServices", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = FirewalledServices{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// List the appliance services and their accessibility rules
func GetFirewalledService(networkId, serviceId string ) (FirewalledService, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/firewalledServices/%s", api.BaseUrl(), networkId, serviceId)
	var payload io.ReadSeeker
	var results = FirewalledService{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type InboundFirewallRules struct {
	Rules []struct {
		Comment       string `json:"comment"`
		Policy        string `json:"policy"`
		Protocol      string `json:"protocol"`
		DestPort      int    `json:"destPort"`
		DestCidr      string `json:"destCidr"`
		SrcPort       string `json:"srcPort"`
		SrcCidr       string `json:"srcCidr"`
		SyslogEnabled bool   `json:"syslogEnabled"`
	} `json:"rules"`
	SyslogDefaultRule bool `json:"syslogDefaultRule"`
}

// Return the inbound firewall rules for an MX network
func GetInboundFirewallRules(networkId string ) (InboundFirewallRules, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/inboundFirewallRules", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = InboundFirewallRules{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}