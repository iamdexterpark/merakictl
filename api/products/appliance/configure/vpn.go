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


type ThirdPartyVPN struct {
	Peers []struct {
		Name           string   `json:"name"`
		PublicIP       string   `json:"publicIp"`
		PrivateSubnets []string `json:"privateSubnets"`
		Secret         string   `json:"secret"`
		IkeVersion     string   `json:"ikeVersion"`
		IpsecPolicies  struct {
			IkeCipherAlgo         []string `json:"ikeCipherAlgo"`
			IkeAuthAlgo           []string `json:"ikeAuthAlgo"`
			IkePrfAlgo            []string `json:"ikePrfAlgo"`
			IkeDiffieHellmanGroup []string `json:"ikeDiffieHellmanGroup"`
			IkeLifetime           int      `json:"ikeLifetime"`
			ChildCipherAlgo       []string `json:"childCipherAlgo"`
			ChildAuthAlgo         []string `json:"childAuthAlgo"`
			ChildPfsGroup         []string `json:"childPfsGroup"`
			ChildLifetime         int      `json:"childLifetime"`
		} `json:"ipsecPolicies,omitempty"`
		NetworkTags         []string `json:"networkTags"`
		RemoteID            string   `json:"remoteId,omitempty"`
		IpsecPoliciesPreset string   `json:"ipsecPoliciesPreset,omitempty"`
	} `json:"peers"`
}

// Return the third party VPN peers for an organization
func GetThirdPartyVPN(organizationId string ) (ThirdPartyVPN, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vpn/thirdPartyVPNPeers", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	var results = ThirdPartyVPN{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
