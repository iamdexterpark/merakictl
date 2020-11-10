package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type LANConfiguration struct {
	Subnet      string `json:"subnet"`
	ApplianceIP string `json:"applianceIp"`
}

// Return single LAN configuration
func GetLANConfiguration(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/singleLan", api.BaseUrl(), networkId)
	var datamodel = LANConfiguration{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
