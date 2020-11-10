package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type UplinkSettings struct {
	BandwidthLimits struct {
		LimitUp   int `json:"limitUp"`
		LimitDown int `json:"limitDown"`
	} `json:"bandwidthLimits"`
}

// Returns The Uplink Settings For Your MG Network
func GetUplinkSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/uplink",
		api.BaseUrl(), networkId)
	var datamodel = UplinkSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}