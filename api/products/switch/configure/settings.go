package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetSwitchNetworkSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/settings",
		api.BaseUrl(), networkId)
	var datamodel = SwitchNetworkSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
