package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type NetworkSettings struct {
	LocalStatusPageEnabled  bool `json:"localStatusPageEnabled"`
	RemoteStatusPageEnabled bool `json:"remoteStatusPageEnabled"`
}

// Return The Settings For A Network
func GetNetworkSettings(networkId string) (NetworkSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/settings", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = NetworkSettings{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
