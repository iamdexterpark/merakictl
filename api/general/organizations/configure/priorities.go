package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

// Data Structure
type BrandingPolicies struct {
	BrandingPolicyIds []string `json:"brandingPolicyIds"`
}

// Return The Branding Policy IDs Of An Organization
func GetBrandingPolicies(organizationId string) (BrandingPolicies, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies/priorities", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = BrandingPolicies{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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
func GetBrandingPolicy(organizationId, brandingPolicyId string) (BrandingPolicy, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies/priorities/%s", api.BaseUrl(),
		organizationId, brandingPolicyId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = BrandingPolicy{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}