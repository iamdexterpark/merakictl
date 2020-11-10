package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

// Data Structure
type BrandingPolicies struct {
	BrandingPolicyIds []string `json:"brandingPolicyIds"`
}

type BrandingPolicy struct {
	BrandingPolicyID string `json:"brandingPolicyId"`
	Name             string `json:"name"`
	Enabled          bool   `json:"enabled"`
	AdminSettings    struct {
		AppliesTo string   `json:"appliesTo"`
		Values    []string `json:"values"`
	} `json:"adminSettings"`
	HelpSettings struct {
		HelpTab                            string `json:"helpTab"`
		GetHelpSubtab                      string `json:"getHelpSubtab"`
		CommunitySubtab                    string `json:"communitySubtab"`
		CasesSubtab                        string `json:"casesSubtab"`
		DataProtectionRequestsSubtab       string `json:"dataProtectionRequestsSubtab"`
		GetHelpSubtabKnowledgeBaseSearch   string `json:"getHelpSubtabKnowledgeBaseSearch"`
		UniversalSearchKnowledgeBaseSearch string `json:"universalSearchKnowledgeBaseSearch"`
		CiscoMerakiProductDocumentation    string `json:"ciscoMerakiProductDocumentation"`
		SupportContactInfo                 string `json:"supportContactInfo"`
		NewFeaturesSubtab                  string `json:"newFeaturesSubtab"`
		FirewallInfoSubtab                 string `json:"firewallInfoSubtab"`
		APIDocsSubtab                      string `json:"apiDocsSubtab"`
		HardwareReplacementsSubtab         string `json:"hardwareReplacementsSubtab"`
		SmForums                           string `json:"smForums"`
	} `json:"helpSettings"`
}

// Return The Branding Policy IDs Of An Organization
func GetBrandingPriority(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies/priorities", api.BaseUrl(), organizationId)

	var datamodel = BrandingPolicies{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return The Branding Policies Of An Organization
func GetBrandingPolicies(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies", api.BaseUrl(),
		organizationId)

	var datamodel = BrandingPolicy{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return The Branding Policy IDs Of An Organization
func GetBrandingPolicy(organizationId, brandingPolicyId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies/%s", api.BaseUrl(),
		organizationId, brandingPolicyId)

	var datamodel = BrandingPolicy{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
