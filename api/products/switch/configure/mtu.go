package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type MTU struct {
	DefaultMtuSize int `json:"defaultMtuSize"`
	Overrides      []struct {
		Switches       []string `json:"switches,omitempty"`
		MtuSize        int      `json:"mtuSize"`
		SwitchProfiles []string `json:"switchProfiles,omitempty"`
	} `json:"overrides"`
}

// Return the MTU configuration
func GetMTU(networkId string) (MTU, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/mtu",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = MTU{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
