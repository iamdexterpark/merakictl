package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type DSCPCOSMapping struct {
	Mappings []struct {
		Dscp  int    `json:"dscp"`
		Cos   int    `json:"cos"`
		Title string `json:"title"`
	} `json:"mappings"`
}

// Return The DSCP To CoS Mappings
func GetDSCPCOSMapping(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/dscpToCosMappings",
		api.BaseUrl(), networkId)
	var datamodel = DSCPCOSMapping{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
