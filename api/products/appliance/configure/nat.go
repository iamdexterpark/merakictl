package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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

type OneToOneNatRules struct {
	Rules []struct {
		Name           string `json:"name"`
		LanIP          string `json:"lanIp"`
		PublicIP       string `json:"publicIp"`
		Uplink         string `json:"uplink"`
		AllowedInbound []struct {
			Protocol         string   `json:"protocol"`
			DestinationPorts []string `json:"destinationPorts"`
			AllowedIps       []string `json:"allowedIps"`
		} `json:"allowedInbound"`
	} `json:"rules"`
}

type PortForwardingRules struct {
	Rules []struct {
		LanIP      string   `json:"lanIp"`
		AllowedIps []string `json:"allowedIps"`
		Name       string   `json:"name"`
		Protocol   string   `json:"protocol"`
		PublicPort string   `json:"publicPort"`
		LocalPort  string   `json:"localPort"`
		Uplink     string   `json:"uplink"`
	} `json:"rules"`
}

// Return the 1:Many NAT mapping rules for an MX network
func GetOneToManyNatRules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/oneToManyNatRules", api.BaseUrl(), networkId)
	var datamodel = OneToManyNatRules{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return the 1:1 NAT mapping rules for an MX network
func GetOneToOneNatRules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/oneToOneNatRules", api.BaseUrl(), networkId)
	var datamodel = OneToOneNatRules{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return the port forwarding rules for an MX network
func GetPortForwardingRules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/portForwardingRules", api.BaseUrl(), networkId)
	var datamodel = PortForwardingRules{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
