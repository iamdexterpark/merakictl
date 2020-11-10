package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type STPSettings struct {
	RstpEnabled       bool `json:"rstpEnabled"`
	StpBridgePriority []struct {
		Switches    []string `json:"switches,omitempty"`
		StpPriority int      `json:"stpPriority"`
		Stacks      []string `json:"stacks,omitempty"`
	} `json:"stpBridgePriority"`
}

// Returns STP Settings
func GetSTPSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stp",
		api.BaseUrl(), networkId)
	var datamodel = STPSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
