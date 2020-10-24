package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetLinkAggregationGroups(networkId string) (LinkAggregationGroups, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/linkAggregations",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = LinkAggregationGroups{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
