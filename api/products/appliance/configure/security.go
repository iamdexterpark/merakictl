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

type OrganizationIntrusionSettings struct {
	AllowedRules []struct {
		RuleID  string `json:"ruleId"`
		Message string `json:"message"`
	} `json:"allowedRules"`
}

// Returns all supported intrusion settings for an organization
func GetOrganizationIntrusionSettings(organizationId string ) (OrganizationIntrusionSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/appliance/security/intrusion", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	var results = OrganizationIntrusionSettings{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type MalwareSettings struct {
	Mode        string `json:"mode"`
	AllowedUrls []struct {
		URL     string `json:"url"`
		Comment string `json:"comment"`
	} `json:"allowedUrls"`
	AllowedFiles []struct {
		Sha256  string `json:"sha256"`
		Comment string `json:"comment"`
	} `json:"allowedFiles"`
}

// Returns all supported malware settings for an MX network
func GetMalwareSettings(networkId string ) (MalwareSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/security/malware", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = MalwareSettings{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}