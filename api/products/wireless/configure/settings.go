package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type WirelessSettings struct {
	MeshingEnabled           bool   `json:"meshingEnabled"`
	Ipv6BridgeEnabled        bool   `json:"ipv6BridgeEnabled"`
	LocationAnalyticsEnabled bool   `json:"locationAnalyticsEnabled"`
	UpgradeStrategy          string `json:"upgradeStrategy"`
}

// Return The Wireless Settings For A Network
func GetWirelessSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/settings",
		api.BaseUrl(), networkId)
	var datamodel = WirelessSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

