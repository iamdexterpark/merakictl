package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type VLANStatus struct {
	VlansEnabled bool `json:"vlansEnabled"`
}

// Returns the enabled status of VLANs for the network
func GetVLANStatus(networkId string ) (VLANStatus, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/settings", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = VLANStatus{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

