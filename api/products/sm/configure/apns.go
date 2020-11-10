package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type APNSCertificate struct {
	Certificate string `json:"certificate"`
}

// Get the organization's APNS certificate
func GetAPNSCertificate(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/sm/apnsCert", api.BaseUrl(), organizationId)
	var datamodel = APNSCertificate{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
