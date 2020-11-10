package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type SwitchPorts []struct {
	SwitchPort
}
type SwitchPort struct {
	PortID                  string   `json:"portId"`
	Name                    string   `json:"name"`
	Tags                    []string `json:"tags"`
	Enabled                 bool     `json:"enabled"`
	PoeEnabled              bool     `json:"poeEnabled"`
	Type                    string   `json:"type"`
	Vlan                    int      `json:"vlan"`
	VoiceVlan               int      `json:"voiceVlan"`
	IsolationEnabled        bool     `json:"isolationEnabled"`
	RstpEnabled             bool     `json:"rstpEnabled"`
	StpGuard                string   `json:"stpGuard"`
	LinkNegotiation         string   `json:"linkNegotiation"`
	PortScheduleID          string   `json:"portScheduleId"`
	Udld                    string   `json:"udld"`
	AccessPolicyType        string   `json:"accessPolicyType"`
	StickyMacAllowList      []string `json:"stickyMacAllowList"`
	StickyMacAllowListLimit int      `json:"stickyMacAllowListLimit"`
	StormControlEnabled     bool     `json:"stormControlEnabled"`
}

// List The Switch Ports For A Switch
func GetSwitchPorts(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/ports",
		api.BaseUrl(), serial)
	var datamodel = SwitchPorts{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return a Switch Port
func GetSwitchPort(serial, portId string) []api.Results{
	baseurl := fmt.Sprintf("%s/devices/%s/switch/ports/%s",
		api.BaseUrl(), serial, portId)
	var datamodel = SwitchPort{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}