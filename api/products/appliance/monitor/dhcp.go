package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type DHCP []struct {
	Subnet    string `json:"subnet"`
	VlanID    int    `json:"vlanId"`
	UsedCount int    `json:"usedCount"`
	FreeCount int    `json:"freeCount"`
}

// Return the DHCP subnet information for an appliance
func GetDHCP(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/appliance/dhcp/subnets", api.BaseUrl(), serial)
	var datamodel = DHCP{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
