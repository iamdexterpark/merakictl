package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetNetworkProfiles(networkId string) (NetworkProfiles, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/profiles",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = NetworkProfiles{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}