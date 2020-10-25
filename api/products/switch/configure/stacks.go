package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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

// Return A Layer 3 Interface DHCP Configuration For A Switch Stack
func GetInterfaceDHCPConfiguration(networkId, switchStackId, interfaceId string) (InterfaceDHCPConfiguration, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%s/routing/interfaces/%s/dhcp",
		api.BaseUrl(), networkId, switchStackId, interfaceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = InterfaceDHCPConfiguration{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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
// List Layer 3 Interfaces For A Switch Stack
func GetStackLayer3Interfaces(networkId, switchStackId string) (StackLayer3Interfaces, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%s/routing/interfaces",
		api.BaseUrl(), networkId, switchStackId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = StackLayer3Interfaces{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return A Layer 3 Interface From A Switch Stack
func GetStackLayer3Interface(networkId, switchStackId, interfaceId string) (StackLayer3Interface, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%s/routing/interfaces/%s",
		api.BaseUrl(), networkId, switchStackId, interfaceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = StackLayer3Interface{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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

// List Layer 3 Static Routes For A Switch Stack
func GetStackStaticRoutes(networkId, switchStackId string) (StackStaticRoutes, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%srouting/interfaces",
		api.BaseUrl(), networkId, switchStackId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = StackStaticRoutes{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return A Layer 3 Static Route For A Switch Stack
func GetStackStaticRoute(networkId, switchStackId, interfaceId string) (StackStaticRoute, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%srouting/interfaces/%s",
		api.BaseUrl(), networkId, switchStackId, interfaceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = StackStaticRoute{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type SwitchStacks []struct {
	SwitchStack
}

type SwitchStack struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Serials []string `json:"serials"`
}

// List The Switch Stacks In A Network
func GetSwitchStacks(networkId string) (SwitchStacks, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SwitchStacks{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Show A Switch Stack
func GetSwitchStack(networkId, switchStackId string) (SwitchStack, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stacks/%s",
		api.BaseUrl(), networkId, switchStackId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SwitchStack{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}