package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type SwitchProfilePorts struct {
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

// Return All The Ports Of A Switch Profile
func GetSwitchProfilePorts(organizationId, configTemplateId, profileId string) (SwitchProfilePorts, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s/switch/profiles/%s/ports",
		api.BaseUrl(), organizationId, configTemplateId, profileId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SwitchProfilePorts{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// Return A Switch Profile Port
func GetSwitchProfilePort(organizationId, configTemplateId, profileId, portId string) (SwitchProfilePorts, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s/switch/profiles/%s/ports/%s",
		api.BaseUrl(), organizationId, configTemplateId, profileId, portId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SwitchProfilePorts{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// List The Switch Profiles For Your Switch Template Configuration