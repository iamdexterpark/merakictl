package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type LANSettings struct {
	DeviceName         string `json:"deviceName"`
	DeviceLanIP        string `json:"deviceLanIp"`
	DeviceSubnet       string `json:"deviceSubnet"`
	FixedIPAssignments []struct {
		Mac  string `json:"mac"`
		Name string `json:"name"`
		IP   string `json:"ip"`
	} `json:"fixedIpAssignments"`
	ReservedIPRanges []struct {
		Start   string `json:"start"`
		End     string `json:"end"`
		Comment string `json:"comment"`
	} `json:"reservedIpRanges"`
}

// Show the LAN Settings of a MG
func GetLANSettings(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/cellularGateway/lan",
		api.BaseUrl(), serial)
	var datamodel = LANSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
