package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetPortVLANSettings(networkId string ) (PortVLANSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/ports", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = PortVLANSettings{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// Return per-port VLAN settings for a single MX port.
func GetPortVLANSetting(networkId, portId string ) (PortVLANSetting, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/ports/%s", api.BaseUrl(), networkId, portId)
	var payload io.ReadSeeker
	var results = PortVLANSetting{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}