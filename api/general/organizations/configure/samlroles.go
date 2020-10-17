package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)
type SAMLRoles []struct {
	SAMLRole
}
type SAMLRole struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	OrgAccess string `json:"orgAccess"`
	Networks  []struct {
		ID     string `json:"id"`
		Access string `json:"access"`
	} `json:"networks"`
	Tags []struct {
		Tag    string `json:"tag"`
		Access string `json:"access"`
	} `json:"tags"`
}

// List the SAML roles for this organization
func GetSAMLRoles(organizationId string) (SAMLRoles, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/samlRoles", api.BaseUrl(),
		organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SAMLRoles{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

func GetSAMLRole(organizationId, samlRoleId string) (SAMLRole, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/samlRoles/%s", api.BaseUrl(),
		organizationId, samlRoleId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SAMLRole{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
