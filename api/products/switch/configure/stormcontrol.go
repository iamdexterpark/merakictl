package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type StormControl struct {
	BroadcastThreshold      int `json:"broadcastThreshold"`
	MulticastThreshold      int `json:"multicastThreshold"`
	UnknownUnicastThreshold int `json:"unknownUnicastThreshold"`
}

// Return The Storm Control Configuration For A Switch Network
func GetStormControl(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stormControl",
		api.BaseUrl(), networkId)
	var datamodel = StormControl{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
