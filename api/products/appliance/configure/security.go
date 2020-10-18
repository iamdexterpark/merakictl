package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type MXIntrusionSettings struct {
	Mode              string `json:"mode"`
	IdsRulesets       string `json:"idsRulesets"`
	ProtectedNetworks struct {
		UseDefault   bool     `json:"useDefault"`
		IncludedCidr []string `json:"includedCidr"`
		ExcludedCidr []string `json:"excludedCidr"`
	} `json:"protectedNetworks"`
}

// Returns all supported intrusion settings for an MX network
func GetMXIntrusionSettings(networkId string ) (MXIntrusionSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/security/intrusion", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = MXIntrusionSettings{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}