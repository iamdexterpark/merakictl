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