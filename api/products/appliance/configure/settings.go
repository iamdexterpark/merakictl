package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type ApplianceSettings struct {
	ClientTrackingMethod string `json:"clientTrackingMethod"`
}

// Return the appliance settings for a network
func GetApplianceSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/settings", api.BaseUrl(), networkId)
	var datamodel = ApplianceSettings{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
