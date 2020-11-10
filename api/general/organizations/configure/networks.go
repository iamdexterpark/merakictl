package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetNetworks(organizationId, configTemplateId, tags, tagsFilterType, perPage,
	startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/networks", api.BaseUrl(),
		organizationId)

	// Parameters for Request URL
	var parameters = map[string]string{
		"configTemplateId": configTemplateId,
		"tags":             tags,
		"tagsFilterType":   tagsFilterType,
		"perPage":          perPage,
		"startingAfter":    startingAfter,
		"endingBefore":     endingBefore}

	var datamodel = Networks{}
	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
