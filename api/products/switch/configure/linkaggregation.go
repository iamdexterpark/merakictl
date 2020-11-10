package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type LinkAggregationGroups []struct {
	LinkAggregationGroup
}
type LinkAggregationGroup struct {
	ID          string `json:"id"`
	SwitchPorts []struct {
		Serial string `json:"serial"`
		PortID string `json:"portId"`
	} `json:"switchPorts"`
}

// List link aggregation groups
func GetLinkAggregationGroups(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/linkAggregations",
		api.BaseUrl(), networkId)
	var datamodel = LinkAggregationGroups{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
