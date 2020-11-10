package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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

type SwitchTemplateConfigProfiles []struct {
	SwitchProfileID string `json:"switchProfileId"`
	Name            string `json:"name"`
	Model           string `json:"model"`
}


// Return All The Ports Of A Switch Profile
func GetSwitchProfilePorts(organizationId, configTemplateId, profileId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s/switch/profiles/%s/ports",
		api.BaseUrl(), organizationId, configTemplateId, profileId)
	var datamodel = SwitchProfilePorts{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// Return A Switch Profile Port
func GetSwitchProfilePort(organizationId, configTemplateId, profileId, portId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s/switch/profiles/%s/ports/%s",
		api.BaseUrl(), organizationId, configTemplateId, profileId, portId)
	var datamodel = SwitchProfilePorts{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// List The Switch Profiles For Your Switch Template Configuration
func GetSwitchTemplateConfigProfiles(organizationId,
	configTemplateId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s/switch/profiles",
		api.BaseUrl(), organizationId, configTemplateId)
	var datamodel = SwitchTemplateConfigProfiles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}