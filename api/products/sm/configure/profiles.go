package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type NetworkProfiles []struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Scope         string      `json:"scope"`
	Tags          []string    `json:"tags"`
	TargetGroupID interface{} `json:"targetGroupId"`
}

// List all profiles in a network
func GetNetworkProfiles(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/profiles",
		api.BaseUrl(), networkId)
	var datamodel = NetworkProfiles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
