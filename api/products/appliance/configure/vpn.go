package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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

type FirewallRules struct {
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

// Return the site-to-site VPN settings of a network.
func GetSiteToSiteVPN(networkId string ) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/vpn/siteToSiteVpn", api.BaseUrl(), networkId)
	var datamodel = SiteToSiteVPN{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return the third party VPN peers for an organization
func GetThirdPartyVPN(organizationId string ) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vpn/thirdPartyVPNPeers", api.BaseUrl(), organizationId)
	var datamodel = ThirdPartyVPN{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return the firewall rules for an organization's site-to-site VPN
func GetFirewallRules(organizationId string ) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vpn/vpnFirewallRules", api.BaseUrl(), organizationId)
	var datamodel = FirewallRules{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

