package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)
type Devices []struct {
	Device
}
type Device struct {
	Name           string  `json:"name"`
	Lat            float64 `json:"lat"`
	Lng            float64 `json:"lng"`
	Serial         string  `json:"serial"`
	Mac            string  `json:"mac"`
	Model          string  `json:"model"`
	Address        string  `json:"address"`
	Notes          string  `json:"notes"`
	LanIP          string  `json:"lanIp"`
	Tags           string  `json:"tags"`
	NetworkID      string  `json:"networkId"`
	BeaconIDParams struct {
		UUID  string `json:"uuid"`
		Major int    `json:"major"`
		Minor int    `json:"minor"`
	} `json:"beaconIdParams"`
	Firmware    string `json:"firmware"`
	FloorPlanID string `json:"floorPlanId"`
}

// GetNetworkDevices - List The Devices In A Network
func GetNetworkDevices(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/devices", api.BaseUrl(), networkId)
	var datamodel = Devices{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}