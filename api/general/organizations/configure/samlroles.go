package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetSAMLRoles(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/samlRoles", api.BaseUrl(),
		organizationId)
	var datamodel = SAMLRoles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSAMLRole(organizationId, samlRoleId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/samlRoles/%s", api.BaseUrl(),
		organizationId, samlRoleId)
	var datamodel = SAMLRole{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
