package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type SiteToSiteVPN struct {
	Mode string `json:"mode"`
	Hubs []struct {
		HubID           string `json:"hubId"`
		UseDefaultRoute bool   `json:"useDefaultRoute"`
	} `json:"hubs"`
	Subnets []struct {
		LocalSubnet string `json:"localSubnet"`
		UseVpn      bool   `json:"useVpn"`
	} `json:"subnets"`
}

// Return the site-to-site VPN settings of a network.
func GetSiteToSiteVPN(networkId string ) (SiteToSiteVPN, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/vpn/siteToSiteVpn", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = SiteToSiteVPN{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}