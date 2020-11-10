package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type Device []struct {
	Name      string   `json:"name"`
	Lat       float64  `json:"lat"`
	Lng       float64  `json:"lng"`
	Address   string   `json:"address"`
	Notes     string   `json:"notes"`
	Tags      []string `json:"tags"`
	NetworkID string   `json:"networkId"`
	Serial    string   `json:"serial"`
	Model     string   `json:"model"`
	Mac       string   `json:"mac"`
	LanIP     string   `json:"lanIp"`
	Firmware  string   `json:"firmware"`
}

// List the devices in an organization
func GetDeviceList(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/devices", api.BaseUrl(),
		organizationId)

	var datamodel = Device{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
