package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetMTU(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/mtu",
		api.BaseUrl(), networkId)
	var datamodel = MTU{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
