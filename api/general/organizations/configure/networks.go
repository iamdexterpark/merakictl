package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type Networks []struct {
	Network
}

type Network struct {
	ID               string   `json:"id"`
	OrganizationID   string   `json:"organizationId"`
	Name             string   `json:"name"`
	TimeZone         string   `json:"timeZone"`
	Tags             []string `json:"tags"`
	ProductTypes     []string `json:"productTypes"`
	EnrollmentString string   `json:"enrollmentString"`
}

// List the networks that the user has privileges on in an organization
func GetNetworks(organizationId string) (Networks, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/networks", api.BaseUrl(),
		organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Networks{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}