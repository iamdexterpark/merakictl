package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type TargetGroups []struct {
	TargetGroup
}
type TargetGroup struct {
	Name  string `json:"name"`
	Scope string `json:"scope"`
	Tags  string `json:"tags"`
	Type  string `json:"type"`
}

// List The Target Groups In This Network
func GetTargetGroups(networkId, withDetails string) (TargetGroups, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/targetGroups",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("withDetails", withDetails)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = TargetGroups{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// Return A Target Group
func GetTargetGroup(networkId, targetGroupId, withDetails string) (TargetGroups, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/targetGroups/%s",
		api.BaseUrl(), networkId, targetGroupId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("withDetails", withDetails)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = TargetGroups{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}