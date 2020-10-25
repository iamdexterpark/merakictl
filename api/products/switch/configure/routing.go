package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type Layer3DHCPInterface struct {
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

// Return A Layer 3 Interface DHCP Configuration For A Switch
func GetLayer3DHCPInterface(serial, interfaceId string) (Layer3DHCPInterface, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/interfaces/%s/dhcp",
		api.BaseUrl(), serial, interfaceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Layer3DHCPInterface{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type Layer3Interfaces []struct {
	Layer3Interface
}
type Layer3Interface struct {
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

// List Layer 3 Interfaces For A Switch
func GetLayer3Interfaces(serial string) (Layer3Interfaces, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/interfaces",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Layer3Interfaces{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return A Layer 3 Interface For A Switch
func GetLayer3Interface(serial, interfaceId string) (Layer3Interface, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/interfaces/%s",
		api.BaseUrl(), serial, interfaceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Layer3Interface{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type MulticastRendezvousPoints [][]struct {
	MulticastRendezvousPoint
}


	type MulticastRendezvousPoint []struct {
	RendezvousPointID string `json:"rendezvousPointId"`
	Serial            string `json:"serial,omitempty"`
	InterfaceName     string `json:"interfaceName,omitempty"`
	InterfaceIP       string `json:"interfaceIp"`
	MulticastGroup    string `json:"multicastGroup"`
	SwitchStackID     string `json:"switchStackId,omitempty"`
}

// List Multicast Rendezvous Points
func GetMulticastRendezvousPoints(networkId string) (MulticastRendezvousPoints, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/routing/multicast/rendezvousPoints",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = MulticastRendezvousPoints{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return A Multicast Rendezvous Point
func GetMulticastRendezvousPoint(networkId, rendezvousPointId string) (MulticastRendezvousPoint, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/routing/multicast/rendezvousPoints/%s",
		api.BaseUrl(), networkId, rendezvousPointId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = MulticastRendezvousPoint{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type MulticastSettings struct {
	DefaultSettings struct {
		IgmpSnoopingEnabled                 bool `json:"igmpSnoopingEnabled"`
		FloodUnknownMulticastTrafficEnabled bool `json:"floodUnknownMulticastTrafficEnabled"`
	} `json:"defaultSettings"`
	Overrides []struct {
		Switches                            []string `json:"switches,omitempty"`
		IgmpSnoopingEnabled                 bool     `json:"igmpSnoopingEnabled"`
		FloodUnknownMulticastTrafficEnabled bool     `json:"floodUnknownMulticastTrafficEnabled"`
		Stacks                              []string `json:"stacks,omitempty"`
	} `json:"overrides"`
}

// Return Multicast Settings For A Network
func GetMulticastSettings(networkId string) (MulticastSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/routing/multicast",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = MulticastSettings{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}