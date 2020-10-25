package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetAlternateManagementInterface(networkId string) (AlternateManagementInterface, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/alternateManagementInterface",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = AlternateManagementInterface{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}