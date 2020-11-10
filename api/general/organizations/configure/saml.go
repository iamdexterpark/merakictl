package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type LDPS []struct {
	LDP
}

type LDP struct {
	IdpID                   string `json:"idpId"`
	ConsumerURL             string `json:"consumerUrl"`
	X509CertSha1Fingerprint string `json:"x509certSha1Fingerprint"`
	SloLogoutURL            string `json:"sloLogoutUrl"`
}

type SAML struct {
	Enabled bool `json:"enabled"`
}


// List the SAML IdPs in your organization.
func GetLDPS(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml/idps", api.BaseUrl(),
		organizationId)
	var datamodel = LDPS{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List a SAML IdP in your organization.
func GetLDP(organizationId, ldpId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml/idp/%s", api.BaseUrl(),
		organizationId, ldpId)

	var datamodel = LDP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}



// Returns the SAML SSO enabled settings for an organization.
func GetSAML(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml", api.BaseUrl(),
		organizationId)

	var datamodel = SAML{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


