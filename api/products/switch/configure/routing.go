package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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

// Return A Layer 3 Interface DHCP Configuration For A Switch
func GetLayer3DHCPInterface(serial, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/interfaces/%s/dhcp",
		api.BaseUrl(), serial, interfaceId)
	var datamodel = Layer3DHCPInterface{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List Layer 3 Interfaces For A Switch
func GetLayer3Interfaces(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/interfaces",
		api.BaseUrl(), serial)
	var datamodel = Layer3Interfaces{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return A Layer 3 Interface For A Switch
func GetLayer3Interface(serial, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/interfaces/%s",
		api.BaseUrl(), serial, interfaceId)
	var datamodel = Layer3Interface{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List Multicast Rendezvous Points
func GetMulticastRendezvousPoints(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/routing/multicast/rendezvousPoints",
		api.BaseUrl(), networkId)
	var datamodel = MulticastRendezvousPoints{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return A Multicast Rendezvous Point
func GetMulticastRendezvousPoint(networkId, rendezvousPointId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/routing/multicast/rendezvousPoints/%s",
		api.BaseUrl(), networkId, rendezvousPointId)
	var datamodel = MulticastRendezvousPoint{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return Multicast Settings For A Network
func GetMulticastSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/routing/multicast",
		api.BaseUrl(), networkId)
	var datamodel = MulticastSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
