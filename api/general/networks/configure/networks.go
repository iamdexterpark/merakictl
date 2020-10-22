package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)



type Network struct {
	ID               string   `json:"id"`
	OrganizationID   string   `json:"organizationId"`
	Name             string   `json:"name"`
	TimeZone         string   `json:"timeZone"`
	Tags             []string `json:"tags"`
	ProductTypes     []string `json:"productTypes"`
	EnrollmentString string   `json:"enrollmentString"`
}

// Return A Network
func GetNetwork(networkId string) (Network, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Network{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}