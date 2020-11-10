package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

// NetworkSettings - Return The Settings For A Network
type NetworkSettings struct {
	LocalStatusPageEnabled  bool `json:"localStatusPageEnabled"`
	RemoteStatusPageEnabled bool `json:"remoteStatusPageEnabled"`
}

// GetNetworkSettings - Return The Settings For A Network
func GetNetworkSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/settings", api.BaseUrl(), networkId)
	var datamodel = NetworkSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
