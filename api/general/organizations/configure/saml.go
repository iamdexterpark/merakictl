package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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

// List the SAML IdPs in your organization.
func GetLDPS(organizationId string) (LDPS, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml/idps", api.BaseUrl(),
		organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = LDPS{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// List a SAML IdP in your organization.
func GetLDP(organizationId, ldpId string) (LDP, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml/idp/%s", api.BaseUrl(),
		organizationId, ldpId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = LDP{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type SAML struct {
	Enabled bool `json:"enabled"`
}


// Returns the SAML SSO enabled settings for an organization.
func GetSAML(organizationId string) (SAML, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml", api.BaseUrl(),
		organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SAML{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


