package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetSwitchPorts(serial string) (SwitchPorts, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/ports",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SwitchPorts{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return a Switch Port
func GetSwitchPort(serial, portId string) (SwitchPort, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/ports/%s",
		api.BaseUrl(), serial, portId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SwitchPort{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}