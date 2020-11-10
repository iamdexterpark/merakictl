package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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

type OrganizationIntrusionSettings struct {
	AllowedRules []struct {
		RuleID  string `json:"ruleId"`
		Message string `json:"message"`
	} `json:"allowedRules"`
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

// Returns all supported intrusion settings for an MX network
func GetMXIntrusionSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/security/intrusion", api.BaseUrl(), networkId)
	var datamodel = MXIntrusionSettings{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns all supported intrusion settings for an organization
func GetOrganizationIntrusionSettings(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/appliance/security/intrusion", api.BaseUrl(), organizationId)
	var datamodel = OrganizationIntrusionSettings{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns all supported malware settings for an MX network
func GetMalwareSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/security/malware", api.BaseUrl(), networkId)
	var datamodel = MalwareSettings{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
