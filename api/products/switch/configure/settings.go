package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type SwitchNetworkSettings struct {
	Vlan             int  `json:"vlan"`
	UseCombinedPower bool `json:"useCombinedPower"`
	PowerExceptions  []struct {
		Serial    string `json:"serial"`
		PowerType string `json:"powerType"`
	} `json:"powerExceptions"`
}

// Returns The Switch Network Settings
func GetSwitchNetworkSettings(networkId string) (SwitchNetworkSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/settings",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SwitchNetworkSettings{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}