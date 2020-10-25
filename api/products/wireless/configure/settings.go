package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type WirelessSettings struct {
	MeshingEnabled           bool   `json:"meshingEnabled"`
	Ipv6BridgeEnabled        bool   `json:"ipv6BridgeEnabled"`
	LocationAnalyticsEnabled bool   `json:"locationAnalyticsEnabled"`
	UpgradeStrategy          string `json:"upgradeStrategy"`
}

// Return The Wireless Settings For A Network
func GetWirelessSettings(networkId string) (WirelessSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/settings",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = WirelessSettings{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

