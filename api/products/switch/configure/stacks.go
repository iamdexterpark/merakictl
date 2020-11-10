package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type InterfaceDHCPConfiguration struct {
	DhcpMode             string   `json:"dhcpMode"`
	DhcpLeaseTime        string   `json:"dhcpLeaseTime"`
	DNSNameserversOption string   `json:"dnsNameserversOption"`
	DNSCustomNameservers []string `json:"dnsCustomNameservers"`
	BootOptionsEnabled   bool     `json:"bootOptionsEnabled"`
	BootNextServer       string   `json:"bootNextServer"`
	BootFileName         string   `json:"bootFileName"`
	DhcpOptions          []struct {
		Code  string `json:"code"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"dhcpOptions"`
	ReservedIPRanges []struct {
		Start   string `json:"start"`
		End     string `json:"end"`
		Comment string `json:"comment"`
	} `json:"reservedIpRanges"`
	FixedIPAssignments []struct {
		Mac  string `json:"mac"`
		Name string `json:"name"`
		IP   string `json:"ip"`
	} `json:"fixedIpAssignments"`
}

type StackLayer3Interfaces []struct {
	StackLayer3Interface
}
type StackLayer3Interface struct {
	InterfaceID      string `json:"interfaceId"`
	Name             string `json:"name"`
	Subnet           string `json:"subnet"`
	InterfaceIP      string `json:"interfaceIp"`
	MulticastRouting string `json:"multicastRouting"`
	VlanID           int    `json:"vlanId"`
	DefaultGateway   string `json:"defaultGateway"`
	OspfSettings     struct {
		Area             string `json:"area"`
		Cost             int    `json:"cost"`
		IsPassiveEnabled bool   `json:"isPassiveEnabled"`
	} `json:"ospfSettings"`
}

type StackStaticRoutes []struct {
	StackStaticRoute
}

type StackStaticRoute struct {
	InterfaceID      string `json:"interfaceId"`
	Name             string `json:"name"`
	Subnet           string `json:"subnet"`
	InterfaceIP      string `json:"interfaceIp"`
	MulticastRouting string `json:"multicastRouting"`
	VlanID           int    `json:"vlanId"`
	DefaultGateway   string `json:"defaultGateway"`
	OspfSettings     struct {
		Area             string `json:"area"`
		Cost             int    `json:"cost"`
		IsPassiveEnabled bool   `json:"isPassiveEnabled"`
	} `json:"ospfSettings"`
}

type SwitchStacks []struct {
	SwitchStack
}

type SwitchStack struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Serials []string `json:"serials"`
}

// Return A Layer 3 Interface DHCP Configuration For A Switch Stack
func GetInterfaceDHCPConfiguration(networkId, switchStackId, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%s/routing/interfaces/%s/dhcp",
		api.BaseUrl(), networkId, switchStackId, interfaceId)
	var datamodel = InterfaceDHCPConfiguration{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List Layer 3 Interfaces For A Switch Stack
func GetStackLayer3Interfaces(networkId, switchStackId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%s/routing/interfaces",
		api.BaseUrl(), networkId, switchStackId)
	var datamodel = StackLayer3Interfaces{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return A Layer 3 Interface From A Switch Stack
func GetStackLayer3Interface(networkId, switchStackId, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%s/routing/interfaces/%s",
		api.BaseUrl(), networkId, switchStackId, interfaceId)
	var datamodel = StackLayer3Interface{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List Layer 3 Static Routes For A Switch Stack
func GetStackStaticRoutes(networkId, switchStackId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%srouting/interfaces",
		api.BaseUrl(), networkId, switchStackId)
	var datamodel = StackStaticRoutes{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return A Layer 3 Static Route For A Switch Stack
func GetStackStaticRoute(networkId, switchStackId, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%srouting/interfaces/%s",
		api.BaseUrl(), networkId, switchStackId, interfaceId)
	var datamodel = StackStaticRoute{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List The Switch Stacks In A Network
func GetSwitchStacks(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks",
		api.BaseUrl(), networkId)
	var datamodel = SwitchStacks{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Show A Switch Stack
func GetSwitchStack(networkId, switchStackId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%s",
		api.BaseUrl(), networkId, switchStackId)
	var datamodel = SwitchStack{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}