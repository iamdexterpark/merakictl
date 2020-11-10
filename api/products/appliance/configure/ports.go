package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type PortVLANSettings []struct {
	PortVLANSetting
}
type PortVLANSetting struct {
	Number              int    `json:"number"`
	Enabled             bool   `json:"enabled"`
	Type                string `json:"type"`
	DropUntaggedTraffic bool   `json:"dropUntaggedTraffic"`
	Vlan                int    `json:"vlan"`
	AccessPolicy        string `json:"accessPolicy"`
}

// List per-port VLAN settings for all ports of a MX.
func GetPortVLANSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/ports", api.BaseUrl(), networkId)
	var datamodel = PortVLANSettings{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return per-port VLAN settings for a single MX port.
func GetPortVLANSetting(networkId, portId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/ports/%s", api.BaseUrl(), networkId, portId)
	var datamodel = PortVLANSetting{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
