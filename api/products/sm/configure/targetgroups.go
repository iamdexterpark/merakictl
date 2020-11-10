package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetTargetGroups(networkId, withDetails string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/targetGroups",
		api.BaseUrl(), networkId)
	var datamodel = TargetGroups{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"withDetails": withDetails}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return A Target Group
func GetTargetGroup(networkId, targetGroupId, withDetails string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/targetGroups/%s",
		api.BaseUrl(), networkId, targetGroupId)
	var datamodel = TargetGroups{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"withDetails": withDetails}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}