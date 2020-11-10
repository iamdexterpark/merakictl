package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type AlternateManagementInterface struct {
	Enabled      bool     `json:"enabled"`
	VlanID       int      `json:"vlanId"`
	Protocols    []string `json:"protocols"`
	AccessPoints []struct {
		Serial                string `json:"serial"`
		AlternateManagementIP string `json:"alternateManagementIp"`
		SubnetMask            string `json:"subnetMask"`
		Gateway               string `json:"gateway"`
		DNS1                  string `json:"dns1"`
		DNS2                  string `json:"dns2"`
	} `json:"accessPoints"`
}

// Return Alternate Management Interface And Devices With IP Assigned
func GetAlternateManagementInterface(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/alternateManagementInterface",
		api.BaseUrl(), networkId)
	var datamodel = AlternateManagementInterface{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
