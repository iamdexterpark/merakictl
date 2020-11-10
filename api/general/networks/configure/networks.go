package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)


// Network - Return A Network
type Network struct {
	ID               string   `json:"id"`
	OrganizationID   string   `json:"organizationId"`
	Name             string   `json:"name"`
	TimeZone         string   `json:"timeZone"`
	Tags             []string `json:"tags"`
	ProductTypes     []string `json:"productTypes"`
	EnrollmentString string   `json:"enrollmentString"`
}

// GetNetwork - Return A Network
func GetNetwork(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s", api.BaseUrl(), networkId)
	var datamodel = Network{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}